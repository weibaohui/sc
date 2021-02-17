// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

// Commit contains information of a Git commit.
type Commit struct {
	//  The author of the commit.
	Author *Signature
	// The committer of the commit.
	Committer *Signature
	// The full commit message.
	Message string
}
