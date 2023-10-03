package model

type Result struct {
	Videos []YtVideo `json:"videos"`
}

type YtVideo struct {
	Title       string  `json:"title"`
	Thumbnail   string  `json:"thumbnail"`
	ChannelData Channel `json:"channel"`
	DetailsData Details `json:"details"`
}

type Channel struct {
	Name  string `json:"name"`
	Image string `json:"image"` // Changed "Image" to "image"
}

type Details struct {
	Views       string `json:"views"`
	ReleaseTime string `json:"releaseTime"`
}
