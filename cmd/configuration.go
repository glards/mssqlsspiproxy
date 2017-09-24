package main

// Configuration holds the configuration options
type Configuration struct {
	ListenPort uint16

	ServerDial string
}

// NewConfiguration creates the default configuration
func NewConfiguration() *Configuration {
	return &Configuration{
		ListenPort: 1432,
		ServerDial: "127.0.0.1:1433",
	}
}
