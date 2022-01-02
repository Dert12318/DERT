package godertlog

import (
	"github.com/Dert12318/Go-DERT-Log/connection"
	"github.com/Dert12318/Go-DERT-Log/log"
)

type Menu struct {
	logging            log.LogMenu
	connectionElastic  connection.ElasticConfig
	connectionPostgres connection.PostgresConfig
}
func(s *Menu) ListMenu () *Menu {
	return s
}