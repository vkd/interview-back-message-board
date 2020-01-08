package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vkd/interview-back-message-board/model"
)

var (
	testMessageID           = "test-ID"
	testMessageName         = "test-name"
	testMessageEmail        = "test-email"
	testMessageText         = "test-text"
	testMessageCteationTime = time.Date(2020, 1, 1, 14, 0, 0, 0, time.UTC)
)

type testStorage []model.Message

func (s *testStorage) PostNewMessage(m model.Message) (model.Message, error) {
	m.ID = testMessageID
	m.CreationTime = testMessageCteationTime
	*s = append(*s, m)
	return m, nil
}

func Test_PostNewMessageHandler(t *testing.T) {
	var storage testStorage
	h := postNewMessageHandler(&storage)

	req, err := http.NewRequest("POST", "/message", strings.NewReader(`{
		"name": "`+testMessageName+`",
		"email": "`+testMessageEmail+`",
		"text": "`+testMessageText+`"
	}`))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	require.Len(t, storage, 0)
	h(ctx)
	require.Len(t, storage, 1)

	m := []model.Message(storage)[0]
	assert.Equal(t, testMessageID, m.ID)
	assert.Equal(t, testMessageName, m.Name)
	assert.Equal(t, testMessageEmail, m.Email)
	assert.Equal(t, testMessageText, m.Text)
	assert.Equal(t, testMessageCteationTime, m.CreationTime)
}
