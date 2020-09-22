package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/models"
)

type ArticleSerializer struct {
	C *gin.Context
	models.Article
}

type ArticlesSerializer struct {
	C        *gin.Context
	Articles []models.Article
}

type ArticleResponse struct {
	ID            uint        `json:"id"`
	Title         string      `json:"title"`
	Desc          string      `json:"desc"`
	State         int         `json:"state"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	Tag           TagResponse `json:"tag"`
	CoverImageUrl string      `json:"cover_image_url"`
}

func (s *ArticleSerializer) Response() ArticleResponse {
	tagSerializer := TagSerializer{s.C, s.Tag}
	response := ArticleResponse{
		ID:            s.ID,
		CreatedAt:     s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt:     s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Title:         s.Title,
		Desc:          s.Desc,
		State:         s.State,
		CoverImageUrl: s.CoverImageUrl,
		Tag:           tagSerializer.Response(),
	}
	return response
}

func (s *ArticlesSerializer) Response() []ArticleResponse {
	response := []ArticleResponse{}
	for _, article := range s.Articles {
		serializer := ArticleSerializer{s.C, article}
		response = append(response, serializer.Response())
	}
	return response
}
