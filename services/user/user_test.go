package user_test

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/sundogrd/content-api/services/user"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/db"
	"github.com/sundogrd/content-api/utils/pointer"
	"testing"
)

func initTestDB() error {
	config.Init()
	viper.SetConfigType("json") // or viper.SetConfigType("YAML")

	var jsonConfig = []byte(`{
	  	"db": {
			"type": "mysql",
			"options": {
				"user": "sundog",
				"password": "sundogPwd",
				"host": "localhost",
				"port": 3306,
				"dbname": "sundog",
				"connectTimeout": "10s"
			}
	  	}
	}`)
	viper.ReadConfig(bytes.NewBuffer(jsonConfig))
	_, err := db.Init()
	return err
}

func TestUserService_Create(t *testing.T) {
	var err error

	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := user.UserServiceInstance().Create(user.CreateRequest{
		Name:      "LWio",
		AvatarUrl: "https://avatars1.githubusercontent.com/u/9214496?v=4",
		Company:   pointer.String("Bytedance"),
		Email:     pointer.String("liang.peare@gmail.com"),
		Extra: user.UserInfoExtra{
			GithubHome: "https://github.com/lwyj123",
		},
	})
	if err != nil {
		t.Fatalf("CreateContent err: %+v", err)
	}
	t.Logf("[User] Create: %+v", res)
}

func TestUserService_Delete(t *testing.T) {
	var err error

	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := user.UserServiceInstance().Delete(user.DeleteRequest{
		UserID: 312461704938139648,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("[User] Delete: %+v", res)
}
func TestUserService_FindOne(t *testing.T) {
	err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res := user.UserServiceInstance().FindOne(user.FindOneRequest{
		UserID: pointer.Int64(312337740408565760),
	})
	t.Logf("[User] FindOne: %+v", res)
}
