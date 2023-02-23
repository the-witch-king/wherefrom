package jikan

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

type JikanGetPersonFullResponse struct {
	Data JikanPersonFull `json:"data"`
}

type JikanPersonFull struct {
	MalID      int    `json:"mal_id"`
	URL        string `json:"url"`
	WebsiteURL string `json:"website_url"`
	Images     struct {
		Jpg struct {
			ImageURL string `json:"image_url"`
		} `json:"jpg"`
	} `json:"images"`
	Name           string   `json:"name"`
	GivenName      string   `json:"given_name"`
	FamilyName     string   `json:"family_name"`
	AlternateNames []string `json:"alternate_names"`
	Birthday       string   `json:"birthday"`
	Favorites      int      `json:"favorites"`
	About          string   `json:"about"`
	Anime          []struct {
		Position string `json:"position"`
		Anime    struct {
			MalID  int    `json:"mal_id"`
			URL    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageURL      string `json:"image_url"`
					SmallImageURL string `json:"small_image_url"`
					LargeImageURL string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageURL      string `json:"image_url"`
					SmallImageURL string `json:"small_image_url"`
					LargeImageURL string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"anime"`
	} `json:"anime"`
	Manga []struct {
		Position string `json:"position"`
		Manga    struct {
			MalID  int    `json:"mal_id"`
			URL    string `json:"url"`
			Images struct {
				Jpg struct {
					ImageURL      string `json:"image_url"`
					SmallImageURL string `json:"small_image_url"`
					LargeImageURL string `json:"large_image_url"`
				} `json:"jpg"`
				Webp struct {
					ImageURL      string `json:"image_url"`
					SmallImageURL string `json:"small_image_url"`
					LargeImageURL string `json:"large_image_url"`
				} `json:"webp"`
			} `json:"images"`
			Title string `json:"title"`
		} `json:"manga"`
	} `json:"manga"`
	Voices []JikanPersonVoice
}
