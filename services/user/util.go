package user

import "encoding/json"

func UnmarshalUserExtraJson(jsonStr string) (*DataInfoExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra DataInfoExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalUserExtraJson(extra *DataInfoExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}

func sdUserToData(dbData SDUser) DataInfo {
	unmarshaled, err := UnmarshalUserExtraJson(dbData.Extra)
	if err != nil {
		unmarshaled = &DataInfoExtra{}
	}
	return DataInfo{
		ID:        dbData.ID,
		UserID:    dbData.UserID,
		Name:      dbData.Name,
		AvatarUrl: dbData.AvatarUrl,
		Company:   dbData.Company,
		Email:     dbData.Email,
		CreatedAt: dbData.CreatedAt,
		UpdatedAt: dbData.UpdatedAt,
		Extra:     *unmarshaled,
	}
}

func sdUsersToDatas(dbData []SDUser) []DataInfo {
	res := make([]DataInfo, 0)
	for _, sdUser := range dbData {
		res = append(res, sdUserToData(sdUser))
	}
	return res
}

func dataToSDUser(dataInfo DataInfo) SDUser {
	marshaled, _ := marshalUserExtraJson(&dataInfo.Extra)
	return SDUser{
		ID:        dataInfo.ID,
		UserID:    dataInfo.UserID,
		Name:      dataInfo.Name,
		AvatarUrl: dataInfo.AvatarUrl,
		Company:   dataInfo.Company,
		Email:     dataInfo.Email,
		CreatedAt: dataInfo.CreatedAt,
		UpdatedAt: dataInfo.UpdatedAt,
		Extra:     marshaled,
	}
}
