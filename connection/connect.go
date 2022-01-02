package connection

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	postgresDB *gorm.DB
	host       string `validate:"required"`
	user       string `validate:"required"`
	pass       string `validate:"required"`
	dbname     string `validate:"required"`
	port       string `validate:"required"`
	sslmode    string `validate:"required"`
	timeZone   string `validate:"required"`
}

type ElasticConfig struct {
	elasticDB *elastic.Client
	host      string `validate:"required"`
	port      string `validate:"required"`
	user      string `validate:"required"`
	password  string `validate:"required"`
	setSniff  bool `validate:"required"`
}

var PostgresLog *gorm.DB
var ElasticLog *elastic.Client

// Func to enter value of Struct

// Postgres Sql
func (s *PostgresConfig) Host(host string) *PostgresConfig {
	s.host = host
	return s
}
func (s *PostgresConfig) User(user string) *PostgresConfig {
	s.user = user
	return s
}
func (s *PostgresConfig) Password(password string) *PostgresConfig {
	s.pass = password
	return s
}
func (s *PostgresConfig) NameDB(dbname string) *PostgresConfig {
	s.dbname = dbname
	return s
}
func (s *PostgresConfig) Port(port string) *PostgresConfig {
	s.port = port
	return s
}
func (s *PostgresConfig) SSLMode(sslmode string) *PostgresConfig {
	s.sslmode = sslmode
	return s
}
func (s *PostgresConfig) Timezone(timeZone string) *PostgresConfig {
	s.timeZone = timeZone
	return s
}

// ElasticSearch
func (s *ElasticConfig) Host(host string) *ElasticConfig {
	s.host = host
	return s
}
func (s *ElasticConfig) Port(port string) *ElasticConfig {
	s.port = port
	return s
}
func (s *ElasticConfig) User(user string) *ElasticConfig {
	s.user = user
	return s
}
func (s *ElasticConfig) Password(password string) *ElasticConfig {
	s.password = password
	return s
}
func (s *ElasticConfig) SetSniff(setSniff bool) *ElasticConfig {
	s.setSniff = setSniff
	return s
}

// Function To connect Log
// Postgre SQL
func (s *PostgresConfig) Connect() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user =%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", s.host, s.user, s.pass, s.dbname, s.port, s.sslmode, s.timeZone)
	postgresDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	s.postgresDB = postgresDB
	PostgresLog = postgresDB
	return postgresDB, err
}

// ElasticSearch
func (s *ElasticConfig) Connect() (*elastic.Client, error) {
	config := fmt.Sprintf("http://%v:%v", s.host, s.port)
	postgresLog, err := elastic.NewClient(elastic.SetURL(config), elastic.SetSniff(s.setSniff), elastic.SetBasicAuth(s.user, s.password))
	s.elasticDB = postgresLog
	ElasticLog = s.elasticDB
	return s.elasticDB, err
}

// func newAsyncHookElasticSearch(Index string, data interface{}) error {
// 	res, err2 := ElasticLog.Index().Index(Index).BodyJson(data)
// }
