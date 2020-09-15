package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/models"
)

type TagSerializer struct {
	C *gin.Context
	models.Tag
}

type TagsSerializer struct {
	C    *gin.Context
	Tags []models.Tag
}

type TagResponse struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	State     int               `json:"state"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	Articles  []ArticleResponse `json:"articles"`
}

func (s *TagSerializer) Response() TagResponse {
	response := TagResponse{
		ID:        s.ID,
		CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Name:      s.Name,
		State:     s.State,
	}
	response.Articles = make([]ArticleResponse, 0)
	for _, article := range s.Articles {
		serializer := ArticleSerializer{s.C, article}
		response.Articles = append(response.Articles, serializer.Response())
	}

	return response
}

func (s *TagsSerializer) Response() []TagResponse {
	response := []TagResponse{}
	for _, tag := range s.Tags {
		serializer := TagSerializer{s.C, tag}
		response = append(response, serializer.Response())
	}
	return response
}
