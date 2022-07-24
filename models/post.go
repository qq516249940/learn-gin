// models/post.go
package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"` // `json:"id" gorm:"primaryKey"` 这些内容称为标签，使用反引号注释来定义键值对
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
