package epub

import (
	"strings"

	"github.com/beevik/etree"
)

func stringEmpty(s string) bool { return strings.TrimSpace(s) == "" }

func elemTextOmitempty(parent *etree.Element, tag, text string) {
	if strings.TrimSpace(text) == "" {
		return
	}
	parent.CreateElement(tag).CreateText(text)
}
