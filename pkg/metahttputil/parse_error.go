package metahttputil

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type ErrXmlResponse struct {
	Title       string
	Description string
}

func (errxmlresponse *ErrXmlResponse) Error() string {
	return fmt.Sprintf("%s: %s", errxmlresponse.Title, errxmlresponse.Description)
}

func CheckHttpError(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	// ok at this point, check if it is a json error response or a html response
	// TODO: figure out if the api responds with any json errors, and if it does, parse them properly
	//if resp.Header.Get("Content-Type") == "application/json" {
	//}

	if strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "text/html") {
		xmlErr, err := ParseHtmlError(resp.Body)
		if err != nil {
			return err
		}
		return xmlErr
	}

	return fmt.Errorf("http error %s: %d", resp.Status, resp.StatusCode)
}

func ParseHtmlError(body io.Reader) (*ErrXmlResponse, error) {
	nd, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	var errXml ErrXmlResponse
	var populate func(n *html.Node)
	populate = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				errXml.Title = n.FirstChild.Data
			}
		}
		if n.Type == html.ElementNode && n.Data == "p" {
			for x := n.FirstChild; x != nil; x = x.NextSibling {
				if x.Data != "br" {
					errXml.Description = errXml.Description + x.Data
				} else {
					errXml.Description = errXml.Description + " "
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			populate(c)
		}
	}
	populate(nd)
	return &errXml, nil
}
