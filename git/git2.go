package git

import (
	"encoding/json"
	"fmt"

	"github.com/weibaohui/sc/utils"
)

func InitGitModule() {
	r, err := Open(".")

	branches, err := r.Branches()
	fmt.Println("共有分支", len(branches))
	utils.CheckIfError(err)
	fmt.Println("搜索所有作者开始")

	for _, branch := range branches {
		id, err := r.BranchCommitID(branch)
		utils.CheckIfError(err)

		fmt.Println("搜索Log开始")
		commits, err := r.Log(id)
		fmt.Println("分支", branch, id, "共有提交", len(commits))

		utils.CheckIfError(err)
		fmt.Println("遍历commits开始")
		for _, c := range commits {
			if counts[c.Author.Name] == nil {
				ac := &AuthorLinesCounter{
					Author:      c.Author.Name,
					CommitCount: 0,
					Addition:    0,
					Deletion:    0,
				}
				counts[c.Author.Name] = ac
			}
			counts[c.Author.Name].CommitCount += 1
		}
		fmt.Println("遍历commits结束")
	}
	fmt.Println("搜索所有作者结束")
	for name, _ := range counts {

		ac := r.SumAuthor(name)
		// fmt.Println(ac)
		counts[name] = ac

	}

	byts, _ := json.Marshal(counts)
	fmt.Println(string(byts))
}
