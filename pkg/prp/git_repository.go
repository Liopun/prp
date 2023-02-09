package prp

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v50/github"
)

type GitRepo interface {
	AddGitPrivateRepo(ctx context.Context, inp GitRepositoryInput) (string, error)
	GetGitRef(ctx context.Context, inp GitBackupInput) (ref *github.Reference, err error)
	GetGitTree(ctx context.Context, ref *github.Reference, inp GitBackupInput) (tree *github.Tree, err error)
	PushGitCommit(ctx context.Context, ref *github.Reference, tree *github.Tree, inp GitBackupInput) error
}

type GitRepository struct {
	client *github.Client
}

func NewGhRepo(client *github.Client) *GitRepository {
	return &GitRepository{
		client: client,
	}
}

func (r *GitRepository) AddGitPrivateRepo(ctx context.Context, inp GitRepositoryInput) (string, error) {
	gitRepo := &github.Repository{
		Name: &inp.RepositoryName,
		Private: &inp.Private,
		Description: &inp.Description,
		AutoInit: github.Bool(true),
	}

	repo, _, err := r.client.Repositories.Create(ctx, "", gitRepo)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("new repo created: %s", repo.GetName()), nil
}

func (r *GitRepository) GetGitRef(ctx context.Context, inp GitBackupInput) (ref *github.Reference, err error) {
	if ref, _, err = r.client.Git.GetRef(ctx, inp.OwnerID, inp.RepositoryName, "refs/heads/"+inp.BranchName); err == nil {
		return ref, nil
	}

	var baseRef *github.Reference
	if baseRef, _, err = r.client.Git.GetRef(ctx, inp.OwnerID, inp.RepositoryName, "refs/heads/"+inp.BranchName); err != nil {
		return nil, err
	}

	newRef := &github.Reference{Ref: github.String("refs/heads/" + inp.BranchName), Object: &github.GitObject{SHA: baseRef.Object.SHA}}
	ref, _, err = r.client.Git.CreateRef(ctx, inp.OwnerID, inp.RepositoryName, newRef)
	return ref, err
}

func (r *GitRepository) GetGitTree(ctx context.Context, ref *github.Reference, inp GitBackupInput) (tree *github.Tree, err error) {
	// Create a tree with what to commit.
	entries := []*github.TreeEntry{}

	// Load each file into the tree.
	for _, fileArg := range inp.CommitFiles {
		file, content, err := getFileContent(fileArg)
		if err != nil {
			return nil, err
		}
		entries = append(entries, &github.TreeEntry{Path: github.String(file), Type: github.String("blob"), Content: github.String(string(content)), Mode: github.String("100644")})
	}

	tree, _, err = r.client.Git.CreateTree(ctx, inp.OwnerID, inp.RepositoryName, *ref.Object.SHA, entries)
	return tree, err
}

func (r *GitRepository) PushGitCommit(ctx context.Context, ref *github.Reference, tree *github.Tree, inp GitBackupInput) error {
	// Get the parent commit to attach the commit to.
	parent, _, err := r.client.Repositories.GetCommit(ctx, inp.OwnerID, inp.RepositoryName, *ref.Object.SHA, nil)
	if err != nil {
		return err
	}
	// This is not always populated, but is needed.
	parent.Commit.SHA = parent.SHA

	// Create the commit using the tree.
	date := time.Now()
	author := &github.CommitAuthor{Date: &github.Timestamp{Time: date}, Name: &inp.OwnerName, Email: &inp.OwnerEmail}
	commit := &github.Commit{Author: author, Message: &inp.CommitMessage, Tree: tree, Parents: []*github.Commit{parent.Commit}}
	newCommit, _, err := r.client.Git.CreateCommit(ctx, inp.OwnerID, inp.RepositoryName, commit)
	if err != nil {
		return err
	}

	// Attach the commit to the master branch.
	ref.Object.SHA = newCommit.SHA
	_, _, err = r.client.Git.UpdateRef(ctx, inp.OwnerID, inp.RepositoryName, ref, false)
	return err
}
