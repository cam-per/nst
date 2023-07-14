package epub

import (
	"archive/zip"
	"io"
	"path/filepath"
)

type instance struct {
	zip.FileHeader

	e *Encoder
	w io.Writer

	id   string
	mime string
}

func (inst *instance) create(e *Encoder, path string) error {
	inst.FileHeader = zip.FileHeader{
		Name:   path,
		Method: zip.Deflate,
	}

	inst.e = e

	var err error

	inst.w, err = inst.e.w.CreateHeader(&inst.FileHeader)
	if err != nil {
		return err
	}

	if stringEmpty(inst.mime) {
		e.files = append(e.files, inst)
	}

	return nil
}

func (inst *instance) href(basepath string) string {
	path, err := filepath.Rel(basepath, inst.Name)
	if err != nil {
		return inst.Name
	}
	return path
}

func (inst *instance) createFile(e *Encoder, src io.Reader, mime, path string) error {
	inst.FileHeader = zip.FileHeader{
		Name:   path,
		Method: zip.Deflate,
	}

	inst.e = e
	inst.mime = mime

	var err error
	inst.w, err = inst.e.w.CreateHeader(&inst.FileHeader)
	if err != nil {
		return err
	}

	if stringEmpty(inst.mime) {
		e.files = append(e.files, inst)
	}

	_, err = io.Copy(inst.w, src)
	return err
}

type Encoder struct {
	w *zip.Writer

	book *Book

	_ncx ncx
	_opf opf

	cover instance

	files    []*instance
	chapters []*Chapter
}

func (e *Encoder) writeMimeType(path string) error {
	var mimetype instance
	err := mimetype.create(e, "mimetype")
	if err != nil {
		return err
	}
	_, err = mimetype.w.Write([]byte("application/epub+zip"))
	return err
}

func (e *Encoder) Encode() error {
	var (
		err error
	)

	err = e.writeMimeType("mimetype")
	if err != nil {
		return err
	}

	err = e._ncx.init()
	if err != nil {
		return err
	}

	e._opf.init()
	err = e._opf.create(e, "OPS/content.opf")
	if err != nil {
		return err
	}

	return e.w.Close()
}

func NewEncoder(w io.Writer, book *Book) *Encoder {
	e := &Encoder{
		w:    zip.NewWriter(w),
		book: book,
	}
	return e
}
