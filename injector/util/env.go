package util

// Env defines for configuration
type Env struct {
	// RjtrackURI is server host and port thats like FQDN
	RjtrackURI string `envconfig:"RJTRACK_URI" default:"0.0.0.0:3000"`
	// DSN is mysql uri
	DSN string `envconfig:"DSN" default:"root:@tcp(127.0.0.1:3306)/injector_development?parseTime=true"`
}
