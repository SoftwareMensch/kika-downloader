package utils

import (
	"gopkg.in/xmlpath.v2"
	"io"
	"bytes"
	"golang.org/x/net/html"
)

// NodeFromHtmlBody return *xmlpath.Node from HTML content
func NodeFromHtmlBody(readCloser io.ReadCloser) (*xmlpath.Node, error) {
	docRoot, err := html.Parse(readCloser)
	if err != nil {
		return nil, nil
	}

	// we make this to fix broken html
	var buf bytes.Buffer
	html.Render(&buf, docRoot)

	reader := bytes.NewReader(buf.Bytes())
	xmlRoot, err := xmlpath.ParseHTML(reader)
	if err != nil {
		return nil, err
	}

	return xmlRoot, nil
}
