package git

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/xxjwxc/gowp/workpool"

	"github.com/weibaohui/sc/config"
	"github.com/weibaohui/sc/utils"
)

var once = sync.Once{}
var summary *Summary

func checkOut() {
	r, err := Open(".")
	utils.CheckIfError(err)
	r.Checkout("git")
}

type AuthorLinesCounters map[string]*AuthorLinesCounter
type AuthorLinesCounter struct {
	Author      string
	CommitCount int // 提交次数
	Addition    int // 增加
	Deletion    int // 删除
}

func (a *AuthorLinesCounter) String() string {
	return fmt.Sprintf("%s commit count %d,added %d,deleted %d", a.Author, a.CommitCount, a.Addition, a.Deletion)
}

type Summary struct {
	Branch          int
	Commit          int
	AuthorCounts    map[string]*AuthorLinesCounter
	Tags            int
	authorCountsMap *sync.Map         // 并发使用
	authorList      map[string]string // 用户列表
}
type Git struct {
	Summary *Summary
}

func (g *Git) Execute() *Git {
	r, err := Open(".")
	branches, err := r.Branches()
	utils.CheckIfError(err)
	g.Summary.Branch = len(branches)

	for _, branch := range branches {
		id, err := r.BranchCommitID(branch)
		utils.CheckIfError(err)

		commits, err := r.Log(id)
		// todo 取最大值 or 每个分支的数值
		g.Summary.Commit = len(commits)
		utils.CheckIfError(err)

		for _, c := range commits {
			if g.Summary.authorList[c.Author.Name] == "" {
				g.Summary.authorList[c.Author.Name] = c.Author.Name
			}
		}
	}

	concurrency := config.GetInstance().Concurrency
	wp := workpool.New(concurrency)
	for i := range g.Summary.authorList {
		name := g.Summary.authorList[i]
		wp.Do(func() error {
			ac := r.SumAuthor(name)
			g.Summary.authorCountsMap.LoadOrStore(name, ac)
			return nil
		})
	}
	err = wp.Wait()
	utils.CheckIfError(err)
	g.Summary.authorCountsMap.Range(func(k, v interface{}) bool {
		g.Summary.AuthorCounts[k.(string)] = v.(*AuthorLinesCounter)
		return true
	})

	return g
}

func (g *Git) String() string {
	bytes, _ := json.Marshal(g.Summary)
	return string(bytes)
}

func GetInstance() *Git {
	return &Git{
		Summary: summary,
	}
}
func init() {
	once.Do(func() {
		summary = &Summary{
			Branch:          0,
			Commit:          0,
			AuthorCounts:    map[string]*AuthorLinesCounter{},
			Tags:            0,
			authorCountsMap: &sync.Map{},
			authorList:      map[string]string{},
		}
	})
}
