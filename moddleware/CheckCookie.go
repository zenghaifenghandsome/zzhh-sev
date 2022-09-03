package moddleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("username"); err == nil {
			fmt.Println("进来了")
			fmt.Println(cookie)
			c.JSON(http.StatusUnauthorized, gin.H{"message": cookie})
			c.Next()
		}
		// 返回错误
		c.JSON(http.StatusOK, gin.H{"message": "cookieErr"})
		// 若验证不通过，不再调用后续的函数处理
		//c.Abort()
	}
}
