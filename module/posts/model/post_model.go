package model

import (
	"time"
)

type Post struct {
	PostID      string    `json:"post_id"`
	PostImage   string    `json:"post_image"`
	PostSubject string    `json:"post_subject"`
	PostContent string    `json:"post_content"`
	Author      Author    `json:"author"`
	Tags        []string  `json:"tags"`
	IsPublish   int       `json:"isPublish"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Author struct {
	AuthorID           string `json:"author_id"`
	AuthorFullName     string `json:"author_fullname"`
	AuthorPhotoProfile string `json:"author_photo_profile"`
}

type Posts []Post

// NewPost func
func NewPost() *Post {
	return &Post{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
