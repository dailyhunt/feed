package cmd

type logFormat string

const (
	JSON logFormat = "json"
	TEXT           = "text"
)

type logOutput string

const (
	FILE   logOutput = "file"
	STDOUT           = "stdout"
)

type logConfig struct {
	Level  string
	Format logFormat
	Output logOutput
	File   struct {
		Filename   string
		MaxSize    int
		Compress   bool
		MaxBackups int
		MaxAge     int
	}
}
