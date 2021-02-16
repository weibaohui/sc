// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"errors"
)

var (
	ErrParentNotExist    = errors.New("parent does not exist")
	ErrSubmoduleNotExist = errors.New("submodule does not exist")
	ErrRevisionNotExist  = errors.New("revision does not exist")
	ErrRemoteNotExist    = errors.New("remote does not exist")
	ErrExecTimeout       = errors.New("execution was timed out")
	ErrNoMergeBase       = errors.New("no merge based was found")
	ErrNotBlob           = errors.New("the entry is not a blob")
)
