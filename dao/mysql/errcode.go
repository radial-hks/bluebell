package mysql

import "errors"

var (
	ErrorCommunityNotExist  = errors.New("CommunityNotExist")
	ErrorCommunityInvalidID = errors.New("CommunityInvalidID ")
)
