package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var connection *gorm.DB

const engine_sql string = "mysql"

// TODO: Poner en env

const username string = ""
const password string = ""
const database string = ""
const hostname string = ""

func InitializeDatabase() {
	connection = ConnectORM(CreateString())
	log.Println("La conexion a la db fue exitosa")
}

//func CloseConnection() {
//	connection.close
//	log.Println("la conexion con la db fue cerrada")
//}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open(mysql.Open(CreateString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection
}

func CreateString() string {
	return engine_sql + username + ":" + password + "@" + hostname + database
}
