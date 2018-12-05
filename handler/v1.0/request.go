package v1_0

type GetVideoThemesReq struct {
}

type GetVideoThemesResp struct {
}

type GetVideoThemeReq struct {
}

type GetVideoThemeResp struct {
	VideoThemeId string `json:"video_theme_id"`
	Name         string
	Description  string
}

type AddVideoThemeReq struct {
	Name        string
	Description string
}

type AddVideoThemeResp struct {
	VideoThemeId string `json:"video_theme_id"`
	Name         string
	Description  string
}

type UpdateVideoThemeReq struct {
}

type UpdateVideoThemeResp struct {
}

type GetVideosReq struct {
}

type GetVideosResp struct {
	VideoId       string `json:"video_id"`
	Title         string `json:"title"`
	StorePath     string `json:"store_path"`
	Thumbnail     string `json:"thumbnail"`
	Description   string `json:"description"`
	IsCharge      bool   `json:"is_charge"`
	FreeTime      string `json:"free_time"`
	ThemeId       int    `json:"theme_id"`
	PublishUserId string `json:"publish_user_id"`
}

type AddVideoReq struct {
	Title         string `json:"title"`
	StorePath     string `json:"store_path"`
	Thumbnail     string `json:"thumbnail"`
	Description   string `json:"description"`
	IsCharge      bool   `json:"is_charge"`
	FreeTime      string `json:"free_time"`
	ThemeId       int    `json:"theme_id"`
	PublishUserId string `json:"publish_user_id"`
}

type AddVideoResp struct {
	VideoId       string `json:"video_id"`
	Title         string `json:"title"`
	StorePath     string `json:"store_path"`
	Thumbnail     string `json:"thumbnail"`
	Description   string `json:"description"`
	IsCharge      bool   `json:"is_charge"`
	FreeTime      string `json:"free_time"`
	ThemeId       int    `json:"theme_id"`
	PublishUserId string `json:"publish_user_id"`
}

type GetVideoReq struct {
}

type GetVideoResp struct {
}

type UpdateVideoReq struct {
}

type UpdateVideoResp struct {
}
