package config

type DBConfig struct {
	Host                string `json:"host" validate:"required"`
	Port                int    `json:"port" validate:"required"`
	Name                string `json:"name" validate:"required"`
	User                string `json:"user" validate:"required"`
	Pass                string `json:"pass" validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

type DB struct {
	Read  string `json:"read" validate:"required"`
	Write string `json:"write" validate:"required"`
}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type Config struct {
	Mode              Mode   `json:"mode"                       validate:"required"`
	ServiceName       string `json:"service_name"               validate:"required"`
	HttpPort          int    `json:"http_port"                  validate:"required"`
	JwtSecret         string `json:"jwt_secret"                 validate:"required"`
	DB                DB     `json:"db"                         validate:"required"`
	MigrationSource   string `json:"migrations"                 validate:"required"`
	RmqQueuePrefix    string `json:"rmq_queue_prefix"           validate:"required"`
	RabbitmqURL       string `json:"rmq_url"                    validate:"required"`
	RmqReconnectDelay int    `json:"rmq_reconnect_delay"        validate:"required"`
	RmqRetryInterval  int    `json:"rmq_retry_interval"         validate:"required"`
	//GrpcUrls          GrpcUrlsConfig `json:"grpc_urls"                  validate:"required"`
	HealthCheckRoute string `json:"health_check_route"`
	DownloadLimit    int    `json:"download_limit"              validate:"required"`
}

var config *Config

func init() {
	config = &Config{}
}

func GetConfig() Config {
	return *config
}
