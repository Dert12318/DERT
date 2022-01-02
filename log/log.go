package log

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type LogMenu struct {
	PostgresConfigLog PostgresConfigLog
	ElasticConfigLog  ElasticConfigLog
	BoolPostgresLog   bool
	BoolElasticLog    bool
}

type PostgresConfigLog struct {
	postgresDB  *gorm.DB
	PostgresLog StandartLog
	NameTable   string
}

type ElasticConfigLog struct {
	elasticDB  *elastic.Client
	ElasticLog StandartLog
	ctx        context.Context
	Indexs     string `validate:"required"`
}

type StandartLog struct {
	typeOfElastic string `validate:"required"`
	level         string `validate:"required"`
	times         string `validate:"required"`
	message       string `validate:"required"`
	location      string `validate:"required"`
	request       string `validate:"required"`
	response      string `validate:"required"`
}

func (e *LogMenu) ConfigLogElastic(ctx context.Context, elasticDB *elastic.Client, Indexs string) *LogMenu {
	e.BoolElasticLog = true
	e.ElasticConfigLog.ctx = ctx
	e.ElasticConfigLog.elasticDB = elasticDB
	e.ElasticConfigLog.Indexs = Indexs
	return e
}

func (e *LogMenu) WriteToLogElastic() (*LogMenu, error) {
	writeIndex, errWriteIndexs := e.ElasticConfigLog.elasticDB.Index().Index(e.ElasticConfigLog.Indexs).BodyJson(e.ElasticConfigLog.ElasticLog).Do(e.ElasticConfigLog.ctx)
	if errWriteIndexs != nil {
		fmt.Println("error :", errWriteIndexs)
		return nil, errWriteIndexs
	}
	fmt.Println("sucess :", writeIndex)
	return e, nil
}

func (e *LogMenu) ConfigLogPostgres(postgresDB *gorm.DB, NameTable string) *LogMenu {
	e.BoolPostgresLog = true
	e.PostgresConfigLog.postgresDB = postgresDB
	e.PostgresConfigLog.NameTable = NameTable
	return e
}

func (e *LogMenu) WriteToLogPostgres() (*LogMenu, error) {
	errCreate := e.PostgresConfigLog.postgresDB.AutoMigrate(StandartLog{})
	if errCreate != nil {
		fmt.Println("error :", errCreate)
		return e, errCreate
	}
	if !e.PostgresConfigLog.postgresDB.Migrator().HasTable(e.PostgresConfigLog.NameTable) {
		err2 := e.PostgresConfigLog.postgresDB.Migrator().RenameTable(e.PostgresConfigLog.PostgresLog, e.PostgresConfigLog.NameTable)
		if err2 != nil {
			fmt.Println("error :", err2)
			return e, err2
		}
	}
	err := e.PostgresConfigLog.postgresDB.Debug().Create(&e.PostgresConfigLog.PostgresLog).Error
	if err != nil {
		fmt.Println("error :", err)
		return e, err
	}
	return e, nil
}

func (e *LogMenu) Errors(index string, message error, level string, location string, request interface{}, response interface{}) (*LogMenu, error) {
	data := StandartLog{}
	data.level = level
	data.location = location
	data.typeOfElastic = "Error"
	data.message = message.Error()
	data.times = time.Now().String()
	strRequest := fmt.Sprintf("%v", request)
	data.request = strRequest
	strResponse := fmt.Sprintf("%v", response)
	data.response = strResponse
	if e.BoolElasticLog {
		e.ElasticConfigLog.ElasticLog = data
		res, err := e.WriteToLogElastic()
		if err != nil {
			return res, err
		}
	}
	if e.BoolPostgresLog {
		e.PostgresConfigLog.PostgresLog = data
		res, err := e.WriteToLogPostgres()
		if err != nil {
			return res, err
		}
	}
	//remove data in struct ElasticLog and PostgresLog
	e.ElasticConfigLog.ElasticLog = StandartLog{}
	e.PostgresConfigLog.PostgresLog = StandartLog{}
	return e, nil
}

func (e *LogMenu) Success(level string, location string, request interface{}, response interface{}) (*LogMenu, error) {
	data := StandartLog{}
	data.level = level
	data.location = location
	data.typeOfElastic = "Success"
	data.message = "Success"
	data.times = time.Now().String()
	strRequest := fmt.Sprintf("%v", request)
	data.request = strRequest
	strResponse := fmt.Sprintf("%v", response)
	data.response = strResponse
	if e.BoolElasticLog {
		e.ElasticConfigLog.ElasticLog = data
		res, err := e.WriteToLogElastic()
		if err != nil {
			return res, err
		}
	}
	if e.BoolPostgresLog {
		e.PostgresConfigLog.PostgresLog = data
		res, err := e.WriteToLogPostgres()
		if err != nil {
			return res, err
		}
	}
	//remove data in struct ElasticLog and PostgresLog
	e.ElasticConfigLog.ElasticLog = StandartLog{}
	e.PostgresConfigLog.PostgresLog = StandartLog{}
	return e, nil
}

func (e *LogMenu) Fatal(level string, message error, location string, request interface{}, response interface{}) (*LogMenu, error) {
	data := StandartLog{}
	data.level = level
	data.location = location
	data.typeOfElastic = "Fatal"
	data.message = message.Error()
	data.times = time.Now().String()
	strRequest := fmt.Sprintf("%v", request)
	data.request = strRequest
	strResponse := fmt.Sprintf("%v", response)
	data.response = strResponse
	if e.BoolElasticLog {
		e.ElasticConfigLog.ElasticLog = data
		res, err := e.WriteToLogElastic()
		if err != nil {
			return res, err
		}
	}
	if e.BoolPostgresLog {
		e.PostgresConfigLog.PostgresLog = data
		res, err := e.WriteToLogPostgres()
		if err != nil {
			return res, err
		}
	}
	//remove data in struct ElasticLog and PostgresLog
	e.ElasticConfigLog.ElasticLog = StandartLog{}
	e.PostgresConfigLog.PostgresLog = StandartLog{}
	panic(message)
}
