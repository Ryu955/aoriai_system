package collect

type Log struct {
	Date  string
	Count int
}
type List []Log

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].Date == l[j].Date {
		return (l[i].Count < l[j].Count)
	} else {
		return (l[i].Date < l[j].Date)
	}
}
