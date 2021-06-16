package test

import (
	"test_no_1/model"
	"test_no_1/question"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("single model", func(t *testing.T) {
		input := model.DataRequest[0]
		resp := question.ConverterRequestToModel(&input)

		assert.Equal(t, resp.ID, input.ID)
		assert.Equal(t, resp.Title, input.Title)
		assert.Equal(t, resp.Lead, input.Lead)
		assert.Equal(t, resp.Content, input.Content)
		assert.Equal(t, resp.ImageCoverURL, input.ImageCoverURL)
		assert.Equal(t, resp.ImageCoverCaption, *input.ImageCoverCaption)
		assert.Equal(t, resp.Status, *input.Status)
		assert.Equal(t, resp.TagsDetail[0].ID, input.Tag1ID)
		assert.Equal(t, resp.TagsDetail[0].Name, input.Tag1Name)
		assert.Equal(t, resp.CategoryDetail.ID, input.CategoryDetailID)
		assert.Equal(t, resp.CategoryDetail.Name, input.CategoryDetailName)
	})

	t.Run("multiple model", func(t *testing.T) {
		var input []*model.ArticleRequest
		for _, row := range model.DataRequest {
			input = append(input, &row)
		}
		multipleResp := question.ConverterMultipleRequestToModel(input)

		for i, resp := range multipleResp {
			assert.Equal(t, resp.ID, input[i].ID)
			assert.Equal(t, resp.Title, input[i].Title)
			assert.Equal(t, resp.Lead, input[i].Lead)
			assert.Equal(t, resp.Content, input[i].Content)
			assert.Equal(t, resp.ImageCoverURL, input[i].ImageCoverURL)
			assert.Equal(t, resp.ImageCoverCaption, *input[i].ImageCoverCaption)
			assert.Equal(t, resp.Status, *input[i].Status)
			assert.Equal(t, resp.TagsDetail[0].ID, input[i].Tag1ID)
			assert.Equal(t, resp.TagsDetail[0].Name, input[i].Tag1Name)
			assert.Equal(t, resp.CategoryDetail.ID, input[i].CategoryDetailID)
			assert.Equal(t, resp.CategoryDetail.Name, input[i].CategoryDetailName)
		}
	})
}
