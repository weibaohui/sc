// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

// Tree represents a flat directory listing in Git.
type Tree struct {
	id   *SHA1
	repo *Repository
}
