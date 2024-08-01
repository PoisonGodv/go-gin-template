package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"test_wxlogin/models"
	"test_wxlogin/utils"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"code":    200,
		"message": data,
	})
}

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	rePassword := c.PostForm("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	if password != rePassword {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "两次密码不一样!",
		})
		return

	}
	data1 := models.FindUserByName(user.Name)
	fmt.Print(data1)
	if data1.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名不能重复!",
		})
		return
	}
	data2 := models.FindUserByPhone(user.Phone)
	fmt.Print(data2)
	if data2.Phone != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "手机号不能重复!",
		})
		return
	}
	data3 := models.FindUserByEmail(user.Email)
	fmt.Print(data3)
	if data3.Email != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "邮箱不能重复!",
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功!",
	})

}

func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除用户成功!",
	})
}

func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Print(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "修改参数不匹配!",
		})
		return
	} else {
		fmt.Print(user)
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "修改用户成功!",
		})
	}

}
