package git

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"

	"sc/utils"
)

type Git struct {
}

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
	since := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2021, 2, 16, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	utils.CheckIfError(err)
	defer cIter.Close()
	// ... just iterates over the commits
	var cCount int
	var stats []*object.FileStat = make([]*object.FileStat, 0, 1000000)
	ticker := time.NewTicker(1 * time.Second)

	go func() {

		for {
			select {
			case <-ticker.C:
				fmt.Println(len(stats), cCount)

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

		// fmt.Printf("%s,%s\r", c.Author.String(), c.Message)
		fileStats, _ := c.Stats()
		for _, stat := range fileStats {
			stats = append(stats, &stat)
			// fmt.Println(stat.Name,stat.Addition,stat.Deletion)
		}
	}

	utils.CheckIfError(err)
	fmt.Println("allcount-", cCount, len(stats))
	// for _, stat := range stats {
	// 	fmt.Println(stat.String())
	// }
	time.Sleep(10 * time.Second)
	return nil
}
