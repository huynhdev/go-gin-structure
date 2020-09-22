package models

type Article struct {
	Model

	TagID         int `json:"tag_id"`
	Tag           Tag
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         int    `json:"state"`
}

func GetArticles(pageNum int, pageSize int) ([]Article, error) {
	var (
		articles []Article
		err      error
	)

	err = db.Joins("Tag").Limit(pageSize).Offset(pageNum).Find(&articles).Error

	return articles, err
}

func GetArticle(id string) (Article, error) {
	var (
		article Article
		err     error
	)

	err = db.Joins("Tag").First(&article, id).Error
	return article, err
}

func AddArticle(data map[string]interface{}) (Article, error) {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	err := db.Create(&article).Joins("Tag").Find(&article).Error
	return article, err
}

func EditArticle(id string, data map[string]interface{}) (Article, error) {
	var article Article
	err := db.Joins("Tag").First(&article, id).Updates(data).Error
	return article, err
}

func DeleteArticle(id string) error {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	return err
}
