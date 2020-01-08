package imporer

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/vkd/interview-back-message-board/model"
)

// MessageStorager - storage of the imported messages.
type MessageStorager interface {
	ImportMessage(model.Message) error
}

// ImportMessages - import messages to the storage from csv file.
func ImportMessages(source io.Reader, storage MessageStorager) error {
	r := csv.NewReader(source)
	r.ReuseRecord = true

	// skip first line with titles
	_, err := r.Read()
	if err != nil {
		return fmt.Errorf("error on read first line on csv file: %w", err)
	}

	for {
		line, err := r.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("error on read csv: %w", err)
		}

		if len(line) != 5 {
			return fmt.Errorf("wrong csv format: it should be 5 fields (%s)", strings.Join(line, ","))
		}

		m := model.Message{
			ID:    line[0],
			Name:  line[1],
			Email: line[2],
			Text:  line[3],
		}
		err = m.SetCreationTime(line[4])
		if err != nil {
			return fmt.Errorf("wrong time format on message (%s): %w", strings.Join(line, ","), err)
		}

		err = storage.ImportMessage(m)
		if err != nil {
			return fmt.Errorf("error on import a message to the storage: %w", err)
		}
	}
}
