package menums

import "github.com/go-playground/validator/v10"

type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

func (sort SortOrder) ConvertToMongoSort() int {
	if sort == Asc {
		return 1
	}
	if sort == Desc {
		return -1
	}
	return 1
}
func ValidateSortOrder(fl validator.FieldLevel) bool {
	sortOrder := SortOrder(fl.Field().String())
	return sortOrder == Asc || sortOrder == Desc
}
