package models

import "time"

type Community struct {
	CommunityID   int64  `db:"community_id" json:"id"`
	CommunityName string `db:"community_name" json:"name"`
}

type CommunityDetail struct {
	CommunityID   int64     `db:"community_id" json:"id"`
	CommunityName string    `db:"community_name" json:"name"`
	Introduction  string    `db:"introduction" json:"introduction,omitempty"`
	CreateTime    time.Time `db:"create_time" json:"create_time"`
	UpdateTime    time.Time `db:"udpate_time" json:"update_time"`
}
