package config

import "github.com/gin-contrib/cors"

func ConfigureCors() (cors.Config, error) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Adjust origins as needed
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	return config, nil
}
