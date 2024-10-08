package database

type wish struct {
	Coin    string `json:"coin"`
	Content string `json:"content"`
	Who     string `json:"who"`
}

type stats struct {
	STR int `json:"STR"`
	CON int `json:"CON"`
	DEX int `json:"DEX"`
	INT int `json:"INT"`
	WIS int `json:"WIS"`
	CHR int `json:"CHR"`
}

type Campaign struct {
	UUID          string `json:"uuid"`
	CampaignOrder int    `json:"campaign_order"`
	Title         string `json:"title"`
	Saga          string `json:"saga"`
	PosterURL     string `json:"poster_url"`
	Synopsis      string `json:"synopsis"`
	Recap         string `json:"recap"`
	PlaylistID    string `json:"playlist_id"`
	IsPlanned     bool   `json:"is_planned"`
	IsPlayed      bool   `json:"is_played"`
	IsVisible     bool   `json:"is_visible"`
	DurationHrs   int    `json:"duration_hrs"`
	Wish          wish   `json:"wish"`
}

type Variable struct {
	SpamResponseChance float64 `json:"spam_response_chance"`
}

type User struct {
	UUID          string   `json:"uuid"`
	TgUserID      string   `json:"tg_user_id"`
	IsWhitelisted bool     `json:"is_whitelisted"`
	IsGM          bool     `json:"is_gm"`
	Variables     Variable `json:"variables"`
}

type Map struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	IsVisible bool   `json:"is_visible"`
}

type Character struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Background string `json:"backround"`
	Goals      string `json:"goals"`
	Level      int    `json:"level"`
	Class      string `json:"Class"`
	PosterURL  string `json:"poster_url"`
	OwnerUUID  string `json:"owner_uuid"`
	Status     string `json:"status"`
	Stats      stats  `json:"stats"`
}

type Secret struct {
	DefaultSpamResponseChance float64  `json:"default_spam_response_chance"`
	SpamResponseChanceFactor  float64  `json:"spam_response_chance_factor"`
	Responses                 Response `json:"responses"`
}

type Response struct {
	SpamResponse string `json:"spam_response"`
}

type ByCampaignOrder []Campaign

func (a ByCampaignOrder) Len() int           { return len(a) }
func (a ByCampaignOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCampaignOrder) Less(i, j int) bool { return a[i].CampaignOrder < a[j].CampaignOrder }
