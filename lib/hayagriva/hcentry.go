package hayagriva

import (
	"fmt"
	"log"
	"time"

	"github.com/scrouthtv/hayagriva-converter/lib/common"
)

type HcEntry struct {
	Type            string
	Title           string
	Authors         []string    `yaml:"author"`
	Date            Date        `yaml:",omitempty"`
	Parents         []HcEntry   `yaml:"parent,omitempty"`
	Abstract        string      `yaml:",omitempty"`
	Genre           string      `yaml:",omitempty"`
	Editors         []string    `yaml:"editor,omitempty"`
	Affiliated      []Person    `yaml:",omitempty"`
	CallNumber      string      `yaml:"call-number,omitempty"`
	Publishers      []Publisher `yaml:"publisher,omitempty"`
	Location        string      `yaml:",omitempty"`
	Organization    string      `yaml:",omitempty"`
	Issue           string      `yaml:",omitempty"`
	Volume          string      `yaml:",omitempty"`
	VolumeTotal     int         `yaml:"volume-total,omitempty"`
	Edition         string      `yaml:",omitempty"`
	PageRange       PageRange   `yaml:"page-range,omitempty"`
	PageTotal       int         `yaml:"page-total,omitempty"`
	TimeRange       string      `yaml:"time-range,omitempty"`
	Timestamp       string      `yaml:",omitempty"`
	Url             Url         `yaml:",omitempty"`
	SerialNumber    Serial      `yaml:"serial-number,omitempty"`
	Language        string      `yaml:",omitempty"`
	Archive         string      `yaml:",omitempty"`
	ArchiveLocation string      `yaml:",omitempty"`
	Note            string      `yaml:",omitempty"`
}

func ToHcEntry(src common.Entry) (HcEntry, error) {
	var dst HcEntry

	err := dst.SetType(src.Type)
	if err != nil {
		return dst, err
	}

	for _, i := range src.Information {
		err = dst.setInformation(i)
		if err != nil {
			return dst, err
		}
	}

	return dst, nil
}

func (e *HcEntry) setInformation(i common.Member) error {
	switch v := i.(type) {
	case common.Author:
		e.Authors = append(e.Authors, v.AuthorName())
	case common.ParentJournal:
		var parent HcEntry
		parent.Type = "periodical"
		parent.Title = v.Name
		e.Parents = append(e.Parents, parent)
	case common.Volume:
		e.Volume = v.Content
	case common.StartPage:
		e.PageRange.Start = v.N
	case common.EndPage:
		e.PageRange.End = v.N
	case common.DateInformation:
		if v.When().IsZero() {
			e.Date.Year = v.Year()
		} else {
			e.Date = NewDate(v.When())
		}
	case common.DOI:
		e.SerialNumber.Doi = v.Number
	case common.PMID:
		e.SerialNumber.Pmid = v.Number
	case common.PMCID:
		e.SerialNumber.Pmcid = v.Number
	case common.SN:
		e.SerialNumber.Serial = v.Number
	case common.Web:
		e.Url.Value = v.Content
		log.Println("url found, may need to set accession date manually")
	case common.Language:
		e.Language = v.Code
	case common.Classification:
		e.Genre = v.Content
	case common.CallNumber:
		e.CallNumber = v.Content
	case common.Abstract:
		e.Abstract = v.Content
	case common.Availability:
		e.Note = v.Content
	case common.Custom:
		e.Note = v.Content
	case common.Notes:
		e.Note = v.Content
	case common.PersonalNotes:
		e.Note = v.Content
	case common.ResearchNotes:
		e.Note = v.Content
	case common.UserDefinableNotes:
		e.Note = v.Content
	case common.Freeform:
		e.Note = v.Data
	case common.Miscellaneous:
		e.Note = v.Data
	case common.Booktitle:
		e.Title = v.Title
	case common.Caption:
		e.Title = v.Content
	case common.Keyword:
		e.Title = v.Content
	case common.Title:
		e.Title = v.Content
	case common.SecondaryTitle:
		e.Title = v.Content
	case common.TertiaryTitle:
		e.Title = v.Content
	case common.TranslatedTitle:
		e.Title = v.Content
	case common.ShortTitle:
		e.Title = v.Content
	case common.WebsiteTitle:
		e.Title = v.Content
	default:
		return WriterError{"writing unknown information", i.String()}
	}

	return nil
}

type PageRange struct {
	Start int
	End   int
}

func (r PageRange) MarshalYAML() ([]byte, error) {
	return []byte(fmt.Sprintf("%d-%d", r.Start, r.End)), nil
}

type Serial struct {
	Doi    string `yaml:",omitempty"`
	Isbn   string `yaml:",omitempty"`
	Issn   string `yaml:",omitempty"`
	Pmid   string `yaml:",omitempty"`
	Pmcid  string `yaml:",omitempty"`
	Arxiv  string `yaml:",omitempty"`
	Serial string `yaml:",omitempty"`
}

type Url struct {
	Value string
	Date  Date `yaml:"date,omitempty"`
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func NewDate(t time.Time) Date {
	return Date{Year: t.Year(), Month: int(t.Month()), Day: t.Day()}
}

func (d Date) String() string {
	if d.Month == 0 || d.Day == 0 {
		return fmt.Sprintf("%04d", d.Year)
	} else {
		return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
	}
}

func (d Date) MarshalYAML() ([]byte, error) {
	s := d.String()
	return []byte(s), nil
}

type Person struct {
	Role  string
	Names []string
}

type Publisher struct {
	Name     string
	Location string `yaml:"location,omitempty"`
}
