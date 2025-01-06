package common

type AccessionNumber struct {
	Number string
}

func (a AccessionNumber) String() string {
	return "accession number: " + a.Number
}

type DOI struct {
	Number string
}

func (d DOI) String() string {
	return "doi: " + d.Number
}

type PMID struct {
	Number string
}

func (p PMID) String() string {
	return "pmid: " + p.Number
}

type PMCID struct {
	Number string
}

func (p PMCID) String() string {
	return "pmcid: " + p.Number
}

type SN struct {
	Number string
}

func (s SN) String() string {
	return "sn: " + s.Number
}

type Web struct {
	Content string
}

func (w Web) String() string {
	return "web: " + w.Content
}

type Language struct {
	Code string
}

func (l Language) String() string {
	return "language: " + l.Code
}

type Classification struct {
	Content string
}

func (c Classification) String() string {
	return "classification: " + c.Content
}

type CallNumber struct {
	Content string
}

func (c CallNumber) String() string {
	return "call number: " + c.Content
}
