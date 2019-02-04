package content

import "fmt"

// DeleteRequest ...
type DeleteRequest struct {
	ContentID int64
}

// DeleteResponse ...
type DeleteResponse struct {
	ContentInfo
}

// Delete ...
func (cr ContentService) Delete(req DeleteRequest) (*DeleteResponse, error) {
	content := &SDContent{}
	if dbc := cr.db.Where("content_id = (?)", req.ContentID).Delete(content); dbc.Error != nil {
		fmt.Printf("[services/content] Delete: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	} else {
		return &DeleteResponse{
			ContentInfo: packContentInfo(*content),
		}, nil
	}

}
