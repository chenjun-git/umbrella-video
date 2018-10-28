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
	fmt.Printf("%+v\n", common.Config.Mysql)
	db.InitMySQL(common.Config.Mysql)
}

func initTestVideoThemeCaseEnv(t *testing.T) {
	initTestVideoThemeTable(t)
}

func initTestVideoThemeTable(t *testing.T) {
	assert := assert.New(t)

	f, err := ioutil.ReadFile(TestVideoThemeSql)
	assert.Nil(err)
	assert.NotNil(f)
	result, err := db.MySQL.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s;", VideoThemeTableName))
	assert.Nil(err)
	assert.NotEmpty(result)
	result, err = db.MySQL.Exec(string(f))
	assert.Nil(err)
	assert.NotEmpty(result)
}

func TestCreateGetDeleteVideoTheme(t *testing.T) {
	assert := assert.New(t)
	initTestVideoThemeCaseEnv(t)

	err := CreateVideoTheme(db.MySQL, "test_name", "test_description")
	assert.Nil(err)

	videoTheme , err := GetVideoThemeByName(db.MySQL, "test_name")
	assert.Equal("test_name", videoTheme.Name)
	assert.Equal("test_description", videoTheme.Description)

	err = DeleteVideoThemeByName(db.MySQL, "test_name")
	assert.Nil(err)
}