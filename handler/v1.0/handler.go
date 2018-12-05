package v1_0

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	commonErrors "github.com/chenjun-git/umbrella-common/errors"

	"github.com/chenjun-git/umbrella-video/common"
	"github.com/chenjun-git/umbrella-video/db"
	"github.com/chenjun-git/umbrella-video/model"
	"github.com/chenjun-git/umbrella-video/utils"
)

func getVideoThemes(w http.ResponseWriter, r *http.Request) {

}

func getVideoThemeById(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(chi.URLParam(r, "video_id"))
	if err != nil {
		return
	}

	videoTheme, err := model.GetVideoThemeById(db.MySQL, videoId)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	} else if videoTheme == nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoThemeNotFound, "video theme not fount"))
		return
	}

	utils.JSON(w, r, wrapCode(videoTheme))
}

func addVideoTheme(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoRequestIOError, err.Error()))
		return
	}
	var req AddVideoThemeReq
	if err = json.Unmarshal(body, &req); err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoJsonUnmarshalErr, err.Error()))
		return
	}

	err = model.CreateVideoTheme(db.MySQL, req.Name, req.Description)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	}

	videoTheme, err := model.GetVideoThemeByName(db.MySQL, req.Name)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	}

	utils.JSON(w, r, wrapCode(map[string]interface{}{
		"Name":videoTheme.Name,
	}))
}

func updateVideoThemeById(w http.ResponseWriter, r *http.Request) {

}

func deleteVideoThemeById(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(chi.URLParam(r, "video_id"))
	if err != nil {
		return
	}

	err = model.DeleteVideoThemeById(db.MySQL, videoId)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	}

	utils.JSON(w, r, wrapCode(""))
}

func getVideos(w http.ResponseWriter, r *http.Request) {

}

func getVideoById(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(chi.URLParam(r, "video_id"))
	if err != nil {
		return
	}

	video, err := model.GetVideoById(db.MySQL, videoId)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	} else if video == nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoNotFound, "video not fount"))
		return
	}

	utils.JSON(w, r, wrapCode(video))
	return 
}

func addVideo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoRequestIOError, err.Error()))
		return
	}
	var req AddVideoReq
	if err = json.Unmarshal(body, &req); err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoJsonUnmarshalErr, err.Error()))
		return
	}

	_, err = model.GetVideoThemeById(db.MySQL, req.ThemeId)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaVideoThemeNotFound, err.Error()))
		return
	}

	err = model.CreateVideo(db.MySQL, req.Title, req.StorePath, req.Thumbnail, req.Description, req.IsCharge, req.ThemeId, req.PublishUserId)
	if err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	}

	// video, err := model.GetVideo(req.Name)
	// if err != nil {
	// 	utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
	// 	return
	// }

	utils.JSON(w, r, wrapCode(map[string]interface{}{
		// "video_theme_id":videoTheme.VideoThemeId,
		// "name":videoTheme.Name,
		// "Description":videoTheme.Description,
	}))
}

func updateVideoById(w http.ResponseWriter, r *http.Request) {
	
}

func deleteVideoById(w http.ResponseWriter, r *http.Request) {
	videoId, err := strconv.Atoi(chi.URLParam(r, "video_id"))
	if err != nil {
		return
	}

	if err := model.DeleteVideoById(db.MySQL, videoId); err != nil {
		utils.JSON(w, r, commonErrors.NewError(common.UmbrellaMysqlError, err.Error()))
		return
	}

	utils.JSON(w, r, wrapCode(""))
}
