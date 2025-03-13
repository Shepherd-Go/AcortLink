package entity

type Success struct {
	URL `json:"url"`
}

type URL struct {
	Original_Url string `json:"original_url,omitempty"`
	Short_Url    string `json:"short_url,omitempty"`
}
