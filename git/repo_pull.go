// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"strings"
	"time"
)

// MergeBaseOptions contains optional arguments for getting merge base.
// Docs: https://git-scm.com/docs/git-merge-base
type MergeBaseOptions struct {
	// The timeout duration before giving up for each shell command execution.
	// The default timeout duration will be used when not supplied.
	Timeout time.Duration
}

// RepoMergeBase returns merge base between base and head revisions of the repository
// in given path.
func RepoMergeBase(repoPath, base, head string, opts ...MergeBaseOptions) (string, error) {
	var opt MergeBaseOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	stdout, err := NewCommand("merge-base", base, head).RunInDirWithTimeout(opt.Timeout, repoPath)
	if err != nil {
		if strings.Contains(err.Error(), "exit status 1") {
			return "", ErrNoMergeBase
		}
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}

// MergeBase returns merge base between base and head revisions of the repository.
func (r *Repository) MergeBase(base, head string, opts ...MergeBaseOptions) (string, error) {
	return RepoMergeBase(r.path, base, head, opts...)
}
