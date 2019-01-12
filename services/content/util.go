package content

import (
	"encoding/json"
)

func UnmarshalContentExtraJson(jsonStr string) (*DataInfoExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra DataInfoExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalContentExtraJson(extra *DataInfoExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}

func sdContentToData(dbData SDContent) DataInfo {
	unmarshaled, err := UnmarshalContentExtraJson(dbData.Extra)
	if err != nil {
		unmarshaled = &DataInfoExtra{}
	}
	return DataInfo{
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
		Extra:       *unmarshaled,
	}
}

func sdContentsToDatas(dbData []SDContent) []DataInfo {
	res := make([]DataInfo, 0)
	for _, sdContent := range dbData {
		res = append(res, sdContentToData(sdContent))
	}
	return res
}

func dataToSDContent(dataInfo DataInfo) SDContent {
	marshaled, _ := marshalContentExtraJson(&dataInfo.Extra)
	return SDContent{
		ID:          dataInfo.ID,
		ContentID:   dataInfo.ContentID,
		Title:       dataInfo.Title,
		Description: dataInfo.Description,
		AuthorID:    dataInfo.AuthorID,
		Category:    dataInfo.Category,
		Type:        dataInfo.Type,
		Body:        dataInfo.Body,
		Version:     dataInfo.Version,
		CreatedAt:   dataInfo.CreatedAt,
		UpdatedAt:   dataInfo.UpdatedAt,
		Extra:       marshaled,
	}
}
