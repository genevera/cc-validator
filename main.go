package main

import (
	"flag"
	"github.com/ztsyed/cc-validator/validate"
	"io/ioutil"
	"log"
)

func main() {
	var ip = flag.String("config", "", "file path to validate")
	flag.Parse()
	data, err := ioutil.ReadFile(*ip)
	if err != nil {
		log.Fatalln("Error reading config file", err.Error())
	}
	log.Println("Going to validate ", *ip)
	report, err := validate.Validate(data)
	if err != nil {
		log.Fatalln("Error validating config file", err.Error())
	}
	if len(report.Entries()) > 0 {
		log.Fatalln("Cloud-Config Invalid: ", report.Entries())
	} else {
		log.Println("Cloud-Config valid:", *ip)
	}
}
