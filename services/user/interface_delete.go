package user

import "fmt"

// DeleteRequest ...
type DeleteRequest struct {
	UserID int64
}

// DeleteResponse ...
type DeleteResponse struct {
	UserInfo
}

// Delete ...
func (us UserService) Delete(req DeleteRequest) (*DeleteResponse, error) {
	var user SDUser
	if dbc := us.db.Where("user_id = (?)", req.UserID).Delete(&user); dbc.Error != nil {
		fmt.Printf("[services/user] Delete: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	} else {
		return &DeleteResponse{
			UserInfo: packUserInfo(user),
		}, nil
	}
}
