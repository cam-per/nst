package nst

var (
	Version = "dev-run"
)

type BookHeader struct {
	URL    string
	Name   string
	Source string
}

type BookInfo struct {
	BookHeader

	AltNames    []string
	Description string

	Year       string
	OriginLang string

	Authors     []string
	Translators []string
	Publishers  []string

	ChaptersCount       int
	TranslateddChapters int
}

func Search(title string) []BookHeader {
	var result []BookHeader
	for _, src := range sources {
		s, err := src.Search(title)
		if err != nil {
			continue
		}
		result = append(result, s...)
	}
	return result
}
