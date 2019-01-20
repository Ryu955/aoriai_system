package main

import (
	"./collect"
	"fmt"
)

func main() {
	a := collect.GetGitLog("201801_GraduationThesis_ryutai")
	fmt.Print(a)
}
