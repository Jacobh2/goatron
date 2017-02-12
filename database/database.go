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
    mysqlTable := os.Getenv("MYSQL_TABLE_NAME")
    mysqlAddr := os.Getenv("MYSQL_ADDRESS")

    //open a db connection
    db, err := gorm.Open("mysql", mysqlUser + ":" + mysqlPwd + "@" + mysqlAddr + "/" + mysqlTable + "?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        return nil, err
    }

    return db, nil
}


