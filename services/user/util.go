package user

import "encoding/json"

func unmarshalUserExtraJson(jsonStr string) (*UserInfoExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra UserInfoExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalUserExtraJson(extra *UserInfoExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}

func packUserInfo(dbData SDUser) UserInfo {
	unmarshaledExtra, err := unmarshalUserExtraJson(dbData.Extra)
	if err != nil {
		unmarshaledExtra = &UserInfoExtra{}
	}
	return UserInfo{
		UserID:    dbData.UserID,
		Name:      dbData.Name,
		AvatarUrl: dbData.AvatarUrl,
		Company:   dbData.Company,
		Email:     dbData.Email,
		CreatedAt: dbData.CreatedAt,
		UpdatedAt: dbData.UpdatedAt,
		Extra:     *unmarshaledExtra,
	}
}
