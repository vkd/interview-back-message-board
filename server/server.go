package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: use admin users storage instead
var users = map[string]string{
	"admin": "back-challenge",
}

// Storager - storage of the messages.
type Storager interface {
	MessagePoster
	MessageListGetter
	MessageGetter
	MessageUpdater
}

// New - return new server instance.
func New(storage Storager) *gin.Engine {
	e := gin.New()

	publicAPI := e.Group("")
	publicAPI.POST("/message", postNewMessageHandler(storage))

	privateAPI := e.Group("", gin.BasicAuth(gin.Accounts(users)))
	messages := privateAPI.Group("/message")
	{
		messages.GET("", getListMessagesHandler(storage))
		messageID := messages.Group("/:id")
		{
			messageID.GET("", getMessageHandler(storage))
			messageID.POST("", updateMessageHandler(storage))
		}
	}

	return e
}

type status string

const (
	statusOK       status = "ok"
	statusError    status = "error"
	statusNotFound status = "not found"
)

func okResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, responseJSON{
		Status: statusOK,
		Data:   data,
	})
}

type responseJSON struct {
	Status status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func errorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, responseJSON{
		Status: statusError,
		Error:  err.Error(),
	})
}

func notFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, responseJSON{
		Status: statusNotFound,
	})
}
