package model

import (
	"database/sql"
	"fmt"

	"github.com/chenjun-git/umbrella-video/db"
)

const (
	VideoTableName = "video"

	VideoFieldId            = "video_id"
	VideoFieldTitle         = "title"
	VideoFieldStorePath     = "store_path"
	VideoFieldThumbnail     = "thumbnail"
	VideoFieldDescription   = "description"
	VideoFieldIsCharge      = "is_charge"
	VideoFieldFreeTime      = "free_time"
	VideoFieldThemeId       = "theme_id"
	VideoFieldPublishUserID = "publish_user_id"
)

type Video struct {
	VideoID       int
	Title         string
	StorePath     string
	Thumbnail     string
	Description   string
	IsCharge      bool
	FreeTime      int
	ThemeID       int
	PublishUserID string
}

func CreateVideo(db db.MySQLExec, title, storePath, thumbnail, description string, isCharge bool, themeID int, publishUserID string) error {
	_SQL := "insert into video(title, store_path, thumbnail, description, is_charge, theme_id, publish_user_id) values(?,?,?,?,?,?,?)"
	_, err := db.Exec(_SQL, title, storePath, thumbnail, description, isCharge, themeID, publishUserID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteVideoBy(db db.MySQLExec, video_id int) error {
	_SQL := "delete from video where video_id = ?"

	_, err := db.Exec(_SQL, video_id)
	if err != nil {
		return err
	}

	return nil
}

func getVideoBy(db db.MySQLExec, filter string, args []interface{}) (*Video, error) {
	_SQL := "select video_id, title, store_path, thumbnail, description, is_charge, free_time, theme_id, publish_user_id from video %s limit 1 offset 0"

	if filter != "" {
		_SQL = fmt.Sprintf(_SQL, " where "+filter)
	} else {
		_SQL = fmt.Sprintf(_SQL, "")
	}

	video := &Video{}
	err := db.QueryRow(_SQL, args...).Scan(
		&video.VideoID,
		&video.Title,
		&video.StorePath,
		&video.Thumbnail,
		&video.Description,
		&video.IsCharge,
		&video.FreeTime,
		&video.ThemeID,
		&video.PublishUserID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return video, nil
}

func GetVideoByID(db db.MySQLExec, video_id int) (*Video, error){
	return getVideoBy(db, "video_id = ?", []interface{}{video_id})
}
