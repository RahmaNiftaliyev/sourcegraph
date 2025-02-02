package webhooks

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sourcegraph/log/logtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/database/dbtest"
	"github.com/sourcegraph/sourcegraph/internal/encryption/keyring"
	"github.com/sourcegraph/sourcegraph/internal/extsvc"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/gitlab/webhooks"
	"github.com/sourcegraph/sourcegraph/internal/types"
)

func TestWebhooksHandler(t *testing.T) {
	logger := logtest.Scoped(t)
	db := database.NewDB(logger, dbtest.NewDB(logger, t))
	u, err := db.Users().Create(context.Background(), database.NewUser{
		Email:           "test@user.com",
		Username:        "testuser",
		EmailIsVerified: true,
	})
	require.NoError(t, err)
	dbWebhooks := db.Webhooks(keyring.Default().WebhookKey)
	gitLabWH, err := dbWebhooks.Create(
		context.Background(),
		extsvc.KindGitLab,
		"http://gitlab.com",
		u.ID,
		types.NewUnencryptedSecret("somesecret"))
	require.NoError(t, err)

	gitHubWH, err := dbWebhooks.Create(
		context.Background(),
		extsvc.KindGitHub,
		"http://github.com",
		u.ID,
		types.NewUnencryptedSecret("githubsecret"),
	)
	require.NoError(t, err)

	gitHubWHNoSecret, err := dbWebhooks.Create(
		context.Background(),
		extsvc.KindGitHub,
		"http://github.com",
		u.ID,
		nil,
	)

	require.NoError(t, err)
	wr := WebhookRouter{
		DB: db,
	}
	gh := GitHubWebhook{WebhookRouter: &wr}

	webhookMiddleware := NewLogMiddleware(
		db.WebhookLogs(keyring.Default().WebhookLogKey),
	)

	base := mux.NewRouter()
	base.Path("/.api/webhooks/{webhook_uuid}").Methods("POST").Handler(webhookMiddleware.Logger(NewHandler(logger, db, gh.WebhookRouter)))
	srv := httptest.NewServer(base)

	t.Run("found GitLab webhook with correct secret returns 200", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/%v", srv.URL, gitLabWH.UUID)

		event := webhooks.EventCommon{
			ObjectKind: "pipeline",
		}
		wr.handlers = map[string]webhookEventHandlers{
			extsvc.KindGitLab: {
				"pipeline": []WebhookHandler{fakeWebhookHandler},
			},
		}
		payload, err := json.Marshal(event)
		require.NoError(t, err)
		req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
		require.NoError(t, err)
		req.Header.Add("X-GitLab-Token", "somesecret")
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("not-found webhook returns 404", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/%v", srv.URL, uuid.New())

		resp, err := http.Post(requestURL, "", nil)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("malformed UUID returns 400", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/SomeInvalidUUID", srv.URL)

		resp, err := http.Post(requestURL, "", nil)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("incorrect GitLab secret returns 400", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/%v", srv.URL, gitLabWH.UUID)

		req, err := http.NewRequest("POST", requestURL, nil)
		require.NoError(t, err)
		req.Header.Add("X-GitLab-Token", "someothersecret")
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("correct GitHub secret returns 200", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/%v", srv.URL, gitHubWH.UUID)

		h := hmac.New(sha1.New, []byte("githubsecret"))
		payload := []byte(`{"body": "text"}`)
		h.Write(payload)
		res := h.Sum(nil)

		wr.handlers = map[string]webhookEventHandlers{
			extsvc.KindGitHub: {
				"member": []WebhookHandler{fakeWebhookHandler},
			},
		}

		req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
		require.NoError(t, err)
		req.Header.Set("X-Hub-Signature", "sha1="+hex.EncodeToString(res))
		req.Header.Set("X-Github-Event", "member")
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		logs, _, err := db.WebhookLogs(keyring.Default().WebhookLogKey).List(context.Background(), database.WebhookLogListOpts{
			WebhookID: &gitHubWH.ID,
		})
		assert.NoError(t, err)
		assert.Len(t, logs, 1)
		for _, log := range logs {
			assert.Equal(t, gitHubWH.ID, *log.WebhookID)
		}
	})

	t.Run("GitHub with no secret returns 200", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/%v", srv.URL, gitHubWHNoSecret.UUID)

		payload := []byte(`{"body": "text"}`)

		req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Github-Event", "member")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("incorrect GitHub secret returns 400", func(t *testing.T) {
		requestURL := fmt.Sprintf("%s/.api/webhooks/%v", srv.URL, gitHubWH.UUID)

		h := hmac.New(sha1.New, []byte("wrongsecret"))
		payload := []byte(`{"body": "text"}`)
		h.Write(payload)
		res := h.Sum(nil)

		req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
		require.NoError(t, err)
		req.Header.Set("X-Hub-Signature", "sha1="+hex.EncodeToString(res))
		req.Header.Set("X-Github-Event", "member")
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func fakeWebhookHandler(ctx context.Context, db database.DB, codeHostURN extsvc.CodeHostBaseURL, event any) error {
	return nil
}
