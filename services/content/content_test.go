package content_test

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
	"github.com/sundogrd/content-api/services/content"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/db"
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
	err := viper.ReadConfig(bytes.NewBuffer(jsonConfig))
	if err != nil {
		return err
	}
	_, err = db.Init()
	return err
}

func TestContentService_Create(t *testing.T) {
	var err error

	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentServiceInstance().Create(content.CreateRequest{
		Title:       "test",
		Description: "desc",
		AuthorID:    123,
		Category:    "cate",
		Type:        1,
		Body:        "## kekeke\n awa",
		Version:     1,
		Extra:       content.BaseInfoExtra{},
	})
	if err != nil {
		t.Fatalf("CreateContent err: %+v", err)
	}
	t.Logf("CreateContent: %+v", res)
}

func TestContentService_Delete(t *testing.T) {
	var err error

	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentServiceInstance().Delete(content.DeleteRequest{
		ContentIDs: []int64{303983183500677120},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("DeleteContent: %+v", res)
}

func TestContentService_Find(t *testing.T) {
	err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentServiceInstance().Find(content.FindRequest{})
	if err != nil {
		t.Fatalf("FindContent err: %+v", err)
	}
	t.Logf("FindContents: %+v", res)
}
func TestContentService_FindOne(t *testing.T) {
	err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentServiceInstance().FindOne(content.FindOneRequest{
		ContentID: 303983137602408448,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("FindContent: %+v", res)
}

func TestContentService_GetRecommendByContent(t *testing.T) {
	var err error
	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentServiceInstance().GetRecommendByContent(content.GetRecommendByContentRequest{
		ContentID: 299696981532479488,
	})
	if err != nil {
		t.Fatalf("RecommendContent err: %+v", err)
	}
	t.Logf("RecommendContent: %+v", res)
}

func TestContentService_Update(t *testing.T) {
	var err error

	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentServiceInstance().Update(content.UpdateRequest{
		Target:      content.SDContent{ContentID: 303983196138115072},
		Title:       "updateTest",
		Description: "descUpdated",
	})
	if err != nil {
		t.Fatalf("UpdateContent err: %+v", err)
	}
	t.Logf("UpdateContent: %+v", res)
}
