package config

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Padliwinata/iam-mfa/bin/config/key"

	"github.com/joho/godotenv"
)

type envConfig struct {
	APMSecretToken     string
	APMUrl             string
	AppEnv             string
	AppName            string
	AppPort            string
	AppVersion         string
	BasicAuthPassword  string
	BasicAuthUsername  string
	CipherKey          string
	ConfigCors         string
	IvKey              string
	KafkaUrl           string
	LogLevel           string
	LogstashHost       string
	LogstashPort       string
	MinioAccessKey     string
	MinioEndpoint      string
	MinioSecretKey     string
	MinioUseSSL        bool
	MongoMasterDBUrl   string
	MongoSlaveDBUrl    string
	PrivateKey         *rsa.PrivateKey
	PublicKey          *rsa.PublicKey
	AccessTokenExpired time.Duration
	ShutdownDelay      int
	SiisKey            string
	SiisUrl            string
	VaccineHost        string
	VaccinePassword    string
	VaccineUsername    string
	TUserHost          string

	RedisHost         string
	RedisPort         string
	RedisPassword     string
	RedisDB           string
	ElasticHost       string
	ElasticUsername   string
	ElasticPassword   string
	ElasticMaxRetries int

	UrlZone              string
	ToogleActiveCheckout string

	StatHour   string
	StatMinute string
	StatSecond string

	TopicConsumerNotifInboxCms string
	TopicSendNotifInboxCms     string
}

func (e envConfig) LogstashPortInt() int {
	i, err := strconv.ParseInt(e.LogstashPort, 10, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func (e envConfig) DnsMariaDB() (string, string) {
	var (
		mariaDbHost     = os.Getenv("MYSQL_HOST")
		mariaDbUsername = os.Getenv("MYSQL_USERNAME")
		mariaDbPassword = os.Getenv("MYSQL_PASSWORD")
		mariaDbName     = os.Getenv("MYSQL_DB_NAME")
	)
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", mariaDbUsername, mariaDbPassword, mariaDbHost, mariaDbName), mariaDbName

}

var envCfg envConfig

func init() {
	err := godotenv.Load()

	if err != nil {
		println(err.Error())
	}

	shutdownDelay, _ := strconv.Atoi(os.Getenv("SHUTDOWN_DELAY"))                // default 0
	minioUseSsl, _ := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))              // default false
	elasticMaxRetries, _ := strconv.Atoi(os.Getenv("ELASTICSEARCH_MAX_RETRIES")) // default false

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	rootApp := strings.TrimSuffix(path, "/bin/config")
	os.Setenv("APP_PATH", rootApp)

	expTime, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRED"))
	if err != nil {
		log.Println(err)
	}

	envCfg = envConfig{
		APMSecretToken:     os.Getenv("ELASTIC_APM_SECRET_TOKEN"),
		APMUrl:             os.Getenv("ELASTIC_APM_SERVER_URL"),
		AppEnv:             os.Getenv("APP_ENV"),
		AppName:            os.Getenv("APP_NAME"),
		AppPort:            os.Getenv("APP_PORT"),
		AppVersion:         os.Getenv("APP_VERSION"),
		BasicAuthPassword:  os.Getenv("BASIC_AUTH_PASSWORD"),
		BasicAuthUsername:  os.Getenv("BASIC_AUTH_USERNAME"),
		CipherKey:          os.Getenv("AES_KEY"),
		ConfigCors:         os.Getenv("CORS_CONFIG"),
		IvKey:              "",
		KafkaUrl:           os.Getenv("KAFKA_HOST_URL"),
		LogLevel:           os.Getenv("LOG_LEVEL"),
		LogstashHost:       os.Getenv("LOGSTASH_HOST"),
		LogstashPort:       os.Getenv("LOGSTASH_PORT"),
		MinioAccessKey:     os.Getenv("MINIO_ACCESS_KEY"),
		MinioEndpoint:      os.Getenv("MINIO_END_POINT"),
		MinioSecretKey:     os.Getenv("MINIO_SECRET_KEY"),
		MinioUseSSL:        minioUseSsl,
		MongoMasterDBUrl:   os.Getenv("MONGO_MASTER_DATABASE_URL"),
		MongoSlaveDBUrl:    os.Getenv("MONGO_SLAVE_DATABASE_URL"),
		PrivateKey:         key.LoadPrivateKey(),
		PublicKey:          key.LoadPublicKey(),
		AccessTokenExpired: expTime,
		ShutdownDelay:      shutdownDelay,
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		RedisDB:            os.Getenv("REDIS_DB"),

		ElasticHost:       os.Getenv("ELASTICSEARCH_HOST"),
		ElasticUsername:   os.Getenv("ELASTICSEARCH_USERNAME"),
		ElasticPassword:   os.Getenv("ELASTICSEARCH_PASSWORD"),
		ElasticMaxRetries: elasticMaxRetries,

		UrlZone:              os.Getenv("URL_ZONE"),
		ToogleActiveCheckout: os.Getenv("TOOGLE_ACTIVE_CHECKOUT"),

		StatHour:   os.Getenv("STAT_HOUR_SCHEDULER"),
		StatMinute: os.Getenv("STAT_MINUTE_SCHEDULER"),
		StatSecond: os.Getenv("STAT_SECOND_SCHEDULER"),

		TopicConsumerNotifInboxCms: os.Getenv("TOPIC_CONSUMER_NOTIF_INBOX_CMS"),
		TopicSendNotifInboxCms:     os.Getenv("TOPIC_SEND_NOTIF_INBOX_CMS"),
	}
}

func GetConfig() *envConfig {
	return &envCfg
}
