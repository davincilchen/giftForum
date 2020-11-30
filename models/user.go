package models

import "giftForum/basemodels"

type LoginUser struct {
	basemodels.User
	UUID string
}
