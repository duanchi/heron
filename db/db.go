package db

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"heurd.com/wand-go/wand/config"
	"log"
	"net/url"
	"strings"
	"xorm.io/core"
)

var Connection *xorm.Engine

func Init () {
	var err error
	parsedDsn, _ := url.Parse(config.Get("Db.Dsn").(string))
	Connection, err = connect(parsedDsn)

	if err != nil {

	}
}

func connect (dsnUrl *url.URL) (connection *xorm.Engine, err error) {

	defer func() {
		e := recover()
		fmt.Printf("%s", e)
		log.Fatal(err)

		return
	}()

	switch dsnUrl.Scheme {
	case "postgres":
		password, _ := dsnUrl.User.Password()
		tableStack := strings.Split(dsnUrl.Path, ".")
		prefix := dsnUrl.Query().Get("prefix")
		table := ""
		schema := ""
		if len(tableStack) > 1 {
			table = tableStack[1]
			schema = tableStack[0]
		} else {
			table = tableStack[0]
		}

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dsnUrl.Hostname(),
			dsnUrl.Port(),
			dsnUrl.User.Username(),
			password,
			table,
			dsnUrl.Query().Get("sslmode"),
		)

		connection, err = xorm.NewEngine("postgres", dsn)
		if err != nil {
			panic(fmt.Sprintf("Database Init Error %s", dsn))
		}

		if schema != "" {
			connection.SetSchema(schema)
		}

		if prefix != "" {
			connection.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, prefix))
		}

		err = connection.Ping()
		if err != nil {
			log.Fatal(err)
			return
		}

	}

	fmt.Println("connect database success!")
	connection.ShowSQL()

	return

	/*connection, err = xorm.NewEngine("postgres", dsn)
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

	return*/
}