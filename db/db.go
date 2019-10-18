package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"xorm.io/core"
)


func Init () {

}

func connect () *xorm.Engine {
	dsn := "host=172.31.16.1 port=3308 user=tb_cloud password=123456 dbname=thingsboard sslmode=disable"
	db, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		fmt.Printf("Database Init Error %s", dsn)
		log.Fatal(err)
		return nil
	}

	db.SetSchema("cloud")
	db.ShowSQL()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgresql success")

	db.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "tb_"))

	return db
}