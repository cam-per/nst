package ranobes

import (
	"github/cam-per/nst"
	"net/url"
)

func init() {
	host = url.URL{
		Scheme: "https",
		Host:   "ranobes.com",
	}

	nst.RegisterSource(&src)
}
