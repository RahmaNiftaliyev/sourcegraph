// Code generated by go-mockgen 1.3.6; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package releasecache

import (
	"context"
	"sync"
)

// MockReleaseCache is a mock implementation of the ReleaseCache interface
// (from the package
// github.com/sourcegraph/sourcegraph/cmd/frontend/internal/httpapi/releasecache)
// used for unit testing.
type MockReleaseCache struct {
	// CurrentFunc is an instance of a mock function object controlling the
	// behavior of the method Current.
	CurrentFunc *ReleaseCacheCurrentFunc
	// HandleFunc is an instance of a mock function object controlling the
	// behavior of the method Handle.
	HandleFunc *ReleaseCacheHandleFunc
	// UpdateNowFunc is an instance of a mock function object controlling
	// the behavior of the method UpdateNow.
	UpdateNowFunc *ReleaseCacheUpdateNowFunc
}

// NewMockReleaseCache creates a new mock of the ReleaseCache interface. All
// methods return zero values for all results, unless overwritten.
func NewMockReleaseCache() *MockReleaseCache {
	return &MockReleaseCache{
		CurrentFunc: &ReleaseCacheCurrentFunc{
			defaultHook: func(string) (r0 string, r1 error) {
				return
			},
		},
		HandleFunc: &ReleaseCacheHandleFunc{
			defaultHook: func(context.Context) (r0 error) {
				return
			},
		},
		UpdateNowFunc: &ReleaseCacheUpdateNowFunc{
			defaultHook: func(context.Context) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockReleaseCache creates a new mock of the ReleaseCache
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockReleaseCache() *MockReleaseCache {
	return &MockReleaseCache{
		CurrentFunc: &ReleaseCacheCurrentFunc{
			defaultHook: func(string) (string, error) {
				panic("unexpected invocation of MockReleaseCache.Current")
			},
		},
		HandleFunc: &ReleaseCacheHandleFunc{
			defaultHook: func(context.Context) error {
				panic("unexpected invocation of MockReleaseCache.Handle")
			},
		},
		UpdateNowFunc: &ReleaseCacheUpdateNowFunc{
			defaultHook: func(context.Context) error {
				panic("unexpected invocation of MockReleaseCache.UpdateNow")
			},
		},
	}
}

// NewMockReleaseCacheFrom creates a new mock of the MockReleaseCache
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockReleaseCacheFrom(i ReleaseCache) *MockReleaseCache {
	return &MockReleaseCache{
		CurrentFunc: &ReleaseCacheCurrentFunc{
			defaultHook: i.Current,
		},
		HandleFunc: &ReleaseCacheHandleFunc{
			defaultHook: i.Handle,
		},
		UpdateNowFunc: &ReleaseCacheUpdateNowFunc{
			defaultHook: i.UpdateNow,
		},
	}
}

// ReleaseCacheCurrentFunc describes the behavior when the Current method of
// the parent MockReleaseCache instance is invoked.
type ReleaseCacheCurrentFunc struct {
	defaultHook func(string) (string, error)
	hooks       []func(string) (string, error)
	history     []ReleaseCacheCurrentFuncCall
	mutex       sync.Mutex
}

// Current delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockReleaseCache) Current(v0 string) (string, error) {
	r0, r1 := m.CurrentFunc.nextHook()(v0)
	m.CurrentFunc.appendCall(ReleaseCacheCurrentFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Current method of
// the parent MockReleaseCache instance is invoked and the hook queue is
// empty.
func (f *ReleaseCacheCurrentFunc) SetDefaultHook(hook func(string) (string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Current method of the parent MockReleaseCache instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReleaseCacheCurrentFunc) PushHook(hook func(string) (string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *ReleaseCacheCurrentFunc) SetDefaultReturn(r0 string, r1 error) {
	f.SetDefaultHook(func(string) (string, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *ReleaseCacheCurrentFunc) PushReturn(r0 string, r1 error) {
	f.PushHook(func(string) (string, error) {
		return r0, r1
	})
}

func (f *ReleaseCacheCurrentFunc) nextHook() func(string) (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReleaseCacheCurrentFunc) appendCall(r0 ReleaseCacheCurrentFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReleaseCacheCurrentFuncCall objects
// describing the invocations of this function.
func (f *ReleaseCacheCurrentFunc) History() []ReleaseCacheCurrentFuncCall {
	f.mutex.Lock()
	history := make([]ReleaseCacheCurrentFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReleaseCacheCurrentFuncCall is an object that describes an invocation of
// method Current on an instance of MockReleaseCache.
type ReleaseCacheCurrentFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReleaseCacheCurrentFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReleaseCacheCurrentFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// ReleaseCacheHandleFunc describes the behavior when the Handle method of
// the parent MockReleaseCache instance is invoked.
type ReleaseCacheHandleFunc struct {
	defaultHook func(context.Context) error
	hooks       []func(context.Context) error
	history     []ReleaseCacheHandleFuncCall
	mutex       sync.Mutex
}

// Handle delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockReleaseCache) Handle(v0 context.Context) error {
	r0 := m.HandleFunc.nextHook()(v0)
	m.HandleFunc.appendCall(ReleaseCacheHandleFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Handle method of the
// parent MockReleaseCache instance is invoked and the hook queue is empty.
func (f *ReleaseCacheHandleFunc) SetDefaultHook(hook func(context.Context) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Handle method of the parent MockReleaseCache instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReleaseCacheHandleFunc) PushHook(hook func(context.Context) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *ReleaseCacheHandleFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *ReleaseCacheHandleFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context) error {
		return r0
	})
}

func (f *ReleaseCacheHandleFunc) nextHook() func(context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReleaseCacheHandleFunc) appendCall(r0 ReleaseCacheHandleFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReleaseCacheHandleFuncCall objects
// describing the invocations of this function.
func (f *ReleaseCacheHandleFunc) History() []ReleaseCacheHandleFuncCall {
	f.mutex.Lock()
	history := make([]ReleaseCacheHandleFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReleaseCacheHandleFuncCall is an object that describes an invocation of
// method Handle on an instance of MockReleaseCache.
type ReleaseCacheHandleFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReleaseCacheHandleFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReleaseCacheHandleFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ReleaseCacheUpdateNowFunc describes the behavior when the UpdateNow
// method of the parent MockReleaseCache instance is invoked.
type ReleaseCacheUpdateNowFunc struct {
	defaultHook func(context.Context) error
	hooks       []func(context.Context) error
	history     []ReleaseCacheUpdateNowFuncCall
	mutex       sync.Mutex
}

// UpdateNow delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockReleaseCache) UpdateNow(v0 context.Context) error {
	r0 := m.UpdateNowFunc.nextHook()(v0)
	m.UpdateNowFunc.appendCall(ReleaseCacheUpdateNowFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the UpdateNow method of
// the parent MockReleaseCache instance is invoked and the hook queue is
// empty.
func (f *ReleaseCacheUpdateNowFunc) SetDefaultHook(hook func(context.Context) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UpdateNow method of the parent MockReleaseCache instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ReleaseCacheUpdateNowFunc) PushHook(hook func(context.Context) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *ReleaseCacheUpdateNowFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *ReleaseCacheUpdateNowFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context) error {
		return r0
	})
}

func (f *ReleaseCacheUpdateNowFunc) nextHook() func(context.Context) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ReleaseCacheUpdateNowFunc) appendCall(r0 ReleaseCacheUpdateNowFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ReleaseCacheUpdateNowFuncCall objects
// describing the invocations of this function.
func (f *ReleaseCacheUpdateNowFunc) History() []ReleaseCacheUpdateNowFuncCall {
	f.mutex.Lock()
	history := make([]ReleaseCacheUpdateNowFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ReleaseCacheUpdateNowFuncCall is an object that describes an invocation
// of method UpdateNow on an instance of MockReleaseCache.
type ReleaseCacheUpdateNowFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ReleaseCacheUpdateNowFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ReleaseCacheUpdateNowFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
