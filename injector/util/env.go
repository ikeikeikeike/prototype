package util

// Env defines for configuration
type Env struct {
	// RjtrackURI is server host and port thats like FQDN
	RjtrackURI string `envconfig:"RJTRACK_URI" default:"0.0.0.0:3000"`
}
