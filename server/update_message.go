package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vkd/interview-back-message-board/model"
)

// MessageUpdater - UpdateMessage method.
type MessageUpdater interface {
	UpdateMessage(id string, m model.Message) (model.Message, error)
}

type updateMessageJSON struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Text  string `json:"text"`
}

func updateMessageHandler(storage MessageUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var j updateMessageJSON
		err := c.BindJSON(&j)
		if err != nil {
			errorResponse(c, fmt.Errorf("error on bind json: %w", err))
			return
		}
		_, err = storage.UpdateMessage(id, model.Message{
			Name:  j.Name,
			Email: j.Email,
			Text:  j.Text,
		})
		if err != nil {
			errorResponse(c, fmt.Errorf("error on update message: %w", err))
			return
		}

		okResponse(c, nil)
	}
}
