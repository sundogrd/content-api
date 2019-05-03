package oauth2

import "github.com/gin-gonic/gin"

func GithubOAuthMiddleware(config string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取或者设置cookie 中的session_id
		session := SessionStart(config, c.Request, c.Writer)
		c.Set("session", session)
		c.Next()
	}
}