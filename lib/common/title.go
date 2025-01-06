package common

type Booktitle struct {
	Title string
}

func (b Booktitle) String() string {
	return "booktitle: " + b.Title
}

type Caption struct {
	Content string
}

func (c Caption) String() string {
	return "caption: " + c.Content
}

type Keyword struct {
	Content string
}

func (k Keyword) String() string {
	return "keyword: " + k.Content
}

type Title struct {
	Content string
}

func (t Title) String() string {
	return "title: " + t.Content
}

type SecondaryTitle struct {
	Content string
}

func (s SecondaryTitle) String() string {
	return "secondary title: " + s.Content
}

type TertiaryTitle struct {
	Content string
}

func (t TertiaryTitle) String() string {
	return "tertiary title: " + t.Content
}

type TranslatedTitle struct {
	Content string
}

func (t TranslatedTitle) String() string {
	return "translated title: " + t.Content
}

type ShortTitle struct {
	Content string
}

func (s ShortTitle) String() string {
	return "short title: " + s.Content
}

type WebsiteTitle struct {
	Content string
}

func (w WebsiteTitle) String() string {
	return "website title: " + w.Content
}
