# Go-DERT-Log
## source package/ requirement
github.com/olivere/elastic/v7 v7.0.30</br>
gorm.io/driver/postgres v1.2.3</br>
gorm.io/gorm v1.22.4
## import package
conn "github.com/Dert12318/Go-DERT-Log/connection" //for connect to DB</br>
log "github.com/Dert12318/Go-DERT-Log/log" //for write Log in DB
## call function
//call the struct for connect to DB</br>
dbConfig := l.PostgresConfig{}</br>
elasticConfig := l.ElasticConfig{}</br>
//try to connect to DB</br>
postgres, err := dbConfig.Host("localhost").NameDB("postgres").Password("Evanroberts14").Port("5432").User("postgres").Timezone("Asia/jakarta").SSLMode("disable").Connect()</br>
if err != nil {</br>
    fmt.Println("error :", err)</br>
    return</br>
}</br>
//example error</br>
_,errConn := elasticConfig.Connect()</br>
// For Elastic</br>
// elastic, err := elasticConfig.Host("localhost").Password("elastic").User("elastic").Port("9200").SetSniff(false).Connect()</br>
Menu.ConfigLogPostgres(db, "NameNewTable")</br>
db, err2 := Menu.Errors( "Message", "Level", "Location", "Request", "Response" )</br>
if err2 != nil {</br>
    fmt.Println("error :", err2)</br>
    return</br>
}else{</br>
    fmt.Println("success log")</br>
}</br>
db, err3 := Menu.Success("Level", "Message", "Location", "Request", "Response" ) </br>
if err2 != nil {</br>
    fmt.Println("error :", err3)</br>
    return</br>
}else{</br>
    fmt.Println("success log")</br>
}</br>
db, err4 := Menu.Fatal("Level", "Message", "Location", "Request", "Response" ) // if success will return panic and log save before panic</br>
if err2 != nil {</br>
    fmt.Println("error :", err4)</br>
    return</br>
}
