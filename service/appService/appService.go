package appService

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"test_wxlogin/models/wxUser"
	"test_wxlogin/utils/jwt"
	"time"
)

type wxLoginResult struct {
	Session_key string
	Unionid     string
	Errmsg      string
	Openid      string
	Errcode     int32
}

func WxLogin(c *gin.Context) {

	//wxLoginCode := c.PostForm("code")

	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]string
	// 反序列化
	_ = json.Unmarshal(b, &m)
	wxLoginCode := m["code"]
	//向微信服务器发送请求获取openid
	params := url.Values{}
	Url, err := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	if err != nil {
		return
	}
	params.Set("appid", viper.GetString("wx-app.appId"))
	params.Set("secret", viper.GetString("wx-app.appSecret"))
	params.Set("js_code", wxLoginCode)
	params.Set("grant_type", "authorization_code")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.JSON(resp.StatusCode, gin.H{
				"code":    resp.StatusCode,
				"message": err,
			})
			return
		}
	}(resp.Body)
	if err != nil {
		c.JSON(resp.StatusCode, gin.H{
			"code":    resp.StatusCode,
			"message": err,
		})
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	var res wxLoginResult
	err = json.Unmarshal(body, &res)

	if err != nil || res.Errcode != 0 {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": err,
		})
	}
	//判断数据库中有没有该用户，没有则存入数据库，有则直接生成token

	loginUser := wxUser.FindByOpenId(res.Openid)

	if *(loginUser) == (wxUser.WxUser{}) {
		inserUser := wxUser.WxUser{
			Openid:    res.Openid,
			CreatTime: time.Now(),
			Img:       "https://lnw-test.oss-cn-chengdu.aliyuncs.com/4256ae4c-ce35-4ded-812e-d970eea80b92.png",
			LastLogin: time.Now(),
		}

		result := wxUser.InsertUser(inserUser)
		if result.Error != nil {
			c.JSON(-1, gin.H{
				"code":    -1,
				"message": result.Error,
			})
			return
		}
	} else {
		(*loginUser).LastLogin = time.Now()
		result := wxUser.UpdateUser(loginUser)
		if result.Error != nil {
			c.JSON(-1, gin.H{
				"code":    -1,
				"message": result.Error,
			})
			return
		}
	}

	//生成token返回res.openid
	token, err := jwtUtils.GenToken(res.Openid)
	if err != nil {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"token": token},
	})

	return
}

func TestMiddleware(c *gin.Context) {
	openid, _ := c.Get("openid")
	fmt.Println(openid, "openid")
	c.JSON(http.StatusOK, gin.H{

		"code": 200,
	})
}
