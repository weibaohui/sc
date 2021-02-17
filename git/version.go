// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"fmt"
	"strings"
	"sync"
)

var (
	// gitVersion stores the Git binary version.
	// NOTE: To check Git version should call BinVersion not this global variable.
	gitVersion     string
	gitVersionOnce sync.Once
	gitVersionErr  error
)

// BinVersion returns current Git binary version that is used by this module.
func BinVersion() (string, error) {
	gitVersionOnce.Do(func() {
		var stdout []byte
		stdout, gitVersionErr = NewCommand("version").Run()
		if gitVersionErr != nil {
			return
		}

		fields := strings.Fields(string(stdout))
		if len(fields) < 3 {
			gitVersionErr = fmt.Errorf("not enough output: %s", stdout)
			return
		}

		// Handle special case on Windows.
		i := strings.Index(fields[2], "windows")
		if i >= 1 {
			gitVersion = fields[2][:i-1]
			return
		}

		gitVersion = fields[2]
	})

	return gitVersion, gitVersionErr
}
