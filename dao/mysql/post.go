package mysql

import (
	"bluebell/models"
	"fmt"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO  post(post_id,title,content,author_id,
				community_id) values(?,?,?,?,?)`
	_, err = DB.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	// sql
	sqlStr := `select
      post_id,
      title,
      content,
      author_id,
      community_id,
      status,
      create_time,
      update_time
      from post WHERE post_id = ?`
	err = DB.Get(post, sqlStr, id)
	fmt.Println(post)
	return
}

func GetPostList() (posts []*models.Post, err error) {

	posts = make([]*models.Post, 0, 2)
	sqlStr := `select
      post_id,
      title,
      content,
      author_id,
      community_id,
      status,
      create_time,
      update_time
      from post
		limit 2`
	//err = DB.Get(post, sqlStr, 0)
	err = DB.Select(&posts, sqlStr)
	return
}
