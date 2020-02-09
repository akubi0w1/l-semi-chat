package conf

// Server config
const (
	addr = "localhost"
	host = "localhost"
	port = "8080"
)

// Database config
const (
	dbEngine   = "mysql"
	dbHost     = "localhost"
	dbPort     = "3316"
	dbUser     = "root"
	dbPassword = "password-lsemi"
	dbName     = "ls_chat"
)

// LoadServerConfig setting server config
func LoadServerConfig() map[string]string {
	serv := map[string]string{
		"addr": addr,
		"host": host,
		"port": port,
	}
	return serv
}

// LoadDBConfig setting database config
func LoadDBConfig() map[string]string {
	db := map[string]string{
		"engine":   dbEngine,
		"host":     dbHost,
		"port":     dbPort,
		"user":     dbUser,
		"password": dbPassword,
		"db":       dbName,
	}
	return db
}
