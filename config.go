package ottolib

var Config = struct {
	Env     string
	Version string
	Host    string
	Port    string
	DB      struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		Schema   string
		SslMode  string
	}
	Kafka struct {
		Host string // include port if necessary and separate by comma if cluster
	}
	Swagger struct {
		Host     string
		Port     string
		User     string
		Password string
	}
	APIM struct {
		Host string
		Port string
	}
}{}
