package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// create post_id
	p.ID = int64(snowflake.GenID())
	// ru ku
	//err = mysql.CreatePost(p)
	// return res
	return mysql.CreatePost(p)
}

//func GetPostHandler(id int64) (post *models.Post, err error) {
//	return mysql.GetPostByID(id)
//}

func GetPostHandler(id int64) (data *models.ApiPostDetail, err error) {
	// post detail
	//data = new(models.ApiPostDetail)
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostByID Failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserIDByID(post.AuthorID)
	if err != nil {
		zap.L().Error("GetUserIDByID(post.AuthorID)", zap.Error(err))
		return
	}
	// community detail
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID", zap.Error(err))
		return
	}
	//data.AuthorName = user.UserName
	//data.CommunityDetail = community
	//data.Post = post
	data = &models.ApiPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList() (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList()
	if err != nil {
		zap.L().Error("mysql.GetPostList Failed", zap.Error(err))
		return nil, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserIDByID(post.AuthorID)
		if err != nil {
			zap.L().Error("GetUserIDByID(post.AuthorID)", zap.Error(err))
			continue
		}
		// community detail
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID", zap.Error(err))
			continue
		}
		datadetail := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, datadetail)
	}

	return
}
