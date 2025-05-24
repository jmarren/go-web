package config

type config struct {
	Host string
	Port string
}

var Port = ":8080"

var Config *config = &config{
	Host: "",
	Port: ":8080",
}
