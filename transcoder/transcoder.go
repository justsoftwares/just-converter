package transcoder

import (
	"just-converter/storage"
	"log"
)

func Transcode(data *storage.AppData) {
	for _, filesList := range data.FilesLists {
		log.Println(filesList.FilesList, filesList.Decoding, filesList.Encoding, filesList.StreamSpecifiers)
	}
	log.Println(data.OutputFolder)
}
