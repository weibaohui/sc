package git

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/weibaohui/sc/config"
	"github.com/weibaohui/sc/utils"
)

var once = sync.Once{}
var summary *Summary

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
	Tags            int
	Commit          map[string]int
	AuthorCounts    map[string]*AuthorLinesCounter
	authorList      map[string]*Signature // 用户列表
	authorCountsMap *sync.Map             // 并发使用
	CurrentBranch   string                // 当前分支
}
type Git struct {
	Summary *Summary
	repo    *Repository
}

// Execute means begin to summary the git repo
func (g *Git) GoExecute() *Git {
	channel := GetChanInstance()
	timer := time.NewTicker(time.Second * 1)

	// 列表 commit ,统计作者
	go func() {

		ref, err := g.repo.SymbolicRef()
		utils.CheckIfError(err)
		g.Summary.CurrentBranch = RefShortName(ref)
		tags, err := g.repo.Tags()
		utils.CheckIfError(err)
		g.Summary.Tags = len(tags)
		branches, err := g.repo.Branches()
		utils.CheckIfError(err)
		g.Summary.Branch = len(branches)

		for _, branch := range branches {
			id, err := g.repo.BranchCommitID(branch)
			utils.CheckIfError(err)

			count, err := g.repo.LogGo(id)
			g.Summary.Commit[branch] = count
			utils.CheckIfError(err)
		}
	}()

	// 按作者统计 代码量
	go func() {
		for {
			select {
			case c := <-commitChan.AuthorEmail:
				_, exists := g.Summary.authorCountsMap.Load(c.Author.Email)
				if !exists {
					Debugf("统计作者%s\n", c.Author.Email)
					ac := g.repo.SumAuthor(c.Author)
					if ac != nil {
						g.Summary.authorCountsMap.Store(c.Author.Email, ac)
					}
					// debug 查看
					if config.GetInstance().Debug {
						if v, ok := g.Summary.authorCountsMap.Load(c.Author.Email); ok && v != nil {
							tmp := v.(*AuthorLinesCounter)
							if tmp != nil {
								Debugf("%s[%s]:提交%d次,%d+,%d-\n", tmp.Name, tmp.Email, tmp.CommitCount, tmp.Addition, tmp.Deletion)
							}
						}
					}

				}
				channel.Process(c)
			case <-timer.C:
				Debugf("收到%d，完成%d\n", channel.receiveCount.Load(), channel.processCount.Load())
				if channel.IsDone() {
					Debugf("全部统计完毕-%d\n", channel.processCount)
					channel.Complete()
					return
				}

			}
		}
	}()

	// 等待统计结束
	for {
		select {
		case <-channel.done:
			Debug("统计结束")
			g.Summary.authorCountsMap.Range(func(k, v interface{}) bool {
				g.Summary.AuthorCounts[k.(string)] = v.(*AuthorLinesCounter)
				return true
			})
			return g
		}
	}

}

// String implement Stringer
func (g *Git) String() string {
	bytes, _ := json.Marshal(g.Summary)
	return string(bytes)
}

// Result contains Summary info
func (g *Git) Result() *Summary {
	return g.Summary
}

// GetInstance return an *Git
func GetInstance() *Git {
	path := config.GetInstance().InitPath
	r, err := Open(path)
	utils.CheckIfError(err)

	return &Git{
		Summary: summary,
		repo:    r,
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
