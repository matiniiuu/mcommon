package mdto

import "github.com/matiniiuu/mcommon/pkg/menums"

type (
	ID struct {
		ID string `param:"id" validate:"required"`
	}
	HeaderLanguage struct {
		Language *string `reqHeader:"Accept-Language"`
	}

	ListRequest struct {
		HeaderLanguage
		Sort       menums.SortOrder `query:"sort" validate:"required,sortOrder"`
		Page       int64            `query:"page" validate:"required"`
		Limit      int64            `query:"limit" validate:"required"`
		SearchText string           `query:"searchText"`
	}
	SuccessResponse struct {
		Message string `json:"message"`
	}

	ListResponse[R any] struct {
		Result []R   `json:"result"`
		Total  int64 `json:"total"`
	}
	DataResponse[R any] struct {
		Data R `json:"data"`
	}
	RemoveManyRequest struct {
		Ids []string `param:"ids" validate:"required"`
	}
	StringArray struct {
		Values []string `json:"values"`
	}
)

func OkResponse() *SuccessResponse {
	return &SuccessResponse{Message: "OK"}
}
