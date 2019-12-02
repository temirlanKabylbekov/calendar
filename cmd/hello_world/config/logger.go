package config

type LoggerConfig struct {
	FilePath string `mapstructure:"file_path"`
	Level    string
	Encoding string
	Message  struct {
		CallerName  string `mapstructure:"caller_name"`
		TimeName    string `mapstructure:"time_name"`
		MessageName string `mapstructure:"message_name"`
		LevelName   string `mapstructure:"level_name"`
	}
}
