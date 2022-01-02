package godertlog

import (
	"context"
	"fmt"
	"time"

	"github.com/Dert12318/Go-DERT-Log/vendor/github.com/olivere/elastic/v7"
	"github.com/Dert12318/Go-DERT-Log/vendor/gorm.io/gorm"
)

type PostgresConfigLog struct {
	postgresDB  *gorm.DB
	PostgresLog StandartLog
}

type ElasticConfigLog struct {
	elasticDB     *elastic.Client
	elasticLog    StandartLog
	ctx           context.Context
	indexs        string
	typeOfElastic string
}

type StandartLog struct {
	level    string
	times    string
	message  string
	location string
	request  string
	response string
}

func (e *ElasticConfigLog) WriteToLog() (*ElasticConfigLog, error) {
	writeIndex, errWriteIndexs := e.elasticDB.Index().Index(e.indexs).BodyJson(e.elasticLog).Do(e.ctx)
	if errWriteIndexs != nil {
		fmt.Println("error :", errWriteIndexs)
		return nil, errWriteIndexs
	}
	fmt.Println("sucess :", writeIndex)
	return e, nil
}

func (e *ElasticConfigLog) Errors(message error, location string, request interface{}, response interface{}) (*ElasticConfigLog, error) {
	e.elasticLog.location = location
	e.elasticLog.message = "Error"
	e.elasticLog.message = message.Error()
	timeNow := time.Now()
	e.elasticLog.times = timeNow.GoString()
	strRequest := fmt.Sprintf("%v", request)
	e.elasticLog.request = strRequest
	strResponse := fmt.Sprintf("%v", response)
	e.elasticLog.request = strResponse
	return e, nil
}

func (e *ElasticConfigLog) Success(location string, request interface{}, response interface{}) (*ElasticConfigLog, error) {
	e.elasticLog.location = location
	e.elasticLog.message = "Success"
	timeNow := time.Now()
	e.elasticLog.times = timeNow.GoString()
	strRequest := fmt.Sprintf("%v", request)
	e.elasticLog.request = strRequest
	strResponse := fmt.Sprintf("%v", response)
	e.elasticLog.request = strResponse
	return e, nil
}

func (e *ElasticConfigLog) Fatal(message error, location string, request interface{}, response interface{}) (*ElasticConfigLog, error) {
	e.elasticLog.location = location
	e.elasticLog.message = "Error"
	e.elasticLog.message = message.Error()
	timeNow := time.Now()
	e.elasticLog.times = timeNow.GoString()
	strRequest := fmt.Sprintf("%v", request)
	e.elasticLog.request = strRequest
	strResponse := fmt.Sprintf("%v", response)
	e.elasticLog.request = strResponse
	panic(message)
	return e, nil
}
