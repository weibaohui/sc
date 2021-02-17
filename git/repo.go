// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/xxjwxc/gowp/workpool"

	"github.com/weibaohui/sc/config"
	"github.com/weibaohui/sc/utils"
)

// Repository contains information of a Git repository.
type Repository struct {
	path                 string
	syncMapCachedCommits *sync.Map
	syncMapCachedTags    *sync.Map
}

// Path returns the path of the repository.
func (r *Repository) Path() string {
	return r.path
}

const LogFormatHashOnly = `format:%H`

// parsePrettyFormatLogToList returns a list of commits parsed from given logs that are
// formatted in LogFormatHashOnly.
func (r *Repository) parsePrettyFormatLogToList(timeout time.Duration, logs []byte) ([]*Commit, error) {
	if len(logs) == 0 {
		return []*Commit{}, nil
	}
	ids := bytes.Split(logs, []byte{'\n'})
	commits := make([]*Commit, 0, len(ids))
	for i := 0; i < len(ids); i++ {
		v, _ := r.CatFileCommit(string(ids[i]), CatFileCommitOptions{Timeout: timeout})
		commits = append(commits, v)
	}
	return commits, nil
}

// parsePrettyFormatLogToList returns a list of commits parsed from given logs that are
// formatted in LogFormatHashOnly.
func (r *Repository) parsePrettyFormatLogToListGo(timeout time.Duration, logs []byte) error {
	if len(logs) == 0 {
		return nil
	}
	ids := bytes.Split(logs, []byte{'\n'})
	concurrency := config.GetInstance().Concurrency
	wp := workpool.New(concurrency * 3)

	for i := 0; i < len(ids); i++ {
		id := ids[i]
		wp.Do(func() error {
			v, _ := r.CatFileCommit(string(id), CatFileCommitOptions{Timeout: timeout})
			if v != nil {
				GetChanInstance().Add(v)
			}
			return nil
		})
	}
	err := wp.Wait()
	utils.CheckIfError(err)
	return nil
}

// Open opens the repository at the given path. It returns an os.ErrNotExist
// if the path does not exist.
func Open(repoPath string) (*Repository, error) {
	repoPath, err := filepath.Abs(repoPath)
	if err != nil {
		return nil, err
	} else if !utils.IsDir(repoPath) {
		return nil, os.ErrNotExist
	}

	return &Repository{
		path:                 repoPath,
		syncMapCachedCommits: &sync.Map{},
		syncMapCachedTags:    &sync.Map{},
	}, nil
}

// CheckoutOptions contains optional arguments for checking out to a branch.
// Docs: https://git-scm.com/docs/git-checkout
type CheckoutOptions struct {
	// The base branch if checks out to a new branch.
	BaseBranch string
	// The timeout duration before giving up for each shell command execution.
	// The default timeout duration will be used when not supplied.
	Timeout time.Duration
}

// Checkout checks out to given branch for the repository in given path.
func RepoCheckout(repoPath, branch string, opts ...CheckoutOptions) error {
	var opt CheckoutOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	cmd := NewCommand("checkout")
	if opt.BaseBranch != "" {
		cmd.AddArgs("-b")
	}
	cmd.AddArgs(branch)
	if opt.BaseBranch != "" {
		cmd.AddArgs(opt.BaseBranch)
	}

	_, err := cmd.RunInDirWithTimeout(opt.Timeout, repoPath)
	return err
}

// Checkout checks out to given branch for the repository.
func (r *Repository) Checkout(branch string, opts ...CheckoutOptions) error {
	return RepoCheckout(r.path, branch, opts...)
}

// CommitOptions contains optional arguments to commit changes.
// Docs: https://git-scm.com/docs/git-commit
type CommitOptions struct {
	// Author is the author of the changes if that's not the same as committer.
	Author *Signature
	// The timeout duration before giving up for each shell command execution.
	// The default timeout duration will be used when not supplied.
	Timeout time.Duration
}

// RepoCommit commits local changes with given author, committer and message for the
// repository in given path.
func RepoCommit(repoPath string, committer *Signature, message string, opts ...CommitOptions) error {
	var opt CommitOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	cmd := NewCommand("commit")
	cmd.AddEnvs("GIT_COMMITTER_NAME="+committer.Name, "GIT_COMMITTER_EMAIL="+committer.Email)

	if opt.Author == nil {
		opt.Author = committer
	}
	cmd.AddArgs(fmt.Sprintf("--author='%s <%s>'", opt.Author.Name, opt.Author.Email))
	cmd.AddArgs("-m", message)

	_, err := cmd.RunInDirWithTimeout(opt.Timeout, repoPath)
	// No stderr but exit status 1 means nothing to commit.
	if err != nil && err.Error() == "exit status 1" {
		return nil
	}
	return err
}

// Commit commits local changes with given author, committer and message for the repository.
func (r *Repository) Commit(committer *Signature, message string, opts ...CommitOptions) error {
	return RepoCommit(r.path, committer, message, opts...)
}

// RevParseOptions contains optional arguments for parsing revision.
// Docs: https://git-scm.com/docs/git-rev-parse
type RevParseOptions struct {
	// The timeout duration before giving up for each shell command execution.
	// The default timeout duration will be used when not supplied.
	Timeout time.Duration
}

// RevParse returns full length (40) commit ID by given revision in the repository.
func (r *Repository) RevParse(rev string, opts ...RevParseOptions) (string, error) {
	var opt RevParseOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	commitID, err := NewCommand("rev-parse", rev).RunInDirWithTimeout(opt.Timeout, r.path)
	if err != nil {
		if strings.Contains(err.Error(), "exit status 128") {
			return "", ErrRevisionNotExist
		}
		return "", err
	}
	return strings.TrimSpace(string(commitID)), nil
}
