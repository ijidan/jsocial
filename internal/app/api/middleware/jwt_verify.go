package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ijidan/jsocial/internal/constant"
	"github.com/ijidan/jsocial/internal/pkg/response"
	"net/http"
)

func JwtVerify() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		ins, _ := response.GetResponseInstance()
		if token == "" {
			rsp := ins.JsonFail(constant.Unauthorized, constant.OK, "token required", nil, "")
			context.JSON(http.StatusOK, rsp)
		}
		//claim, err := jwt.ParseJwtToken(token, jwt.Secret)
		//if err != nil {
		//	rsp := ins.JsonFail(erro.Unauthorized, erro.OK, "token error", nil, "")
		//	context.JSON(http.StatusOK, rsp)
		//}

		//exp := cast.ToInt(claim["exp"])
		//if exp < time.Now().Second() {
		//	rsp := ins.JsonFail(erro.Unauthorized, erro.OK, "token expired", nil, "")
		//	context.JSON(http.StatusOK, rsp)
		//}
		//uid := cast.ToInt(claim["uid"])
		//context.Set("user_id", uid)

		context.Next()
	}
}
