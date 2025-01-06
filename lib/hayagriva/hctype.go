package hayagriva

import (
	"fmt"
	"log"
	"strconv"

	"github.com/scrouthtv/hayagriva-converter/lib/common"
)

func (e HcEntry) SetType(t common.Type) error {
	switch t {
	case common.TypeAbstract, common.TypeAncientText, common.TypeChart,
		common.TypeClassicalWork, common.TypeConferencePaper, common.TypeEquation,
		common.TypeFigure, common.TypeHearing, common.TypeGrant,
		common.TypeInterview, common.TypeOnlineMultimedia, common.TypePamphlet,
		common.TypePersonalCommunication, common.TypePressRelease, common.TypeSlide,
		common.TypeStatute, common.TypeUnpublished:
		e.Type = "misc"
		log.Println("unknown type found, consider using parent with `misc`")
	case common.TypeAudiovisual:
		e.Type = "video"
		e.Genre = "Audiovisual"
	case common.TypeAggregatedDatabase:
		e.Type = "anthology"
	case common.TypeArtwork:
		e.Type = "artwork"
	case common.TypeBillResolution:
		e.Type = "legislation"
	case common.TypeBlog:
		e.Type = "blog"
	case common.TypeBookWhole:
		e.Type = "book"
	case common.TypeCase:
		e.Type = "case"
	case common.TypeBookSection:
		e.Type = "chapter"
	case common.TypeComputerProgram:
		e.Type = "repository"
	case common.TypeConferenceProceedings:
		e.Type = "proceedings"
	case common.TypeCatalog:
		e.Type = "anthology"
	case common.TypeDatafile:
		e.Type = "reference"
		e.Genre = "Datafile"
	case common.TypeOnlineDatabase:
		e.Type = "web"
		e.Genre = "Database"
	case common.TypeDictionary:
		e.Type = "reference"
		e.Genre = "Dictionary"
	case common.TypeElectronicBook:
		e.Type = "book"
		e.Genre = "Electronic Book"
	case common.TypeElectronicBookSection:
		e.Type = "chapter"
		e.Genre = "Electronic Book"
	case common.TypeEditedBook:
		log.Println("edited book found,consider using parent with `original`")
		e.Type = "book"
	case common.TypeElectronicArticle:
		e.Type = "article"
		e.Genre = "Electronic Article"
	case common.TypeWebPage:
		e.Type = "web"
	case common.TypeEncyclopedia:
		e.Type = "reference"
		e.Genre = "Encyclopedia"
	case common.TypeGeneric:
		e.Type = "misc"
	case common.TypeGovernment:
		e.Type = "legislation"
	case common.TypeInternetCommunication:
		e.Type = "thread"
	case common.TypePressArticle:
		e.Type = "article"
		e.Genre = "Press Article"
	case common.TypePeriodicalFull:
		e.Type = "periodical"
	case common.TypePeriodicalArticle:
		e.Type = "article"
	case common.TypeLegal:
		e.Type = "case"
	case common.TypeManuscript:
		e.Type = "manuscript"
	case common.TypeCartographicData:
		e.Type = "reference"
		e.Genre = "Cartographic Data"
	case common.TypeMagazineArticle:
		e.Type = "article"
		e.Genre = "Magazine Article"
	case common.TypeFilmBroadcast:
		e.Type = "video"
	case common.TypeMusic:
		e.Type = "audio"
	case common.TypeNewspaper:
		e.Type = "newspaper"
	case common.TypePatent:
		e.Type = "patent"
	case common.TypePodcast:
		e.Type = "audio"
		e.Genre = "Podcast"
	case common.TypeReport:
		e.Type = "report"
	case common.TypeSerial:
		e.Type = "periodical"
	case common.TypeSound:
		e.Type = "audio"
	case common.TypeStandard:
		e.Type = "misc"
	case common.TypeThesis:
		e.Type = "thesis"
	case common.TypeUnenacted:
		e.Type = "bill"
		e.Genre = "Unenacted"
	case common.TypeVideo:
		e.Type = "video"
	case common.TypeWebpage:
		e.Type = "web"
	default:
		return WriterError{"writing type", "unknown type " + strconv.Itoa(int(t))}
	}

	return nil
}

type WriterError struct {
	What     string
	Actually string
}

func (e WriterError) Error() string {
	return fmt.Sprintf("writer error during %s: got %s", e.What, e.Actually)
}
