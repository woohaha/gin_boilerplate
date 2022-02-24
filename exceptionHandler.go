package boilerplate

import (
	"log"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *gin.Context) error

func HandleException(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			var exp *APIException
			if h, ok := err.(*APIException); ok {
				exp = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					exp = UnknownError(e.Error())
				} else {
					log.Println(err.Error())
					exp = ServerError()
				}
			} else {
				log.Println(err.Error())
				exp = ServerError()
			}
			exp.setRequestURI(c.Request)
			c.JSON(exp.Code, exp)
			return
		}
	}
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	handleErr.setRequestURI(c.Request)
	c.JSON(handleErr.Code, handleErr)
	return
}
