package content

import (
	"encoding/json"
)

func unmarshalContentExtraJson(jsonStr string) (*ContentInfoExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra ContentInfoExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalContentExtraJson(extra *ContentInfoExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}

func packContentInfo(dbData SDContent) ContentInfo {
	unmarshaledExtra, err := unmarshalContentExtraJson(dbData.Extra)
	if err != nil {
		unmarshaledExtra = &ContentInfoExtra{}
	}
	return ContentInfo{
		ID:          dbData.ID,
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
