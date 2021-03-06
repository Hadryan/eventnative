package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ksensehq/eventnative/appconfig"
	"net/http"
)

const TokenName = "token"

//TokenAuth check that provided token is valid and exists in auth config
func TokenAuth(main gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(appconfig.Instance.AuthorizedTokens) > 0 {
			queryValues := c.Request.URL.Query()
			token := queryValues.Get(TokenName)
			_, ok := appconfig.Instance.AuthorizedTokens[token]
			if !ok {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Set(TokenName, token)
		}
		main(c)
	}
}
