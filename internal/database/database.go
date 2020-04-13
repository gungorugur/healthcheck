package database

//IsHealthy should return true if database is reachable for example mongodb has ping and db.stats() commands.
func IsHealthy() bool {
	return true
}
