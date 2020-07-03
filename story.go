package goCYOA

// Story struct type required for navigation
type Story map[string]Chapter

// Chapter struct defines that structure used to match the key-value pairs that represent paragraphs in cyoaData.json
type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

// Option struct
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
