package middleware

import (
	"time"
	"todo_list/pkg/e"
	"todo_list/pkg/utils"

	"github.com/gin-gonic/gin"
)

// gin框架中间价的标准形式 gin.HandlerFunc
func JWT() gin.HandlerFunc {
	//用于访问参数和响应数据
	return func(c *gin.Context) {
		//注意初始化
		code := e.SUCCESS
		// 获取请求的http，头部字段通常包含JWT
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthEmptyToken
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout //token失效
			}
		}
		if code != e.SUCCESS {
			//响应客户端json对象
			c.JSON(e.SUCCESS, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			//中端处理流程
			c.Abort()
			return
		}
		//若令牌有效则执行下一个中间件
		c.Next()
	}
}
