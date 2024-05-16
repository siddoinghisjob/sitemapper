package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"sitemap/utils"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

func main() {
	website := flag.String("w", "http://twitter.com/", "Webiste to scrap data from.")
	flag.Parse()

	aTags := utils.NewParser(*website)
	toXml := utils.Urlset{
		Xmlns: xmlns,
	}
	for _, page := range aTags {
		toXml.Urls = append(toXml.Urls, utils.Loc{Value: page.Link})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}
