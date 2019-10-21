package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"heurd.com/wand-go/wand/config"
	"log"
	"xorm.io/core"
)

var Connection *xorm.Engine

func Init () {
	fmt.Print("abdaf")
	var dsn = config.Get("Db.Dsn").(string)
	connect(dsn)
}

func connect (dsn string) *xorm.Engine {
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