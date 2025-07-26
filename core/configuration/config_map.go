package configuration

type Config struct {
	App struct {
		Name string
		Port int
		Env  string
	}
	Database struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	Redis struct {
		Host     string
		Port     int
		DB       int
		Password string
	}
	RabbitMQ struct {
		Host     string
		Port     int
		User     string
		Password string
		VHost    string
	}
}
