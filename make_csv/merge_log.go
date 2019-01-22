package make_csv

import (
	"github.com/Ryu955/shintyoku/collect"
)

type Log struct {
	Name string
	List collect.List
}
type MergeLog []Log

func AddLog(name string, list collect.List, merge MergeLog) MergeLog {
	log := Log{name, list}
	merge = append(merge, log)
	return merge
}
