package epub

import (
	"fmt"
	"io"
	"net/url"

	"github.com/beevik/etree"
)

type ncx struct {
	instance

	playOrder int

	doc  *etree.Document
	root *etree.Element
	head *etree.Element
	nav  *etree.Element
}

func (n *ncx) createMeta(name string, content any) {
	meta := n.head.CreateElement("meta")
	meta.CreateAttr(name, fmt.Sprint(content))
}

func (n *ncx) init() error {
	n.instance.id = "ncx"
	n.instance.mime = "application/x-dtbncx+xml"

	err := n.instance.create(n.e, "OPS/toc.ncx")
	if err != nil {
		return err
	}

	n.doc = etree.NewDocument()

	n.root = n.doc.CreateElement("ncx")
	n.root.CreateAttr("xmlns", "http://www.daisy.org/z3986/2005/ncx/")
	n.root.CreateAttr("version", "2005-1")

	n.head = n.root.CreateElement("head")
	n.nav = n.root.CreateElement("navMap")

	return nil
}

func (n *ncx) addNav(id, src, text string) {
	n.playOrder++

	point := n.nav.CreateElement("navPoint")
	point.CreateAttr("id", id)
	point.CreateAttr("playOrder", fmt.Sprint(n.playOrder))

	point.CreateElement("navLabel").CreateElement("text").CreateText(text)
	point.CreateElement("content").CreateAttr("src", src)
}

func (n *ncx) encode() error {
	n.doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	if stringEmpty(n.e.book.Title) {
		n.root.CreateElement("docTitle").
			CreateElement("text").
			CreateText(n.e.book.Title)
	}

	if stringEmpty(n.e.book.Creator) {
		n.root.CreateElement("docAuthor").
			CreateElement("text").
			CreateText(n.e.book.Creator)
	}

	n.createMeta("dtb:uid", n.e.book.UID)
	n.createMeta("dtb:depth", 2)
	n.createMeta("dtb:totalPageNumber", 0)
	n.createMeta("dtb:maxPageNumber", 0)

	n.doc.Indent(2)
	_, err := n.doc.WriteTo(n.w)
	return err
}

type opf struct {
	instance

	doc      *etree.Document
	root     *etree.Element
	metadata *etree.Element
	manifest *etree.Element
	spine    *etree.Element
}

func (o *opf) init() {
	o.doc = etree.NewDocument()

	o.root = o.doc.CreateElement("package")
	o.root.CreateAttr("xmlns", "http://www.idpf.org/2007/opf")
	o.root.CreateAttr("version", "3.0")
	o.root.CreateAttr("unique-identifier", "bookid")

	o.metadata = o.root.CreateElement("metadata")
	o.metadata.CreateAttr("xmlns:dc", "http://purl.org/dc/elements/1.1/")
	o.metadata.CreateAttr("xmlns:opf", "http://www.idpf.org/2007/opf")

	o.manifest = o.root.CreateElement("manifest")

	o.spine = o.root.CreateElement("spine")
}

func (o *opf) encode() error {
	o.doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	bookid := o.metadata.CreateElement("dc:identifier")
	bookid.CreateAttr("id", "bookid")
	bookid.CreateText(o.e.book.UID)

	elemTextOmitempty(o.metadata, "dc:title", o.e.book.Title)
	elemTextOmitempty(o.metadata, "dc:creator", o.e.book.Creator)
	elemTextOmitempty(o.metadata, "dc:language", o.e.book.Lang)

	o.spine.CreateAttr("toc", o.e._ncx.id)

	o.doc.Indent(2)
	_, err := o.doc.WriteTo(o.w)
	return err
}

type Chapter struct {
	inst instance

	url url.URL

	body []byte
}

type Book struct {
	UID         string
	Title       string
	Creator     string
	Description string
	Lang        string

	CoverImage io.Reader
}
