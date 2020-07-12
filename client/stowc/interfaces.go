package stowc

type Driver interface {
	FileReader
	FileWriter
	FileLister
	FileDeleter
}

type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type FileWriter interface {
	WriteFile(path string, data []byte) error
}

type FileLister interface {
	ListFiles() (out chan string, err error)
}

type FileDeleter interface {
	DeleteFile(f string) (err error)
}
