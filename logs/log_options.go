package logs

type LogOptions struct {
	FileName     string
	FilePath     string
	FileSize     int
	MaxFileCount int
	Debug        bool
	Compress     bool
	Json         bool
	Color        bool
}
