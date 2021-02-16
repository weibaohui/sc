// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRefShortName(t *testing.T) {
	tests := []struct {
		ref    string
		expVal string
	}{
		{
			ref:    "refs/heads/master",
			expVal: "master",
		},
		{
			ref:    "refs/tags/v1.0.0",
			expVal: "v1.0.0",
		},
		{
			ref:    "refs/pull/98",
			expVal: "refs/pull/98",
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, RefShortName(test.ref))
		})
	}
}

func TestRepository_ShowRefVerify(t *testing.T) {
	t.Run("reference does not exsit", func(t *testing.T) {
		rev, err := testrepo.ShowRefVerify("bad_reference")
		assert.NotNil(t, err)
		assert.Empty(t, rev)
	})

	rev, err := testrepo.ShowRefVerify("refs/heads/release-1.0")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "0eedd79eba4394bbef888c804e899731644367fe", rev)
}

func TestRepository_BranchCommitID(t *testing.T) {
	t.Run("branch does not exsit", func(t *testing.T) {
		rev, err := testrepo.BranchCommitID("bad_branch")
		assert.NotNil(t, err)
		assert.Empty(t, rev)
	})

	rev, err := testrepo.BranchCommitID("release-1.0")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "0eedd79eba4394bbef888c804e899731644367fe", rev)
}

func TestRepository_TagCommitID(t *testing.T) {
	t.Run("tag does not exsit", func(t *testing.T) {
		rev, err := testrepo.TagCommitID("bad_tag")
		assert.NotNil(t, err)
		assert.Empty(t, rev)
	})

	rev, err := testrepo.TagCommitID("v1.0.0")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "0eedd79eba4394bbef888c804e899731644367fe", rev)
}

func TestRepository_HasReference(t *testing.T) {
	tests := []struct {
		ref    string
		opt    ShowRefVerifyOptions
		expVal bool
	}{
		{
			ref:    RefsHeads + "master",
			expVal: true,
		},
		{
			ref:    RefsTags + "v1.0.0",
			expVal: true,
		},
		{
			ref:    "master",
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, testrepo.HasReference(test.ref, test.opt))
		})
	}
}

func TestRepository_HasBranch(t *testing.T) {
	tests := []struct {
		ref    string
		opt    ShowRefVerifyOptions
		expVal bool
	}{
		{
			ref:    "master",
			expVal: true,
		},
		{
			ref:    RefsHeads + "master",
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, testrepo.HasBranch(test.ref, test.opt))
		})
	}
}

func TestRepository_HasTag(t *testing.T) {
	tests := []struct {
		ref    string
		opt    ShowRefVerifyOptions
		expVal bool
	}{
		{
			ref:    "v1.0.0",
			expVal: true,
		},
		{
			ref:    RefsTags + "v1.0.0",
			expVal: false,
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expVal, testrepo.HasTag(test.ref, test.opt))
		})
	}
}

func TestRepository_SymbolicRef(t *testing.T) {
	r, cleanup, err := setupTempRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	// Get HEAD
	ref, err := r.SymbolicRef()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, RefsHeads+"master", ref)

	// Set a symbolic reference
	_, err = r.SymbolicRef(SymbolicRefOptions{
		Name: "TEST-REF",
		Ref:  RefsHeads + "develop",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Get the symbolic reference we just set
	ref, err = r.SymbolicRef(SymbolicRefOptions{
		Name: "TEST-REF",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, RefsHeads+"develop", ref)
}

func TestRepository_ShowRef(t *testing.T) {
	tests := []struct {
		opt     ShowRefOptions
		expRefs []*Reference
	}{
		{
			opt: ShowRefOptions{
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
			opt: ShowRefOptions{
				Tags:     true,
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
			refs, err := testrepo.ShowRef(test.opt)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.expRefs, refs)
		})
	}
}

func TestRepository_Branches(t *testing.T) {
	expBranches := map[string]bool{
		"master":      true,
		"develop":     true,
		"release-1.0": true,
	}
	branches, err := testrepo.Branches()
	if err != nil {
		t.Fatal(err)
	}
	for _, b := range branches {
		delete(expBranches, b)
	}

	if len(expBranches) > 0 {
		t.Fatalf("expect to be empty but got %v", expBranches)
	}
}

func TestRepository_DeleteBranch(t *testing.T) {
	r, cleanup, err := setupTempRepo()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	tests := []struct {
		opt DeleteBranchOptions
	}{
		{
			opt: DeleteBranchOptions{
				Force: false,
			},
		},
		{
			opt: DeleteBranchOptions{
				Force: true,
			},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			branch := strconv.Itoa(int(time.Now().UnixNano()))
			err := r.Checkout(branch, CheckoutOptions{
				BaseBranch: "master",
			})
			if err != nil {
				t.Fatal(err)
			}

			assert.True(t, r.HasReference(RefsHeads+branch))

			err = r.Checkout("master")
			if err != nil {
				t.Fatal(err)
			}

			err = r.DeleteBranch(branch, test.opt)
			if err != nil {
				t.Fatal(err)
			}

			assert.False(t, r.HasReference(RefsHeads+branch))
		})
	}
}
