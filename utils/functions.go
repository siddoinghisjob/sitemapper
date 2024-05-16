package utils

// IOT, Cloud and ML
import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var base string

func NewParser(website string) []URL {
	resp, err := http.Get(website)
	if err != nil {
		fmt.Println("Error occured")
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base = baseUrl.String()
	content, e := html.Parse(resp.Body)

	if e != nil {
		fmt.Printf("Error")
		panic(e)
	}

	return Dfs(content)
}

func Dfs(node *html.Node) []URL {
	if node.Data == "a" && node.Type == html.ElementNode {
		currURL := ToURLS(node)
		if currURL.filterLinks() {
			return []URL{currURL}
		} else {
			return nil
		}
	}

	var hrefs []URL
	for fChild := node.FirstChild; fChild != nil; fChild = fChild.NextSibling {
		tmp := Dfs(fChild)
		if tmp != nil {
			hrefs = append(hrefs, tmp...)
		}
	}

	return hrefs
}

func ToURLS(hrefs *html.Node) URL {
	var url URL
	text := strings.Join(strings.Fields((helperToURLS(hrefs))), " ")
	link := hrefExtractor(hrefs)
	url = URL{Text: text, Link: link}
	return url
}

func helperToURLS(href *html.Node) string {
	if href.Type == html.TextNode {
		return href.Data
	}

	text := ""
	for child := href.FirstChild; child != nil; child = child.NextSibling {
		text += child.Data
	}
	return text
}

func hrefExtractor(href *html.Node) string {
	for _, attr := range href.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func (u *URL) filterLinks() bool {
	if strings.HasPrefix(u.Link, "/") {
		newLink := base + u.Link
		u.Link = newLink
	}

	return strings.HasPrefix(u.Link, base)
}
