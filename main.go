package main

import (
	"flag"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"github.com/mft-labs/capturesql/utils"
)

var (
	config *ini.File

)
func main() {
	var conf string
	flag.StringVar(&conf,"conf","app.conf","Config file")
	flag.Parse()
	log.Printf("Running Capture SQL\n")
	util := &utils.Util{}
	err := util.LoadConfig(conf)
	if err!=nil {
		log.Printf("Failed to load config:%v",err)
		os.Exit(1)
	}
	proc := &Process{Util:util}
	proc.RunQueries()
}
