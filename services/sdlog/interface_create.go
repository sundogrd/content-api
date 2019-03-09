package sdlog

import (
	"fmt"
	"time"
)

// CreateRequest ...
type CreateRequest struct {
	TargetID  int64
	UserID    int64
	Type      SDLogType
	Extra     SDLogExtra
	CreatedAt time.Time
}

// CreateResponse ...
type CreateResponse struct {
	SDLog
}

// Create ...
func (ss SDLogService) Create(req CreateRequest) (*CreateResponse, error) {

	sdlogExtraStr, err := marshalSDLogExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/sdlog] Create: json marshal error: %+v", err)
		sdlogExtraStr, _ = marshalSDLogExtraJson(&SDLogExtra{})
	}

	sdlogId, _ := ss.idBuilder.NextId()

	sdlog := SDLogModel{
		ID:       sdlogId,
		TargetID: req.TargetID,
		UserID:   req.UserID,
		Type:     string(req.Type),
		Extra:    sdlogExtraStr,
	}
	if dbc := ss.db.Create(&sdlog); dbc.Error != nil {
		fmt.Printf("[services/sdlog] Create: db createerror: %+v", dbc.Error)
		return nil, dbc.Error
	}

	responseExtra, err := unmarshalSDLogExtraJson(sdlog.Extra)
	if err != nil {
		fmt.Printf("[services/sdlog] Create: marshalSDLogExtraJson error: %+v", err)
		responseExtra = &SDLogExtra{}
	}
	res := &CreateResponse{
		SDLog: SDLog{
			LogID:    sdlog.ID,
			TargetID: sdlog.TargetID,
			UserID:   sdlog.UserID,
			Type:     SDLogType(sdlog.Type),
			Extra:    *responseExtra,
		},
	}
	return res, err
}
