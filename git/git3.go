package git

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-git/go-git/v5"

	"github.com/weibaohui/sc/utils"
)

type Git struct {
}

var counts = map[string]*AuthorLinesCounter{}
var counts_sync = sync.Map{}

func GetInstance() *Git {
	r, err := git.PlainOpen(".")
	utils.CheckIfError(err)
	// Length of the HEAD history
	utils.Info("git rev-list HEAD --count")

	// ... retrieving the HEAD reference
	ref, err := r.Head()
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.CheckIfError(err)

	// ... retrieves the commit history
	// since := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	// until := time.Date(2021, 2, 16, 0, 0, 0, 0, time.UTC)
	// cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	utils.CheckIfError(err)
	defer cIter.Close()
	// ... just iterates over the commits
	var cCount int
	ticker := time.NewTicker(1 * time.Second)

	go func() {

		for {
			select {
			case <-ticker.C:
				// fmt.Println(len(counts), cCount)
			}
		}
	}()

	for {
		c, err := cIter.Next()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		cCount++
		var r *AuthorLinesCounter
		v, ok := counts_sync.Load(c.Author.Name)
		if ok {
			r = v.(*AuthorLinesCounter)
			r.CommitCount += 1
		} else {
			r = &AuthorLinesCounter{
				Author:      c.Author.Name,
				CommitCount: 0,
				Addition:    0,
				Deletion:    0,
			}
			counts_sync.Store(c.Author.Name, r)
			// fmt.Println("sync.map store",c.Author.Name)
		}

		// fmt.Printf("%s,%s\r", c.Author.String(), c.Message)

		fileStats, _ := c.Stats()
		for _, stat := range fileStats {
			// fmt.Println(stat.Name, stat.Addition, stat.Deletion)
			v, ok := counts_sync.Load(c.Author.Name)
			if ok {
				r1 := v.(*AuthorLinesCounter)
				r1.Addition += stat.Addition
				r1.Deletion += stat.Deletion
			}

		}

	}

	utils.CheckIfError(err)
	fmt.Println("allcount-", cCount)
	// for _, stat := range stats {
	// 	fmt.Println(stat.String())
	// }

	return nil
}
