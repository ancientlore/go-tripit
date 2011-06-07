package main

import (
	"template"
	"os"
	"flag"
	"log"
	"fmt"
	"strings"
)

var pkg = flag.String("package", "tripit", "Package for vector class")
var nam = flag.String("name", "", "Name of type held by vector")
var typ = flag.String("type", "", "Actual type held by vector")

func main() {
	var info = make(map[string]string)

	flag.Parse()

	if *pkg == "" {
		log.Print("Package name is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	info["Package"] = *pkg

	if *nam == "" {
		log.Print("Type name is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	info["Name"] = *nam

	if *typ == "" {
		info["Type"] = info["Name"]
	} else {
		info["Type"] = *typ
	}
	if info["Type"][0] == '*' {
		info["IsPtr"] = "yes"
	}

	template := template.New(nil)
	template.SetDelims("{{{", "}}}")
	err := template.Parse(getTemplate())
	if err != nil {
		panic("Cannot parse template")
	}

	filename := fmt.Sprintf("%svector.go", strings.ToLower(info["Name"]))
	file, err := os.Create(filename)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}
	defer file.Close()
	template.Execute(file, info)
}

