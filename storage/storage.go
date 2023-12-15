package storage

type FilesList struct {
	FilesList        []string
	StreamSpecifiers []string
	Decoding         string
	Encoding         string
}

type AppData struct {
	FilesLists   []FilesList
	OutputFolder string
}

func NewAppData() *AppData {
	return &AppData{}
}
