// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLsRemote(t *testing.T) {
	tests := []struct {
		url     string
		opt     LsRemoteOptions
		expRefs []*Reference
	}{
		{
			url: testrepo.Path(),
			opt: LsRemoteOptions{
				Heads:    true,
				Patterns: []string{"release-1.0"},
			},
			expRefs: []*Reference{
				{
					ID:      "0eedd79eba4394bbef888c804e899731644367fe",
					Refspec: "refs/heads/release-1.0",
				},
			},
		}, {
			url: testrepo.Path(),
			opt: LsRemoteOptions{
				Tags:     true,
				Patterns: []string{"v1.0.0"},
			},
			expRefs: []*Reference{
				{
					ID:      "0eedd79eba4394bbef888c804e899731644367fe",
					Refspec: "refs/tags/v1.0.0",
				},
			},
		}, {
			url: testrepo.Path(),
			opt: LsRemoteOptions{
				Refs:     true,
				Patterns: []string{"v1.0.0"},
			},
			expRefs: []*Reference{
				{
					ID:      "0eedd79eba4394bbef888c804e899731644367fe",
					Refspec: "refs/tags/v1.0.0",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			refs, err := LsRemote(test.url, test.opt)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.expRefs, refs)
		})
	}
}

func TestIsURLAccessible(t *testing.T) {
	tests := []struct {
		url    string
		expVal bool
	}{
		{
			url:    testrepo.Path(),
			expVal: true,
		}, {
			url:    os.TempDir(),
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, IsURLAccessible(DefaultTimeout, test.url))
		})
	}
}

func TestRepository_AddRemote(t *testing.T) {
	path := tempPath()
	defer func() {
		_ = os.RemoveAll(path)
	}()

	err := Init(path, InitOptions{
		Bare: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	r, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}

	// Addition testrepo as the mirror remote and fetch right away
	err = r.AddRemote("origin", testrepo.Path(), AddRemoteOptions{
		Fetch:       true,
		MirrorFetch: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Check a non-default branch: release-1.0
	assert.True(t, r.HasReference(RefsHeads+"release-1.0"))
}

func TestRepository_RemoveRemote(t *testing.T) {
	r, cleanup, err := setupTempRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	err = r.RemoveRemote("origin", RemoveRemoteOptions{})
	assert.Nil(t, err)

	err = r.RemoveRemote("origin", RemoveRemoteOptions{})
	assert.Equal(t, ErrRemoteNotExist, err)
}
