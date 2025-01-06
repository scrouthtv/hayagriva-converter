package common

type Type uint8

const (
	TypeAbstract Type = iota
	TypeAudiovisual
	TypeAggregatedDatabase
	TypeAncientText
	TypeArtwork
	TypeBillResolution
	TypeBlog
	TypeBookWhole
	TypeCase
	TypeBookSection
	TypeChart
	TypeClassicalWork
	TypeComputerProgram
	TypeConferenceProceedings
	TypeConferencePaper
	TypeCatalog
	TypeDatafile
	TypeOnlineDatabase
	TypeDictionary
	TypeElectronicBook
	TypeElectronicBookSection
	TypeEditedBook
	TypeElectronicArticle
	TypeWebPage
	TypeEncyclopedia
	TypeEquation
	TypeFigure
	TypeGeneric
	TypeGovernment
	TypeGrant
	TypeHearing
	TypeInternetCommunication
	TypePressArticle
	TypeInterview
	TypePeriodicalFull
	TypePeriodicalArticle
	TypeLegal
	TypeManuscript
	TypeCartographicData
	TypeMagazineArticle
	TypeFilmBroadcast
	TypeOnlineMultimedia
	TypeMusic
	TypeNewspaper
	TypePamphlet
	TypePatent
	TypePersonalCommunication
	TypePodcast
	TypePressRelease
	TypeReport
	TypeSerial
	TypeSlide
	TypeSound
	TypeStandard
	TypeStatute
	TypeThesis
	TypeUnenacted
	TypeUnpublished
	TypeVideo
	TypeWebpage
)

func (t Type) String() string {
	switch t {
	case TypeAbstract:
		return "type abstract"
	case TypeAudiovisual:
		return "type audiovisual"
	case TypeAggregatedDatabase:
		return "type aggregated database"
	case TypeAncientText:
		return "type ancient text"
	case TypeArtwork:
		return "type artwork"
	case TypeBillResolution:
		return "type bill resolution"
	case TypeBlog:
		return "type blog"
	case TypeBookWhole:
		return "type book whole"
	case TypeCase:
		return "type case"
	case TypeBookSection:
		return "type book section"
	case TypeChart:
		return "type chart"
	case TypeClassicalWork:
		return "type classical work"
	case TypeComputerProgram:
		return "type computer program"
	case TypeConferenceProceedings:
		return "type conference proceedings"
	case TypeConferencePaper:
		return "type conference paper"
	case TypeCatalog:
		return "type catalog"
	case TypeDatafile:
		return "type datafile"
	case TypeOnlineDatabase:
		return "type online database"
	case TypeDictionary:
		return "type dictionary"
	case TypeElectronicBook:
		return "type electronic book"
	case TypeElectronicBookSection:
		return "type electronic book section"
	case TypeEditedBook:
		return "type edited book"
	case TypeElectronicArticle:
		return "type electronic article"
	case TypeWebPage:
		return "type web page"
	case TypeEncyclopedia:
		return "type encyclopedia"
	case TypeEquation:
		return "type equation"
	case TypeFigure:
		return "type figure"
	case TypeGeneric:
		return "type generic"
	case TypeGovernment:
		return "type government"
	case TypeGrant:
		return "type grant"
	case TypeHearing:
		return "type hearing"
	case TypeInternetCommunication:
		return "type internet communication"
	case TypePressArticle:
		return "type press article"
	case TypeInterview:
		return "type interview"
	case TypePeriodicalFull:
		return "type periodical full"
	case TypePeriodicalArticle:
		return "type periodical article"
	case TypeLegal:
		return "type legal"
	case TypeManuscript:
		return "type manuscript"
	case TypeCartographicData:
		return "type cartographic data"
	case TypeMagazineArticle:
		return "type magazine article"
	case TypeFilmBroadcast:
		return "type film broadcast"
	case TypeOnlineMultimedia:
		return "type online multimedia"
	case TypeMusic:
		return "type music"
	case TypeNewspaper:
		return "type newspaper"
	case TypePamphlet:
		return "type pamphlet"
	case TypePatent:
		return "type patent"
	case TypePersonalCommunication:
		return "type personal communication"
	case TypePodcast:
		return "type podcast"
	case TypePressRelease:
		return "type press release"
	case TypeReport:
		return "type report"
	case TypeSerial:
		return "type serial"
	case TypeSlide:
		return "type slide"
	case TypeSound:
		return "type sound"
	case TypeStandard:
		return "type standard"
	case TypeStatute:
		return "type statute"
	case TypeThesis:
		return "type thesis"
	case TypeUnenacted:
		return "type unenacted"
	case TypeUnpublished:
		return "type unpublished"
	case TypeVideo:
		return "type video"
	case TypeWebpage:
		return "type webpage"
	default:
		return "unknown type"
	}
}
