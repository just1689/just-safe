package model

type Driver interface {
	FileReader
	FileWriter
	FileLister
}

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type FileWriter interface {
	WriteFile(path string, data []byte) error
}

type FileLister interface {
	ListFiles(path string) (out chan string, err error)
}
