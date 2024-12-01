package model

type Song struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func NewSong(request SongRequest, detail SongDetail) Song {
	return Song{
		Group:       request.Group,
		Song:        request.Song,
		ReleaseDate: detail.ReleaseDate,
		Text:        detail.Text,
		Link:        detail.Link,
	}
}
