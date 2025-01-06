package common

type Author interface {
	AuthorName() string
}

type PrimaryAuthor struct {
	Name string
}

func (a PrimaryAuthor) AuthorName() string {
	return a.Name
}

func (a PrimaryAuthor) String() string {
	return "primary author: " + a.Name
}

type SecondaryAuthor struct {
	Name string
}

func (a SecondaryAuthor) AuthorName() string {
	return a.Name
}

func (a SecondaryAuthor) String() string {
	return "secondary author: " + a.Name
}

type TertiaryAuthor struct {
	Name string
}

func (a TertiaryAuthor) AuthorName() string {
	return a.Name
}

func (a TertiaryAuthor) String() string {
	return "tertiary author: " + a.Name
}

type QuaternaryAuthor struct {
	Name string
}

func (a QuaternaryAuthor) AuthorName() string {
	return a.Name
}

func (a QuaternaryAuthor) String() string {
	return "quaternary author: " + a.Name
}

type QuinaryAuthor struct {
	Name string
}

func (a QuinaryAuthor) AuthorName() string {
	return a.Name
}

func (a QuinaryAuthor) String() string {
	return "quinary author: " + a.Name
}

type WebsiteAuthor struct {
	Name string
}

func (a WebsiteAuthor) AuthorName() string {
	return a.Name
}

func (a WebsiteAuthor) String() string {
	return "website author: " + a.Name
}

type TranslatedAuthor struct {
	Name string
}

func (a TranslatedAuthor) AuthorName() string {
	return a.Name
}

func (a TranslatedAuthor) String() string {
	return "translated author: " + a.Name
}
