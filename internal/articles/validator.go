package articles

import (
	"hafiztri123/be-sharing-vision/internal/utils"
	"slices"
)

func (dto *CreateArticleDTO) Validate() []utils.ValidationErrorPayload {
	var errors []utils.ValidationErrorPayload

	if len(dto.Title) < 20 {
		errors = append(errors, utils.ValidationErrorPayload{
			Message: "Title less than 20 characters",
			Key:     "title",
		})
	}

	if len(dto.Content) < 200 {
		errors = append(errors, utils.ValidationErrorPayload{
			Message: "Content less than 200 characters",
			Key:     "content",
		})

	}

	if len(dto.Category) < 3 {
		errors = append(errors, utils.ValidationErrorPayload{
			Message: "Category less than 3 characters",
			Key:     "category",
		})

	}

	status := []string{"publish", "draft", "thrash"}

	if !slices.Contains(status, dto.Status) {
		errors = append(errors, utils.ValidationErrorPayload{
			Message: "Status didn't fit one of these options: 'publish', 'draft' 'thrash'",
			Key:     "status",
		})

	}

	if len(errors) > 0 {
		return errors
	}

	return nil

}
