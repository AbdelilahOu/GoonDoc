package model

type YtVideo struct {
	Title     string  `json:"title"`
	Thumbnail string  `json:"thumbnail"`
	Channel   Channel `json:channel`
	Details   Details `json:details`
}

type Channel struct {
	Name  string `json:"name"`
	Image string `json:"Image"`
}

type Details struct {
	Views       string `json:"views"`
	ReleaseTime string `json:"releaseTime"`
}
