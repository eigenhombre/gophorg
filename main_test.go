package main

import (
	"bytes"
	"errors"
	"golang.org/x/net/html"
	"io"
	"strings"
	"testing"
)

// Most of this code will eventually be reworked or moved into
// production code.  Using tests as a primitive REPL for now.

// Adapted from https://stackoverflow.com/a/38855264/611752:
func FindElement(elementName string, doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == elementName {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(node *html.Node) string {
	var buf bytes.Buffer
	html.Render(io.Writer(&buf), node)
	return buf.String()
}

func TestHTMLParsing(t *testing.T) {
	doc, _ := html.Parse(strings.NewReader(htm))
	bn, _ := FindElement("body", doc)
	body := renderNode(bn)
	// FIXME: Why does it add a newline?!
	desired := `<body>
    body content
    <p>more content</p>

</body>`
	if 0 != (strings.Compare(body, desired)) {
		t.Log("Unexpected parse output")
		t.Fail()
	}
}

const htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
    body content
    <p>more content</p>
</body>
</html>`
