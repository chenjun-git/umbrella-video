package model

import (
	"database/sql"
	"fmt"

	"github.com/chenjun-git/umbrella-video/db"
)

const (
	VideoThemeTableName = "video_theme"

	VideoThemeThemeID     = "theme_id"
	VideoThemeName        = "name"
	VideoThemeDescription = "description"
)

type VideoTheme struct {
	ThemeID     int
	Name        string
	Description string
}

func CreateVideoTheme(db db.MySQLExec, name, description string) error {
	_SQL := "insert into video_theme(name, description) values(?,?)"
	_, err := db.Exec(_SQL, name, description)
	if err != nil {
		return err
	}

	return nil
}

func DeleteVideoThemeById(db db.MySQLExec, id int) error {
	_SQL := "delete from video_theme where id = ?"

	_, err := db.Exec(_SQL, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteVideoThemeByName(db db.MySQLExec, name string) error {
	_SQL := "delete from video_theme where name = ?"

	_, err := db.Exec(_SQL, name)
	if err != nil {
		return err
	}

	return nil
}

func getVideoThemeBy(db db.MySQLExec, filter string, args []interface{}) (*VideoTheme, error) {
	_SQL := "select theme_id, name, description from video_theme %s limit 1 offset 0"

	if filter != "" {
		_SQL = fmt.Sprintf(_SQL, " where "+filter)
	} else {
		_SQL = fmt.Sprintf(_SQL, "")
	}

	videoTheme := &VideoTheme{}
	err := db.QueryRow(_SQL, args...).Scan(
		&videoTheme.ThemeID,
		&videoTheme.Name,
		&videoTheme.Description,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return videoTheme, nil
}

func GetVideoThemeById(db db.MySQLExec, id int) (*VideoTheme, error) {
	return getVideoThemeBy(db, "theme_id = ?", []interface{}{id})
}

func GetVideoThemeByName(db db.MySQLExec, name string) (*VideoTheme, error) {
	return getVideoThemeBy(db, "name = ?", []interface{}{name})
}
