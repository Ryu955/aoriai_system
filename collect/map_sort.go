package collect

type Log struct {
	date string
	value int
}
type List []Log

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].date< l[j].date)
	} else {
		return (l[i].value < l[j].value)
	}
}