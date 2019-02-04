package content

import (
	"errors"
	"fmt"
)

// UpdateRequest ...
type UpdateRequest struct {
	Target SDContent

	Title       string
	Description string
	Category    string
	Type        int16
	Body        string
	//Extra       DataInfoExtra
}

// UpdateResponse ...
type UpdateResponse struct {
	ContentInfo
}

// Update ...
func (cr ContentService) Update(req UpdateRequest) (*UpdateResponse, error) {
	var target SDContent

	// TODO: 加上extra的update
	cr.db.Where("content_id = ?", req.Target.ContentID).Take(&target)
	var modified bool = false
	if req.Title != "" {
		target.Title = req.Title
		modified = true
	}
	if req.Description != "" {
		target.Description = req.Description
		modified = true
	}
	if req.Category != "" {
		target.Category = req.Category
		modified = true
	}
	// TODO: content_type自动解析赋值
	//if req.Type != 0 {
	//	target.Type = req.Type
	//	modified = true
	//}
	if req.Body != "" {
		target.Body = req.Body
		modified = true
	}

	if modified {
		target.Version += 1
	} else {
		return nil, errors.New("NoContentModified")
	}
	if dbc := cr.db.Save(&target); dbc.Error != nil {
		fmt.Printf("[services/content] Update: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	return &UpdateResponse{
		ContentInfo: packContentInfo(target),
	}, nil

}
