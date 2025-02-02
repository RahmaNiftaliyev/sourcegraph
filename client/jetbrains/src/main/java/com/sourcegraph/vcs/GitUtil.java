package com.sourcegraph.vcs;

import com.intellij.dvcs.repo.Repository;
import com.intellij.openapi.project.Project;
import git4idea.GitRemoteBranch;
import git4idea.repo.GitRemote;
import git4idea.repo.GitRepository;
import git4idea.repo.GitRepositoryManager;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.stream.Collectors;

public class GitUtil {

    @NotNull
    // Returned format: git@github.com:sourcegraph/sourcegraph.git
    public static String getRemoteRepoUrl(@NotNull GitRepository repository, @NotNull Project project) throws Exception {
        GitRemote remote = getBestRemote(repository, project);
        if (remote == null) {
            throw new Exception("No configured git remote for \"sourcegraph\" or \"origin\".");
        }
        String url = remote.getFirstUrl();
        if (url == null) {
            throw new Exception("No URL found for git remote \"" + remote.getName() + "\".");
        }
        return url;
    }

    /**
     * @param branchName E.g. "main"
     */
    public static boolean doesRemoteBranchExist(@NotNull GitRepository repository, @NotNull String branchName) {
        List<String> remoteBranchNames = repository.getInfo().getRemoteBranchesWithHashes()
            .keySet().stream().map(GitRemoteBranch::getNameForLocalOperations).collect(Collectors.toList());
        return remoteBranchNames.stream().anyMatch(name -> name.equals(branchName));
    }

    @Nullable
    private static GitRemote getBestRemote(@NotNull GitRepository repository, @NotNull Project project) {
        GitRemote sourcegraphRemote = getRemote(repository, project, "sourcegraph");
        if (sourcegraphRemote != null) {
            return sourcegraphRemote;
        }
        return getRemote(repository, project, "origin");
    }

    @Nullable
    private static GitRemote getRemote(@NotNull Repository repository, @NotNull Project project, @NotNull String remoteName) {
        GitRepository gitRepository = GitRepositoryManager.getInstance(project).getRepositoryForRoot(repository.getRoot());
        if (gitRepository == null) {
            return null;
        }

        List<GitRemote> matchingRemotes = gitRepository.getRemotes().stream().filter(x -> x.getName().equals(remoteName)).collect(Collectors.toList());
        try {
            return matchingRemotes.get(0);
        } catch (IndexOutOfBoundsException e) {
            return null;
        }
    }
}
