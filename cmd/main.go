package main

import (
	"fmt"
	"github.com/Ryu955/shintyoku/collect"
	"github.com/Ryu955/shintyoku/make_csv"
	"github.com/Ryu955/shintyoku/setting"
)

func main() {
	repos := setting.GetRepoName()

	var logs make_csv.MergeLog

	for _, repo := range repos {
		fmt.Println(repo)
		logs = make_csv.AddLog(repo, collect.GetGitLog(repo), logs)
		fmt.Println(collect.GetGitLog(repo))
		var x collect.List
		x = collect.GetGitLog(repo)
		fmt.Println(x)
	}

	fmt.Println(logs[1])
	fmt.Println(logs[1].List[1])
	fmt.Println(logs[1].List[1].Date)
	fmt.Println(logs[1].List[1].Count)

	fmt.Println(logs[1].List[0])
	fmt.Println(logs[1].List[0].Date)
	fmt.Println(logs[1].List[0].Count)

	//a := collect.GetGitLog("201801_GraduationThesis_ryutai")
	//fmt.Println(a)
}
