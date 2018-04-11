package util

// Env defines for configuration
type Env struct {
	// URI is server host and port thats like FQDN
	URI string `envconfig:"ENV_URI" default:"0.0.0.0:3000"`
	// DSN is mysql uri
	DSN string `envconfig:"ENV_DSN" default:"root:@tcp(127.0.0.1:3306)/injector_development?parseTime=true"`
	// Mode defines project environment model
	Mode string `envconfig:"ENV_MODE" default:"development"`
}

// IsProd returns determined product environment mode
func (e *Env) IsProd() bool {
	return e.Mode == "production"
}
