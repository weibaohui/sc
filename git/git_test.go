// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"bytes"
	"flag"
	"fmt"
	stdlog "log"
	"os"
	"testing"

	goversion "github.com/mcuadros/go-version"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

const repoPath = "testdata/testrepo.git"

var testrepo *Repository

func TestMain(m *testing.M) {
	verbose := flag.Bool("verbose", false, "")
	flag.Parse()

	if *verbose {
		SetOutput(os.Stdout)
	}

	// Set up the test repository
	if !isExist(repoPath) {
		if err := Clone("https://github.com/gogs/git-module-testrepo.git", repoPath, CloneOptions{
			Bare: true,
		}); err != nil {
			stdlog.Fatal(err)
		}
	}

	var err error
	testrepo, err = Open(repoPath)
	if err != nil {
		stdlog.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestSetPrefix(t *testing.T) {
	old := logPrefix
	new := "[custom] "
	SetPrefix(new)
	defer SetPrefix(old)

	assert.Equal(t, new, logPrefix)
}

func Test_log(t *testing.T) {
	old := logOutput
	defer SetOutput(old)

	tests := []struct {
		format    string
		args      []interface{}
		expOutput string
	}{
		{
			format:    "",
			expOutput: "[git-module] \n",
		},
		{
			format:    "something",
			expOutput: "[git-module] something\n",
		},
		{
			format:    "val: %v",
			args:      []interface{}{123},
			expOutput: "[git-module] val: 123\n",
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)

			log(test.format, test.args...)
			assert.Equal(t, test.expOutput, buf.String())
		})
	}
}

func TestBinVersion(t *testing.T) {
	g := errgroup.Group{}
	for i := 0; i < 30; i++ {
		g.Go(func() error {
			version, err := BinVersion()
			assert.Nil(t, err)

			if !goversion.Compare(version, "1.8.3", ">=") {
				return fmt.Errorf("version: expected >= 1.8.3 but got %q", version)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		t.Fatal(err)
	}
}
