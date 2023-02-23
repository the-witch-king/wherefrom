package mal

type MALUserAnimeListResponse struct {
	Data   []MALAnime `json:"data"`
	Paging MALPaging  `json:"paging"`
}

type MALAnime struct {
	Node struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		MainPicture struct {
			Medium string `json:"medium"`
			Large  string `json:"large"`
		}
	} `json:"node"`
	ListStatus struct {
		Status             string `json:"status"`
		Score              int    `json:"score"`
		NumWatchedEpisodes int    `json:"num_watched_episodes"`
		IsRewatching       bool   `json:"is_rewatching"`
		UpdatedAt          string `json:"updated_at"`
	} `json:"list_status"`
}

type MALPaging struct {
	Next string `json:"next"`
}
