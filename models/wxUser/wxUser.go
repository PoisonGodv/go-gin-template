package wxUser

import (
	"gorm.io/gorm"
	"test_wxlogin/utils"
	"time"
)

type WxUser struct {
	gorm.Model
	Openid    string
	Phone     string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Img       string //头像
	ClientIp  string
	LastLogin time.Time
	CreatTime time.Time
}

func (WxUser) TableName() string {
	return "appuser_info"
}

func FindByOpenId(openid string) *WxUser {
	wxUser := WxUser{}
	utils.DB.Where("openid = ?", openid).First(&wxUser)
	return &wxUser
}

func InsertUser(user WxUser) *gorm.DB {
	return utils.DB.Create(&user)
}
func UpdateUser(user *WxUser) *gorm.DB {
	return utils.DB.Save(user)
}
