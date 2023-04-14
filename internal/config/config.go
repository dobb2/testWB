package config

type Config struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	User     string `env:"POSTGRES_USER" envDefault:"dobb2"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"root"`
	Port     string `env:"POSTGRES_PORT" envDefault:"54320"`
	Db       string `env:"POSTGRES_DB" envDefault:"testWB"`
	Address  string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`

	NatsClusterID string `env:"NATS_CLUSTER" envDefault:"test-cluster-wb"`
	NatsClientID  string `env:"NATS_CLIENT" envDefault:"WB-service"`
	NatsDurableID string `env:"NATS_DURABLE" envDefault:"order-service-durable"`
}
