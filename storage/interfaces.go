package storage

type Driver struct {
	DirCreator
	FileReader
	FileWriter
	FileLister
}

type DirCreator interface {
	CreateDir(path string) error
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
