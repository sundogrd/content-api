package content

import (
	"encoding/json"
)

func unmarshalContentExtraJson(jsonStr string) (*BaseInfoExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra BaseInfoExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalContentExtraJson(extra *BaseInfoExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}

func packBaseInfo(dbData SDContent) BaseInfo {
	unmarshaledExtra, err := unmarshalContentExtraJson(dbData.Extra)
	if err != nil {
		unmarshaledExtra = &BaseInfoExtra{}
	}
	return BaseInfo{
		ContentID:   dbData.ContentID,
		Title:       dbData.Title,
		Description: dbData.Description,
		AuthorID:    dbData.AuthorID,
		Category:    dbData.Category,
		Type:        dbData.Type,
		Body:        dbData.Body,
		BodyType:    dbData.BodyType,
		Version:     dbData.Version,
		CreatedAt:   dbData.CreatedAt,
		UpdatedAt:   dbData.UpdatedAt,
		Extra:       *unmarshaledExtra,
	}
}
