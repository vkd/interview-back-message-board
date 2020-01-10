package storage

import (
	"sort"
	"strconv"
	"time"

	"github.com/vkd/interview-back-message-board/model"
)

// SliceMessages - slice implementation of the message storage.
//
// TODO: make this structure as thread unsafe
type SliceMessages []model.Message

// GetMessage - get one message by id.
func (s *SliceMessages) GetMessage(id string) (model.Message, bool, error) {
	var m model.Message
	// TODO: not optimal search
	for _, m = range *s {
		if m.ID == id {
			return m, true, nil
		}
	}
	return m, false, nil
}

// GetMessages - get list of messages.
func (s *SliceMessages) GetMessages(_ string) ([]model.Message, error) {
	// TODO: use order param
	res := []model.Message(*s)
	// TODO: not optimal sort
	sort.Slice(res, func(i, j int) bool {
		return res[i].CreationTime.After(res[j].CreationTime)
	})
	return res, nil
}

// PostNewMessage - add new message.
func (s *SliceMessages) PostNewMessage(m model.Message) (model.Message, error) {
	m.ID = strconv.Itoa(len(*s))
	m.CreationTime = time.Now()
	*s = append(*s, m)
	return m, nil
}

// UpdateMessage - update message record.
func (s *SliceMessages) UpdateMessage(id string, updateMessage model.Message) (model.Message, bool, error) {
	var m model.Message
	// TODO: not optimal search
	for i, m := range *s {
		if m.ID == id {
			// requirements says "update the text of a specific message"
			// []model.Message(*s)[i].Name = updateMessage.Name
			// []model.Message(*s)[i].Email = updateMessage.Email
			[]model.Message(*s)[i].Text = updateMessage.Text
			return m, true, nil
		}
	}
	return m, false, nil
}

// ImportMessage - import message.
func (s *SliceMessages) ImportMessage(m model.Message) error {
	*s = append(*s, m)
	return nil
}
