package mconfig

type (
	//* example of how you should define your Config Based on BASE config
	// Config struct {
	// 	Base `yaml:",inline" env:",inline"`
	// }
	Base struct {
		Database Database `yaml:"database"`
		Server   Server   `yaml:"server"`
		I18n     I18n     `yaml:"i18n"`
		Logger   Logger   `yaml:"logger"`
		ENV      string   `env:"ENVIRONMENT"`
	}
	Database struct {
		MongoDB MongoDB `yaml:"mongodb"`
	}

	MongoDB struct {
		Url      string `yaml:"mongodb_url" env:"MONGODB_URL"`
		Database string `yaml:"mongodb_db_name" env:"MONGODB_DB_NAME"`
	}

	Server struct {
		Port     int `yaml:"port" env:"PORT"`
		GrpcPort int `yaml:"grpc_port" env:"GRPC_PORT"`
	}

	I18n struct {
		BundlePath string `yaml:"bundle_path" env:"BUNDLE_PATH"`
	}

	Logger struct {
		MaxAge          string `yaml:"max_age" env:"MAX_AGE"`
		MaxSize         string `yaml:"max_size" env:"MAX_SIZE"`
		FilenamePattern string `yaml:"filename_pattern" env:"FILENAME_PATTERN"`
		RotationTime    string `yaml:"rotation_time" env:"ROTATION_TIME"`
		InternalPath    string `yaml:"internal_path" env:"INTERNAL_PATH"`
	}
)
