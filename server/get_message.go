package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vkd/interview-back-message-board/model"
)

type MessageGetter interface {
	GetMessage(id string) (model.Message, bool, error)
}

func getMessageHandler(storage MessageGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		m, ok, err := storage.GetMessage(id)
		if err != nil {
			errorResponse(c, fmt.Errorf("error on get message by id: %w", err))
			return
		}

		if !ok {
			notFoundResponse(c)
			return
		}

		okResponse(c, m)
	}
}
