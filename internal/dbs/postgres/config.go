package postgres

type Config struct {
	Database string `config:"POSTGRES_DB" yaml:"database"`
	User     string `config:"POSTGRES_USER" yaml:"user"`
	Password string `config:"POSTGRES_PASSWORD" yaml:"password"`
	Host     string `config:"POSTGRES_HOST" yaml:"host"`
	Port     string `config:"POSTGRES_PORT" yaml:"port"`
	Ssl      string `config:"POSTGRES_SSL" yaml:"ssl"`
}

func (c Config) String() string {
	return "host=" + c.Host + " port=" + c.Port + " user=" + c.User + " password=" + c.Password + " dbname=" + c.Database + " sslmode=" + c.Ssl
}
