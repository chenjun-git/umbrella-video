package model

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chenjun-git/umbrella-video/common"
	"github.com/chenjun-git/umbrella-video/db"
)

func init() {
	common.InitConfig(configPath)
	db.InitMySQL(common.Config.Mysql)
}

func initTestVideoCaseEnv(t *testing.T) {
	initTestVideoTable(t)
}

func initTestVideoTable(t *testing.T) {
	assert := assert.New(t)

	f, err := ioutil.ReadFile(TestVideoSql)
	assert.Nil(err)
	assert.NotNil(f)
	result, err := db.MySQL.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", VideoTableName))
	assert.Nil(err)
	assert.NotEmpty(result)
	result, err = db.MySQL.Exec(string(f))
	assert.Nil(err)
	assert.NotEmpty(result)
}

func TestCreateGetDeleteVideo(t *testing.T) {
	assert := assert.New(t)
	initTestVideoCaseEnv(t)

	err := CreateVideo(db.MySQL, "test_title", "test_path", "test_thumbnail", "test_description", false, 1, "test_pushlish_user_id")
	assert.Nil(err)

	video, err := GetVideoByID(db.MySQL, 1)
	assert.Nil(err)
	assert.Equal("test_title", video.Title)

	err = DeleteVideoBy(db.MySQL, 1)
	assert.Nil(err)
}