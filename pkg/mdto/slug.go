package mdto

type (
	SlugDto struct {
		Slug string `json:"slug" validate:"required"`
	}
	SlugExistenceResponse struct {
		Exists bool `json:"exists" validate:"required"`
	}
)
