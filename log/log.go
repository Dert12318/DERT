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
	PostgresDB  *gorm.DB
	PostgresLog StandartLog
}

type ElasticConfigLog struct {
	ElasticDB  *elastic.Client
	ElasticLog StandartLog
	ctx        context.Context
	Indexs     string
}

type StandartLog struct {
	TypeOfElastic string
	Level         string
	Times         string
	Message       string
	Location      string
	Request       string
	Response      string
}

func (e *LogMenu) ConfigLogElastic(ctx context.Context, ElasticDB *elastic.Client, Indexs string) *LogMenu {
	e.BoolElasticLog = true
	e.ElasticConfigLog.ctx = ctx
	e.ElasticConfigLog.ElasticDB = ElasticDB
	e.ElasticConfigLog.Indexs = Indexs
	return e
}

func (e *LogMenu) WriteToLogElastic() (*LogMenu, error) {
	writeIndex, errWriteIndexs := e.ElasticConfigLog.ElasticDB.Index().Index(e.ElasticConfigLog.Indexs).BodyJson(e.ElasticConfigLog.ElasticLog).Do(e.ElasticConfigLog.ctx)
	if errWriteIndexs != nil {
		fmt.Println("error :", errWriteIndexs)
		return nil, errWriteIndexs
	}
	fmt.Println("sucess :", writeIndex)
	return e, nil
}

func (e *LogMenu) ConfigLogPostgres(PostgresDB *gorm.DB, NameTable string) *LogMenu {
	e.BoolPostgresLog = true
	e.PostgresConfigLog.PostgresDB = PostgresDB
	return e
}

func (e *LogMenu) WriteToLogPostgres() (*LogMenu, error) {
	model := StandartLog{}
	errCreate := e.PostgresConfigLog.PostgresDB.AutoMigrate(model)
	if errCreate != nil {
		fmt.Println("error :", errCreate)
		return e, errCreate
	}
	err := e.PostgresConfigLog.PostgresDB.Debug().Create(&e.PostgresConfigLog.PostgresLog).Error
	if err != nil {
		fmt.Println("error :", err)
		return e, err
	}
	return e, nil
}

func (e *LogMenu) Errors(index string, Message error, Level string, Location string, Request interface{}, Response interface{}) (*LogMenu, error) {
	data := StandartLog{}
	data.Level = Level
	data.Location = Location
	data.TypeOfElastic = "Error"
	data.Message = Message.Error()
	data.Times = time.Now().String()
	strRequest := fmt.Sprintf("%v", Request)
	data.Request = strRequest
	strResponse := fmt.Sprintf("%v", Response)
	data.Response = strResponse
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

func (e *LogMenu) Success(Level string, Location string, Request interface{}, Response interface{}) (*LogMenu, error) {
	data := StandartLog{}
	data.Level = Level
	data.Location = Location
	data.TypeOfElastic = "Success"
	data.Message = "Success"
	data.Times = time.Now().String()
	strRequest := fmt.Sprintf("%v", Request)
	data.Request = strRequest
	strResponse := fmt.Sprintf("%v", Response)
	data.Response = strResponse
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

func (e *LogMenu) Fatal(Level string, Message error, Location string, Request interface{}, Response interface{}) (*LogMenu, error) {
	data := StandartLog{}
	data.Level = Level
	data.Location = Location
	data.TypeOfElastic = "Fatal"
	data.Message = Message.Error()
	data.Times = time.Now().String()
	strRequest := fmt.Sprintf("%v", Request)
	data.Request = strRequest
	strResponse := fmt.Sprintf("%v", Response)
	data.Response = strResponse
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
	panic(Message)
}
