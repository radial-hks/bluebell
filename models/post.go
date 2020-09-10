package models

import "time"

// n
type Post struct {
	ID          int64     `json:"id" form:"id" db:"id"`
	PostID      int64     `json:"post_id" form:"post_id" db:"post_id"`
	AuthorID    int64     `json:"author_id" form:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" form:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" form:"status" db:"status"`
	Title       string    `json:"title" form:"username" db:"title" binding:"required"`
	Content     string    `json:"content" form:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

// TIEIZI DEtail
type ApiPostDetail struct {
	AuthorName string `json:"user_name"`
	*Post
	*CommunityDetail `json:"communityDetail"`
}
