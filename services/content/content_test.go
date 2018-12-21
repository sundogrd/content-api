package content_test

import (
	"bytes"
	"context"
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
	viper.ReadConfig(bytes.NewBuffer(jsonConfig))
	_, err := db.Init()
	return err
}

// TestContentFindOne ...
func TestContentFindOne(t *testing.T) {
	ctx := context.Background()
	err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res := content.ContentRepositoryInstance().FindOne(ctx, content.FindOneRequest{
		ContentID: 300746136048635904,
	})
	t.Logf("FindContent: %+v", res)
}

// TestContentFind ...
func TestContentFind(t *testing.T) {
	ctx := context.Background()
	err := initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentRepositoryInstance().Find(ctx, content.FindRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("FindContents: %+v", res)
}


// TestContentCreate ...
func TestContentCreate(t *testing.T) {
	var err error

	ctx := context.Background()
	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentRepositoryInstance().Create(ctx, content.CreateRequest{
		Title:       "test",
		Description: "desc",
		AuthorID:    123,
		Category:    "cate",
		Type:        0,
		Body:        "## kekeke\n awa",
		Version:     1,
		Extra:       content.DataInfoExtra{},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

// TestContentDelete ...
func TestContentDelete(t *testing.T) {
	var err error

	ctx := context.Background()
	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentRepositoryInstance().Delete(ctx, content.DeleteRequest{
		ContentID: 299696847465746432,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

// TestContentUpdate ...
func TestContentUpdate(t *testing.T) {
	var err error

	ctx := context.Background()
	err = initTestDB()
	if err != nil {
		t.Fatal(err)
	}
	res, err := content.ContentRepositoryInstance().Update(ctx, content.UpdateRequest{
		Target:      content.PFContent{ContentID: 299696981532479488},
		Title:       "updateTest",
		Description: "descUpdated",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
