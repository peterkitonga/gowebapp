package config

import "github.com/alexedwards/scs/v2"

// AppConfig: holds the application config
type AppConfig struct {
	InProduction bool
	Session      *scs.SessionManager
}
