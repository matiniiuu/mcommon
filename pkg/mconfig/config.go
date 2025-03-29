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
		Uploader Uploader `yaml:"uploader"`
		Queue    Queue    `yaml:"queue"`
		ENV      string   `env:"ENVIRONMENT"`
	}

	Uploader struct {
		AwsS3      AwsS3      `yaml:"aws_s3"`
		Cloudinary Cloudinary `yaml:"cloudinary"`
	}
	AwsS3 struct {
		S3AccessKey string `yaml:"s3_access_key" env:"S3_ACCESS_KEY"`
		S3SecretKey string `yaml:"s3_secret_key" env:"S3_SECRET_KEY"`
		S3Region    string `yaml:"s3_region" env:"S3_REGION"`
		S3Bucket    string `yaml:"s3_bucket" env:"S3_BUCKET"`
		S3BaseUrl   string `yaml:"s3_base_url" env:"S3_BASE_URL"`
	}
	Cloudinary struct {
		Url string `yaml:"cloudinary_url" env:"CLOUDINARY_URL"`
	}

	Database struct {
		MongoDB    MongoDB    `yaml:"mongodb"`
		PostgreSQL PostgreSQL `yaml:"postgresql"`
	}

	MongoDB struct {
		Url      string `yaml:"mongodb_url" env:"MONGODB_URL"`
		Database string `yaml:"mongodb_db_name" env:"MONGODB_DB_NAME"`
	}
	PostgreSQL struct {
		Url string `yaml:"postgresql_url" env:"POSTGRESQL_URL"`
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
	Queue struct {
		RabbitMq RabbitMq `yaml:"rabbitmq"`
	}
	RabbitMq struct {
		Url string `yaml:"rabbitmq_url" env:"RABBITMQ_URL"`
	}
)
