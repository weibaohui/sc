package git

import (
	git "github.com/gogs/git-module"

	"github.com/weibaohui/sc/utils"
)

func InitGitModule() {
	_, err := git.Open(".")
	utils.CheckIfError(err)
}
