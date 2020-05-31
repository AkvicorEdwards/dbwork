package main

import (
	"dbwork/config"
	"dbwork/da"
	"dbwork/handler"
	"dbwork/tpl"
	"fmt"
	"github.com/AkvicorEdwards/argsparser"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	parseArgs()
	tpl.Parse()
	handler.ParsePrefix()
	server := http.Server{
		Addr:              config.Addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Minute,
	}
	log.Printf("http://127.0.0.1%s", config.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func parseArgs() {
	argsparser.Version = `v1.0`
	argsparser.Help = `
Usage: dbwork [option...] [arguments...]

Example: 

	fi port 8080

The commands are:

	init_database  [username] [user password] [database name]
                                     Create Database And Table
	port           [port]            Set http listening port
	dbname         [database name]   Set database name
	dbuser         [username]        Set database user
	dbpassword     [user password]   Set user password

Default:
	port           8080
	dbname         dbwork
	dbuser         root
	dbpassword     password
`

	argsparser.Add("help", 0, func(str []string) {
		fmt.Println(argsparser.Help)
		os.Exit(0)
	})
	argsparser.Add("version", 0, func(str []string) {
		fmt.Println(argsparser.Version)
		os.Exit(0)
	})
	argsparser.Add("init_database", 3, func(str []string) {
		da.Init(str[1], str[2], str[3])
		os.Exit(0)
	})
	argsparser.Add("port", 1, func(str []string) {
		config.Addr = fmt.Sprintf(":%s", str[1])
	})
	argsparser.Add("dbname", 1, func(str []string) {
		config.DBName = fmt.Sprintf("%s", str[1])
	})
	argsparser.Add("dbuser", 1, func(str []string) {
		config.DBUser = fmt.Sprintf("%s", str[1])
	})
	argsparser.Add("dbpassword", 1, func(str []string) {
		config.DBPassword = fmt.Sprintf("%s", str[1])
	})

	argsparser.Parse()
}
