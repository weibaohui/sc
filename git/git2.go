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
	Email       string // 作者邮箱
	Name        string // 作者名称
	CommitCount int    // 提交次数
	Addition    int    // 增加
	Deletion    int    // 删除
}

func (a *AuthorLinesCounter) String() string {
	return fmt.Sprintf("%s(%s) commit count %d,added %d,deleted %d", a.Email, a.Name, a.CommitCount, a.Addition, a.Deletion)
}

type Summary struct {
	Branch          int
	Commit          map[string]int
	AuthorCounts    map[string]*AuthorLinesCounter
	Tags            int
	authorCountsMap *sync.Map             // 并发使用
	authorList      map[string]*Signature // 用户列表
}
type Git struct {
	Summary *Summary
}

// Execute means begin to summary the git repo
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
		g.Summary.Commit[branch] = len(commits)
		utils.CheckIfError(err)

		for _, c := range commits {
			if g.Summary.authorList[c.Author.Email] == nil {
				g.Summary.authorList[c.Author.Email] = c.Author
			}
		}
	}

	concurrency := config.GetInstance().Concurrency
	wp := workpool.New(concurrency)
	for i := range g.Summary.authorList {
		author := g.Summary.authorList[i]
		wp.Do(func() error {
			ac := r.SumAuthor(author)
			g.Summary.authorCountsMap.LoadOrStore(author.Email, ac)
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

// String implement Stringer
func (g *Git) String() string {
	bytes, _ := json.Marshal(g.Summary)
	return string(bytes)
}

// GetInstance return an *Git
func GetInstance() *Git {
	return &Git{
		Summary: summary,
	}
}
func init() {
	once.Do(func() {
		summary = &Summary{
			Branch:          0,
			Commit:          map[string]int{},
			AuthorCounts:    map[string]*AuthorLinesCounter{},
			Tags:            0,
			authorCountsMap: &sync.Map{},
			authorList:      map[string]*Signature{},
		}
	})
}
