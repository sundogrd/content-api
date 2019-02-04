package content

import "fmt"

// DeleteRequest ...
type DeleteRequest struct {
	ContentIDs []int64
}

// DeleteResponse ...
type DeleteResponse struct {
	ContentInfo
}

// Delete ...
func (cr ContentService) Delete(req DeleteRequest) (*DeleteResponse, error) {
	content := &SDContent{}
	if dbc := cr.db.Where("content_id IN (?)", req.ContentIDs).Delete(content); dbc.Error != nil {
		fmt.Printf("[services/content] Delete: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	} else {
		fmt.Printf("%+v\n", dbc)
	}
	return &DeleteResponse{
		ContentInfo: packContentInfo(*content),
	}, nil
}
