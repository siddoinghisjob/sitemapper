package utils

type URL struct {
	Text string
	Link string
}

type Loc struct {
	Value string `xml:"loc"`
}

type Urlset struct {
	Urls  []Loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}
