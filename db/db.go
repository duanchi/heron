package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"go.heurd.com/heron-go/heron/config"
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
		if e != nil {
			fmt.Printf("%s", e)
			log.Fatal(err)
		}
		return
	}()

	switch dsnUrl.Scheme {
	case "postgres":
		password, _ := dsnUrl.User.Password()
		dbStack := strings.Split(strings.Trim(dsnUrl.Path, "/"), "/")
		prefix := dsnUrl.Query().Get("prefix")
		dbname := ""
		schema := ""
		if len(dbStack) > 1 {
			dbname = dbStack[0]
			schema = dbStack[1]
		} else {
			dbname = dbStack[0]
		}

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dsnUrl.Hostname(),
			dsnUrl.Port(),
			dsnUrl.User.Username(),
			password,
			dbname,
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

	case "mysql":

		host := dsnUrl.Host
		prefix := dsnUrl.Query().Get("prefix")
		dsnUrl.Query().Del("prefix")
		query := dsnUrl.Query().Encode()

		if query == "" {
			query = ""
		} else {
			query = "?" + query
		}

		if host[0:1] == "/" {
			host = "unix(" + host + ")"
		} else {
			host = "tcp(" + host + ")"
		}

		dsn := dsnUrl.User.String() + "@" + host + dsnUrl.Path + query

		connection, err = xorm.NewEngine("mysql", dsn)
		if err != nil {
			panic(fmt.Sprintf("Database Init Error %s", dsn))
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
}