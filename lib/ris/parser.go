package ris

import (
	"fmt"
	"strconv"
	"time"

	"github.com/scrouthtv/hayagriva-converter/lib/common"
)

type EndTag struct{}

func (e EndTag) String() string {
	return "ris end tag"
}

func ParseToken(t Token) (common.Member, error) {
	switch t.Key {
	case "TY":
		return TypeFromRISString(t.Value)
	case "ER":
		return EndTag{}, nil
	case "AU", "A1":
		return common.PrimaryAuthor{Name: t.Value}, nil
	case "ED", "A2":
		return common.SecondaryAuthor{Name: t.Value}, nil
	case "A3":
		return common.TertiaryAuthor{Name: t.Value}, nil
	case "A4":
		return common.QuaternaryAuthor{Name: t.Value}, nil
	case "A5":
		return common.QuinaryAuthor{Name: t.Value}, nil
	case "A6":
		return common.WebsiteAuthor{Name: t.Value}, nil
	case "TA":
		return common.TranslatedAuthor{Name: t.Value}, nil
	case "AB", "N2":
		return common.Abstract{Content: t.Value}, nil
	case "AV":
		return common.Availability{Content: t.Value}, nil
	case "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8":
		lvl := int(t.Key[1] - '0')
		return common.Custom{Level: lvl, Content: t.Value}, nil
	case "N1", "NO":
		return common.Notes{Content: t.Value}, nil
	case "PA":
		return common.PersonalNotes{Content: t.Value}, nil
	case "RN":
		return common.ResearchNotes{Content: t.Value}, nil
	case "U1", "U2", "U3", "U4", "U5", "U6", "U7", "U8", "U9", "U10",
		"U11", "U12", "U13", "U14", "U15":
		lvl, err := strconv.Atoi(t.Key[1:])
		if err != nil {
			return nil, NewParserError("unknown user definable notes level", t.Key)
		}
		return common.UserDefinableNotes{Level: lvl, Content: t.Value}, nil
	case "AN":
		return common.AccessionNumber{Number: t.Value}, nil
	case "DI", "DO", "DOI":
		return common.DOI{Number: t.Value}, nil
	case "PMID":
		return common.PMID{Number: t.Value}, nil
	case "PMCID":
		return common.PMCID{Number: t.Value}, nil
	case "SN", "ID", "IP", "IS":
		return common.SN{Number: t.Value}, nil
	case "UR":
		return common.Web{Content: t.Value}, nil
	case "FD":
		return common.Freeform{Data: t.Value}, nil
	case "LA":
		return common.Language{Code: t.Value}, nil
	case "CL":
		return common.Classification{Content: t.Value}, nil
	case "CN":
		return common.CallNumber{Content: t.Value}, nil
	case "M1", "M2", "M3":
		lvl := int(t.Key[1] - '0')
		return common.Miscellaneous{Level: lvl, Data: t.Value}, nil
	case "BT":
		return common.Booktitle{Title: t.Value}, nil
	case "CA", "CT":
		return common.Caption{Content: t.Value}, nil
	case "K1", "KW":
		return common.Keyword{Content: t.Value}, nil
	case "T1", "TI":
		return common.Title{Content: t.Value}, nil
	case "T2":
		return common.SecondaryTitle{Content: t.Value}, nil
	case "T3":
		return common.TertiaryTitle{Content: t.Value}, nil
	case "TT":
		return common.TranslatedTitle{Content: t.Value}, nil
	case "ST":
		return common.ShortTitle{Content: t.Value}, nil
	case "WT":
		return common.WebsiteTitle{Content: t.Value}, nil
	case "DA":
		da, err := parseDate(t)
		if err != nil {
			return nil, err
		}
		return common.Date{T: da}, nil
	case "RD", "Y2":
		da, err := parseDate(t)
		if err != nil {
			return nil, err
		}
		return common.RetrievedDate{T: da}, nil
	case "WP":
		da, err := parseDate(t)
		if err != nil {
			return nil, err
		}
		return common.WebDate{T: da}, nil
	case "PY", "Y1":
		da, err := parseDate(t)
		if err != nil {
			return nil, err
		}
		return common.Date{T: da}, nil
	case "YR":
		yr, err := strconv.Atoi(t.Value)
		if err != nil {
			return nil, NewParserError("publication year", err.Error())
		} else {
			return common.PublicationYear{Yr: yr}, nil
		}
	case "JA", "J2", "JF", "JO":
		return common.ParentJournal{Name: t.Value}, nil
	case "SP":
		n, err := strconv.Atoi(t.Value)
		if err != nil {
			return nil, NewParserError("start page number", err.Error())
		} else {
			return common.StartPage{N: n}, nil
		}
	case "EP":
		n, err := strconv.Atoi(t.Value)
		if err != nil {
			return nil, NewParserError("end page number", err.Error())
		} else {
			return common.EndPage{N: n}, nil
		}
	case "VL", "VO", "SV":
		return common.Volume{Content: t.Value}, nil

		/*
		   CP 	City/place of publication.[9][16] Issue.[21][18]
		   CR 	Cited references.[15]
		   CY 	Place published, e.g. city, conference location, country, or activity location.[6][14][21][8][16][18][19][20]
		   DB 	Name of database.[6][14][15][18][20]
		   DP 	Database provider.[6][14][18][20]
		   DS 	Data source.[15]
		   ET 	Edition, e.g. epub (electronic publication?) date, date published, session, action of higher court, version, requirement, description of material, international patent classification, or description.[6][14][16][18][19][20]
		   H1 	Location (library).[16]
		   H2 	Location (call number).[16]
		   L1 	File attachments, e.g. figure.[6][14][18][19][20] "Link to PDF. There is no practical length limit to this field. URL addresses can be entered individually, one per tag or multiple addresses can be entered on one line using a semi-colon as a separator. These links should end with a file name, and not simply a landing page. Use the UR tag for URL links."[9][22] Internet link.[16] Local file.[18]
		   L2 	URL.[14] "Link to Full-text. There is no practical length limit to this field. URL addresses can be entered individually, one per tag or multiple addresses can be entered on one line using a semi-colon as a separator."[9][22] Internet link.[16]
		   L3 	DOI.[14] Related records.[9][22] Internet link.[16]
		   L4 	Figure, e.g. URL or file attachments.[6][14][18][19][20] Images.[9][22] Internet link.[16] Local file.[18]
		   LB 	Label.[6][14][18][20]
		   LK 	Links.[15]
		   LL 	Sponsoring library location.[15]
		   NV 	Number of volumes, e.g. extent of work, reporter abbreviation, catalog number, study number, document number, version, amount received, session number, frequency, manuscript number, US patent classification, communication number, series volume, or statute number. Ignored for Press Release (PRESS).[6][14][18][19][20]
		   OL 	Output language (using one of the documented numeric codes).[15]
		   OP 	Original publication, e.g. contents, history, content, version history, original grant number, or priority numbers.[6][14][18] Other pages.[15][20] Original foreign title.[15]
		   PB 	Publisher, e.g. court, distributor, sponsoring agency, library/archive, assignee, institution, source, or university / degree grantor.[6][14][9][21][8][15][16][17][18][19][20]
		   PP 	Place of publication.[15]
		   RI 	Reviewed item, geographic coverage, or article number.[14][18][20]
		   RP 	Reprint status, e.g. reprint edition, review date, or notes. Has three possible values: "IN FILE", "NOT IN FILE", or "ON REQUEST". "ON REQUEST" must be followed by an MM/DD/YY date in parentheses.[6][14][9][26][8][17][18][20]
		   RT 	Reference type.[15]
		   SE 	Section, screens, code section, message number, pages, chapter, filed date, number of pages, original release date, version, e-pub date, duration of grant, section number, start page, international patent number, or running time.[14][18][20]
		   SF 	Subfile/database.[15]
		   SL 	Sponsoring library.[15]
		   SR 	Source type: Print(0) or Electronic(1).[15]
		   WV 	Website version.[15]
		*/

	default:
		return nil, NewParserError("unknown key type", t.Key)
	}
}

func TypeFromRISString(s string) (common.Type, error) {
	switch s {
	case "ABST":
		return common.TypeAbstract, nil
	case "ADVS":
		return common.TypeAudiovisual, nil
	case "AGGR":
		return common.TypeAggregatedDatabase, nil
	case "ANCIENT":
		return common.TypeAncientText, nil
	case "ART":
		return common.TypeArtwork, nil
	case "BILL":
		return common.TypeBillResolution, nil
	case "BLOG":
		return common.TypeBlog, nil
	case "BOOK":
		return common.TypeBookWhole, nil
	case "CASE":
		return common.TypeLegal, nil
	case "CHAP":
		return common.TypeBookSection, nil
	case "CHART":
		return common.TypeChart, nil
	case "CLSWK":
		return common.TypeClassicalWork, nil
	case "COMP":
		return common.TypeComputerProgram, nil
	case "CONF":
		return common.TypeConferenceProceedings, nil
	case "CPAPER":
		return common.TypeConferencePaper, nil
	case "CTLG":
		return common.TypeCatalog, nil
	case "DATA":
		return common.TypeDatafile, nil
	case "DBASE":
		return common.TypeOnlineDatabase, nil
	case "DICT":
		return common.TypeDictionary, nil
	case "EBOOK":
		return common.TypeElectronicBook, nil
	case "ECHAP":
		return common.TypeElectronicBookSection, nil
	case "EDBOOK":
		return common.TypeEditedBook, nil
	case "EJOUR":
		return common.TypeElectronicArticle, nil
	case "ELEC":
		return common.TypeWebpage, nil
	case "ENCYC":
		return common.TypeEncyclopedia, nil
	case "EQUA":
		return common.TypeEquation, nil
	case "FIGURE":
		return common.TypeFigure, nil
	case "GEN":
		return common.TypeGeneric, nil
	case "GOVDOC":
		return common.TypeGovernment, nil
	case "GRNT":
		return common.TypeGrant, nil
	case "HEAR":
		return common.TypeHearing, nil
	case "ICOMM":
		return common.TypeInternetCommunication, nil
	case "INPR":
		return common.TypePressArticle, nil
	case "INTV":
		return common.TypeInterview, nil
	case "JFULL":
		return common.TypePeriodicalFull, nil
	case "JOUR":
		return common.TypePeriodicalArticle, nil
	case "LEGAL":
		return common.TypeLegal, nil
	case "MANSCPT":
		return common.TypeManuscript, nil
	case "MAP":
		return common.TypeCartographicData, nil
	case "MGZN":
		return common.TypeMagazineArticle, nil
	case "MPCT":
		return common.TypeFilmBroadcast, nil
	case "MULTI":
		return common.TypeOnlineMultimedia, nil
	case "MUSIC":
		return common.TypeMusic, nil
	case "NEWS":
		return common.TypeNewspaper, nil
	case "PAMP":
		return common.TypePamphlet, nil
	case "PAT":
		return common.TypePatent, nil
	case "PCOMM":
		return common.TypePersonalCommunication, nil
	case "POD":
		return common.TypePodcast, nil
	case "PRESS":
		return common.TypePressRelease, nil
	case "RPRT":
		return common.TypeReport, nil
	case "SER":
		return common.TypeSerial, nil
	case "SLIDE":
		return common.TypeSlide, nil
	case "SOUND":
		return common.TypeSound, nil
	case "STAND":
		return common.TypeStandard, nil
	case "STAT":
		return common.TypeStatute, nil
	case "STD":
		return common.TypeGeneric, nil
	case "THES":
		return common.TypeThesis, nil
	case "UNBILL":
		return common.TypeUnenacted, nil
	case "UNPB":
		return common.TypeUnpublished, nil
	case "UNPD":
		return common.TypeUnpublished, nil
	case "VIDEO":
		return common.TypeVideo, nil
	case "WEB":
		return common.TypeWebpage, nil
	default:
		return common.TypeReport, NewParserError("unknown publication type", s)
	}
}

var formats = []string{
"02 Jan 2006",
"02 Jan. 2006",
"Jan. 2006",
"January 2006",
}

func parseDate(tok Token) (time.Time, error) {
	for _, format := range formats {
		t, err := time.Parse(format, tok.Value)
		if err == nil {
			return t, nil
		}
	}
	switch len(tok.Value) {
	case 2:
		return time.Parse("06", tok.Value)
	case 4:
		return time.Parse("2006", tok.Value)
	case 10:
		t, err := time.Parse("2006/01/02", tok.Value)
		if err == nil {
			return t, nil
		}

		return time.Parse("2006-01-02", tok.Value)
	default:
		return time.Time{}, NewParserError("date", fmt.Sprintf("unknown date format: %s", tok.Value))
	}
}

type ParserError struct {
	What   string
	Actual string
}

func NewParserError(what, actual string) ParserError {
	return ParserError{
		What:   what,
		Actual: actual,
	}
}

func (e ParserError) Error() string {
	return fmt.Sprintf("parser error: %s: '%s'", e.What, e.Actual)
}
