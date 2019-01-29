package model

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Uid      string `json:"uid"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     Icon   `json:"icon"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}
