package main

import (
	"fmt"
	"github.com/Ryu955/shintyoku/collect"
	"github.com/Ryu955/shintyoku/setting"
)

func main() {
	setting.GetRepoName()
	a := collect.GetGitLog("201801_GraduationThesis_ryutai")
	fmt.Println(a)
}
