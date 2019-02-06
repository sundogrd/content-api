package content

import "fmt"

type GetRecommendByContentRequest struct {
	ContentID int64
}
type GetRecommendByContentResponse struct {
	ContentList []BaseInfo
}

func (cr ContentService) GetRecommendByContent(req GetRecommendByContentRequest) (*GetRecommendByContentResponse, error) {
	var recommendContents []SDContent
	if dbc := cr.db.Limit(4).Order("updated_at desc").Find(&recommendContents); dbc.Error != nil {
		fmt.Printf("[services/content] GetRecommendByContent: db error: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}
	BaseInfos := make([]BaseInfo, 0)
	for _, v := range recommendContents {
		BaseInfos = append(BaseInfos, packBaseInfo(v))
	}
	return &GetRecommendByContentResponse{
		ContentList: BaseInfos,
	}, nil
}
