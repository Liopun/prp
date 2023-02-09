package prp

type GitRepositoryInput struct {
	RepositoryName string
	Description string
	Private bool
}

type GitBackupInput struct {
	RepositoryName string
	BranchName string
	OwnerID string
	OwnerName string
	OwnerEmail string
	CommitMessage string
	CommitFiles []string
}