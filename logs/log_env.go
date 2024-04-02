package logs

type LogEnv struct {
	FilePath     string `env:"log_file_path" env_default:"/data/logs"`
	FileName     string `env:"log_file_name"`
	FileSize     int    `env:"log_file_size" env_default:"50"`
	FileMaxCount int    `env:"log_file_max_count" env_default:"32"`
	Compress     bool   `env:"log_compress" env_default:"true"`
	FormatJson   bool   `env:"log_format_json" env_default:"true"`
}
