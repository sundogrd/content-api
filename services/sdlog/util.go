package sdlog

import "encoding/json"

func unmarshalSDLogExtraJson(jsonStr string) (*SDLogExtra, error) {
	var jsonBlob = []byte(jsonStr)
	var extra SDLogExtra
	err := json.Unmarshal(jsonBlob, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

func marshalSDLogExtraJson(extra *SDLogExtra) (string, error) {
	marshaled, err := json.Marshal(extra)
	if err != nil {
		return "{}", err
	}
	return string(marshaled), nil
}
