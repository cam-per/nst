package nst

import (
	"io"
)

type Progress struct {
}

type HandlerProgress func(p Progress)

type SourceProvider interface {
	ID() string

	Search(title string) ([]BookHeader, error)
	Info(path string) (BookInfo, error)
	Save(w io.Writer, path string) error

	OnProgress(handler HandlerProgress)
}

var sources []SourceProvider

func RegisterSource(source SourceProvider) {
	sources = append(sources, source)
}
