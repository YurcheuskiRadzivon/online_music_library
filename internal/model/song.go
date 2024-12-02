package model

type Song struct {
	SoundId     int    `json:"sound_id,omitempty"`
	Group       string `json:"group,omitempty"`
	Song        string `json:"song,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Text        string `json:"text,omitempty"`
	Link        string `json:"link,omitempty"`
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
