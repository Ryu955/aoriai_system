package main

import (
	"fmt"
	"github.com/Ryu955/shintyoku/collect"
	"github.com/Ryu955/shintyoku/setting"
)

func main() {
	repos := setting.GetRepoName()

	for _, repo := range repos {
		fmt.Println(repo)
		fmt.Println(collect.GetGitLog(repo))
	}
}
