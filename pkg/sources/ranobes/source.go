package ranobes

import (
	"github/cam-per/nst"
	"net/url"
)

var (
	host url.URL

	src source
)

type source struct {
	handlerProgress nst.HandlerProgress
}
