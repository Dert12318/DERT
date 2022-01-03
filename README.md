# Go-DERT-Log
## source package/ requirement
github.com/olivere/elastic/v7 v7.0.30**Enter**
gorm.io/driver/postgres v1.2.3
gorm.io/gorm v1.22.4
## import package
conn "github.com/Dert12318/Go-DERT-Log/connection" //for connect to DB
log "github.com/Dert12318/Go-DERT-Log/log" //for write Log in DB
## call function
//call the struct for connect to DB
dbConfig := l.PostgresConfig{}
elasticConfig := l.ElasticConfig{}
//try to connect to DB
postgres, err := dbConfig.Host("localhost").NameDB("postgres").Password("Evanroberts14").Port("5432").User("postgres").Timezone("Asia/jakarta").SSLMode("disable").Connect()
if err != nil {
    fmt.Println("error :", err)
    return
}
//example error
_,errConn := elasticConfig.Connect()
// For Elastic
// elastic, err := elasticConfig.Host("localhost").Password("elastic").User("elastic").Port("9200").SetSniff(false).Connect()
Menu.ConfigLogPostgres(db, "lalala")
db, err2 := Menu.Errors( "Message", "Level", "Location", "Request", "Response" )
if err2 != nil {
    fmt.Println("error :", err2)
    return
}else{
    fmt.Println("success log")
}
db, err3 := Menu.Success("Level", "Message", "Location", "Request", "Response" ) 
if err2 != nil {
    fmt.Println("error :", err3)
    return
}else{
    fmt.Println("success log")
}
db, err4 := Menu.Fatal("Level", "Message", "Location", "Request", "Response" ) // if success will return panic and log save before panic
if err2 != nil {
    fmt.Println("error :", err4)
    return
}
