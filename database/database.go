package database

import(
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() (*gorm.DB, error) {
    //open a db connection
    db, err := gorm.Open("mysql", "todo:1234@/todo?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        return nil, err
    }

    return db, nil
}


