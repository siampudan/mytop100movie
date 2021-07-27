package apperror

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Message string
	Status  int
}

func New(status int, message string) *AppError {
	return &AppError{Status: status, Message: message}
}

func (apperr *AppError) Error() string {
	return apperr.Message
}

func Response(c *gin.Context, err error) {
	switch err.(type) {
	case *AppError:
		appErr := err.(*AppError)
		if appErr.Message == "" {
			c.AbortWithStatus(appErr.Status)
		} else {
			c.AbortWithStatusJSON(appErr.Status, appErr)
		}
		return
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
