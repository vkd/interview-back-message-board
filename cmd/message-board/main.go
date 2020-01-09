package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/vkd/interview-back-message-board/importer"
	"github.com/vkd/interview-back-message-board/server"
	"github.com/vkd/interview-back-message-board/storage"
)

func main() {
	port := flag.String("port", ":8000", "Port of service")
	importCsv := flag.String("import-file", "", "Import csv file")
	flag.Parse()

	var storage storage.SliceMessages

	if importCsv != nil && *importCsv != "" {
		log.Printf("Start importing messages from %s file", *importCsv)

		file, err := os.Open(*importCsv)
		if err != nil {
			log.Fatalf("Error on open csv file: %v", err)
		}
		err = imporer.ImportMessages(file, &storage)
		if err != nil {
			log.Fatalf("Error on import csv file: %v", err)
		}
		err = file.Close()
		if err != nil {
			log.Fatalf("Error on close file: %v", err)
		}

		log.Printf("Messages are imported")
	}

	srv := server.New(&storage)
	err := srv.Run(*port)
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf("Error on run server: %v", err)
		}
	}
}
