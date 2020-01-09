package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vkd/interview-back-message-board/model"
)

// MessagePoster - PostNewMessage method.
type MessagePoster interface {
	PostNewMessage(m model.Message) (model.Message, error)
}

type newMessageJSON struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Text  string `json:"text"`
}

func postNewMessageHandler(storage MessagePoster) gin.HandlerFunc {
	return func(c *gin.Context) {
		var j newMessageJSON
		err := c.BindJSON(&j)
		if err != nil {
			errorResponse(c, fmt.Errorf("error on parse request: %w", err))
			return
		}

		m := model.Message{
			Name:  j.Name,
			Email: j.Email,
			Text:  j.Text,
		}
		newM, err := storage.PostNewMessage(m)
		if err != nil {
			errorResponse(c, fmt.Errorf("error on write new message: %w", err))
			return
		}

		okResponse(c, newM.ID)
	}
}
