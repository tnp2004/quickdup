package configs

type Config struct {
	Server *Server
}

type Server struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Server: &Server{
			Port: "8080",
		},
	}
}
