package request



type Setin struct {
	Sentence []string `json:"sentence"`
	FontSize []int    `json:"fontsize"`
	Spacing  int      `json:"spacing"`
}
