package content

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// ReadRequest ...
type ReadRequest struct {
	ContentID int64
	//Extra       DataInfoExtra
}

// ReadResponse ...
type ReadResponse struct {
	ContentID int64
	ReadCount int64
}

// Update ...
func (cs ContentService) Read(req ReadRequest) (*ReadResponse, error) {
	var target SDContentCount

	fmt.Printf("[services/content] Read %d content", req.ContentID)
	// TODO: 加上extra的update
	if err := cs.db.Where("content_id = ? and count_key = 'read_count'", req.ContentID).First(&target).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) == true {
			// error handling...
			target = SDContentCount{
				ContentID: req.ContentID,
				CountKey: "read_count",
				CountValue: 1,
			}
			cs.db.Create(&target)  // newUser not user
		}
	} else {
		cs.db.Model(&target).Where("content_id = ? and count_key = 'read_count'", req.ContentID).Update("count_value", target.CountValue + 1)
	}

	return &ReadResponse{
		ContentID: req.ContentID,
		ReadCount: target.CountValue,
	}, nil

}
