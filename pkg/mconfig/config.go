package mconfig

import "github.com/matiniiuu/mcommon/pkg/paytabs/enums"

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
		Email    Email    `yaml:"email"`
		SMS      SMS      `yaml:"sms"`
		Paytabs  Paytabs  `yaml:"paytabs"`
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
	Email struct {
		AwsSES AwsSES `yaml:"aws_ses"`
	}
	AwsSES struct {
		SESAccessKey  string `yaml:"ses_access_key" env:"SES_ACCESS_KEY"`
		SESSecretKey  string `yaml:"ses_secret_key" env:"SES_SECRET_KEY"`
		SESRegion     string `yaml:"ses_region" env:"SES_REGION"`
		SESSourceMail string `yaml:"ses_source_mail" env:"SES_SOURCE_MAIL"`
	}
	SMS struct {
		Authentica Authentica `yaml:"authentica"`
	}
	Authentica struct {
		AuthorizationKey string `yaml:"authorization_key" env:"AUTHENTICA_AUTHORIZATION_KEY"`
	}

	Paytabs struct {
		ProfileId string              `yaml:"profile_id" env:"PAYTABS_PROFILE_ID"`
		ServerKey string              `yaml:"server_key" env:"PAYTABS_SERVER_KEY"`
		Region    enums.PaytabsRegion `yaml:"region" env:"PAYTABS_REGION"`
	}
)
