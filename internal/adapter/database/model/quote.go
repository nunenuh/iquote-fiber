package model

import (
	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model

	ID           int        `gorm:"primaryKey" json:"id"`
	QText        string     `gorm:"column:qtext;type:varchar(255);not null" json:"text"`
	Tags         string     `gorm:"column:tags;type:varchar(255);not null" json:"tags"`
	AuthorID     *int       `json:"author_id"`
	Author       Author     `json:"author"`
	Categories   []Category `gorm:"many2many:quote_categories"`
	UserWhoLiked []User     `gorm:"many2many:quote_likes" json:"users_who_liked"`
}

func (q *Quote) LikedCount() int {
	return len(q.UserWhoLiked)
}
func (q *Quote) IsDeleted() bool {
	return !q.DeletedAt.Time.IsZero()
}
