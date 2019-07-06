package main

type Server struct {
	Name        string `yaml:"name"`
	Scheme      string `yaml:"scheme"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Connections int    `yaml:"connections"`
}

func (server Server) Url() string {
	return server.Scheme + "://" + server.Host + ":" + server.Port
}
