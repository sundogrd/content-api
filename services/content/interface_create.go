package content

import "fmt"

// CreateRequest ...
type CreateRequest struct {
	Title       string
	Description string
	AuthorID    int64
	Category    string
	Type        ContentType
	Body        string
	Extra       BaseInfoExtra
}

// CreateResponse ...
type CreateResponse struct {
	BaseInfo
}

// Create ...
func (cr ContentService) Create(req CreateRequest) (*CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	contentExtraStr, err := marshalContentExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/content] Create: json marshal error: %+v", err)
		contentExtraStr, _ = marshalContentExtraJson(&BaseInfoExtra{})
	}

	contentId, _ := cr.idBuilder.NextId()
	// TODO: contentType自动解析赋值

	content := SDContent{
		ContentID:   contentId,
		Title:       req.Title,
		Description: req.Description,
		AuthorID:    req.AuthorID,
		Category:    req.Category,
		Type:        1, // 先写死只有图文
		Body:        req.Body,
		BodyType:    3, // 先写死为Markdown
		Version:     1,
		Extra:       contentExtraStr,
	}
	if dbc := cr.db.Create(&content); dbc.Error != nil {
		fmt.Printf("[services/content] Create: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	responseExtra, err := unmarshalContentExtraJson(content.Extra)
	if err != nil {
		fmt.Printf("[services/content] Create: UnmarshalContentJson error: %+v", err)
		responseExtra = &BaseInfoExtra{}
	}
	res := &CreateResponse{
		BaseInfo: BaseInfo{
			ContentID:   content.ContentID,
			Title:       content.Title,
			Description: content.Description,
			AuthorID:    content.AuthorID,
			Category:    content.Category,
			Type:        1, // 写死只有图文
			Body:        content.Body,
			BodyType:    3, // 先写死为Markdown
			Version:     content.Version,
			CreatedAt:   content.CreatedAt,
			UpdatedAt:   content.UpdatedAt,
			Extra:       *responseExtra,
		},
	}
	return res, nil
}
