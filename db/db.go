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
	var dsn = config.Get("Db.Dsn").(string)
	var err error
	Connection, err = connect(dsn)

	if err != nil {

	}
}

func connect (dsn string) (connection *xorm.Engine, err error) {
	connection, err = xorm.NewEngine("postgres", dsn)
	if err != nil {
		fmt.Printf("Database Init Error %s", dsn)
		log.Fatal(err)
		return
	}

	connection.SetSchema("cloud")
	connection.ShowSQL()

	err = connection.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("connect postgresql success")

	connection.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "tb_"))

	return
}