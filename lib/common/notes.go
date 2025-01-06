package common

import "strconv"

type Abstract struct {
	Content string
}

func (a Abstract) String() string {
	return "abstract: " + a.Content
}

type Availability struct {
	Content string
}

func (a Availability) String() string {
	return "availability: " + a.Content
}

type Custom struct {
	Level   int
	Content string
}

func (c Custom) String() string {
	return "custom lvl(" + strconv.Itoa(c.Level) + "): " + c.Content
}

type Notes struct {
	Content string
}

func (n Notes) String() string {
	return "notes: " + n.Content
}

type PersonalNotes struct {
	Content string
}

func (n PersonalNotes) String() string {
	return "personal notes: " + n.Content
}

type ResearchNotes struct {
	Content string
}

func (n ResearchNotes) String() string {
	return "research notes: " + n.Content
}

type UserDefinableNotes struct {
	Level   int
	Content string
}

func (n UserDefinableNotes) String() string {
	return "user definable notes lvl(" + strconv.Itoa(n.Level) + "): " + n.Content
}

type Freeform struct {
	Data string
}

func (f Freeform) String() string {
	return "freeform: " + f.Data
}

type Miscellaneous struct {
	Level int
	Data  string
}

func (m Miscellaneous) String() string {
	return "miscellaneous lvl(" + strconv.Itoa(m.Level) + "): " + m.Data
}
