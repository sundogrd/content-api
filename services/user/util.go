package user

import "encoding/json"

func unmarshalUserExtraJson(jsonStr string) (*BaseInfoExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra BaseInfoExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalUserExtraJson(extra *BaseInfoExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}

func packBaseInfo(dbData SDUser) BaseInfo {
	unmarshaledExtra, err := unmarshalUserExtraJson(dbData.Extra)
	if err != nil {
		unmarshaledExtra = &BaseInfoExtra{}
	}
	return BaseInfo{
		UserID:    dbData.UserID,
		Name:      dbData.Name,
		AvatarURL: dbData.AvatarURL,
		Company:   dbData.Company,
		Email:     dbData.Email,
		CreatedAt: dbData.CreatedAt,
		UpdatedAt: dbData.UpdatedAt,
		Extra:     *unmarshaledExtra,
	}
}
