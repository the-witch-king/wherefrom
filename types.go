package main

// =============
// = AWS TYPES =
// =============
type UserPayload struct {
	UserName string `json:"username"`
	ActorId string `json:"actorId"`
}

// =========
// = OTHER =
// =========
type Appearance struct {
	Show string `json:"show"`
	Character string `json:"character"`
}


// =============
// = MAL TYPES =
// =============
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

// ===============
// = JIKAN TYPES =
// ===============
type JikanGetPersonVoicesResponse struct {
	Data []JikanPersonVoice `json:"data"`
}

type JikanPersonVoice struct {
	Role  string `json:"role"`
	Anime struct {
		MalID  int    `json:"mal_id"`
		URL    string `json:"url"`
		Images struct {
			Jpg struct {
			} `json:"jpg"`
			Webp struct {
			} `json:"webp"`
		} `json:"images"`
		Title string `json:"title"`
	} `json:"anime"`
	Character struct {
		MalID  int    `json:"mal_id"`
		URL    string `json:"url"`
		Images struct {
		} `json:"images"`
		Name string `json:"name"`
	} `json:"character"`
}
