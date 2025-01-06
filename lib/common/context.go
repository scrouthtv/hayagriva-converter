package common

import "strconv"

type ParentJournal struct {
	Name string
}

func (j ParentJournal) String() string {
	return "parent journal: " + j.Name
}

type Volume struct {
	Content string
}

func (v Volume) String() string {
	return "volume: " + v.Content
}

type StartPage struct {
	N int
}

func (s StartPage) String() string {
	return "start page: " + strconv.Itoa(s.N)
}

type EndPage struct {
	N int
}

func (e EndPage) String() string {
	return "end page: " + strconv.Itoa(e.N)
}
