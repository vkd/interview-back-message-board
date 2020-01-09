package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vkd/interview-back-message-board/model"
)

// MessageListGetter - GetMessages method.
type MessageListGetter interface {
	GetMessages(order string) ([]model.Message, error)
}

func getListMessagesHandler(storage MessageListGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		order := c.Query("order")
		list, err := storage.GetMessages(order)
		if err != nil {
			errorResponse(c, fmt.Errorf("error on get list of messages: %w", err))
			return
		}

		okResponse(c, list)
	}
}
