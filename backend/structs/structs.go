package structs

type Translate struct {
	Text string `json:"text"`
	To   string `json:"to"`
}
type TranslationRes struct {
	Translation []Translate `json:"translations"`
}

type ChatLog struct {
	Name string `gorm:"name"`
	Text string `gorm:"text"`
	Time string `gorm:"time"`
}

type JsonMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type JsonReturn struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
