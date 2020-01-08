package imporer

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vkd/interview-back-message-board/model"
)

type testMessageStorage []model.Message

func (s *testMessageStorage) ImportMessage(m model.Message) error {
	*s = append(*s, m)
	return nil
}

func TestImportMessages(t *testing.T) {
	csvContent := `id,name,email,text,creation_time
2C7BCEC7-CD14-D6E5-3FBF-F95513754290,test_name0,testemail0@test.email,"test,message0",2017-12-14T06:20:33-08:00
2C7BCEC7-CD14-D6E5-3FBF-F95513754291,test_name1,testemail1@test.email,"test,message1",2017-12-14T06:20:33-08:00
2C7BCEC7-CD14-D6E5-3FBF-F95513754292,test_name2,testemail2@test.email,"test,message2",2017-12-14T06:20:33-08:00
`
	testCreatedTime, err := time.Parse(time.RFC3339, "2017-12-14T06:20:33-08:00")
	require.NoError(t, err)

	var testStorage testMessageStorage

	err = ImportMessages(strings.NewReader(csvContent), &testStorage)
	require.NoError(t, err)
	require.Len(t, testStorage, 3)

	for i, m := range ([]model.Message)(testStorage) {
		si := strconv.Itoa(i)
		assert.Equal(t, "2C7BCEC7-CD14-D6E5-3FBF-F9551375429"+si, m.ID)
		assert.Equal(t, "test_name"+si, m.Name)
		assert.Equal(t, "testemail"+si+"@test.email", m.Email)
		assert.Equal(t, "test,message"+si, m.Text)
		assert.Equal(t, testCreatedTime, m.CreationTime)
	}
}
