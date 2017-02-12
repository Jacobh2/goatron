package database

import(
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() (*gorm.DB, error) {
    //Get the envs for mysql connection
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPwd := os.Getenv("MYSQL_PASSWORD")
    mysqlDb := os.Getenv("MYSQL_DB_NAME")
    mysqlAddr := os.Getenv("MYSQL_ADDRESS")

    if(mysqlAddr == ""){
        mysqlAddr = os.Getenv("mysql") + ":3306"
    }

    println("**** DATABASE CONNECTION ****")
    println("User: " + mysqlUser)
    println("Addr: " + mysqlAddr)


    //open a db connection
    db, err := gorm.Open("mysql", mysqlUser + ":" + mysqlPwd + "@" + mysqlAddr + "/" + mysqlDb + "?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        return nil, err
    }

    return db, nil
}


