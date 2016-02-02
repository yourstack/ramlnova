package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/buddhamagnet/raml"
)

var (
	ramlFile  string
	genFile   string
	templates = make(map[string]Template)
)

func init() {
	flag.StringVar(&ramlFile, "ramlfile", "api.raml", "RAML file to parse")
	flag.StringVar(&genFile, "genfile", "tmp/routes.php", "Filename to use for output")
}

// Process processes a RAML file and returns an API definition.
func process(file string) (*raml.APIDefinition, error) {
	routes, err := raml.ParseFile(file)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing RAML file: %s\n", err.Error())
	}
	return routes, nil
}

func main() {
	//加载文件夹中的templates模板文件
	err := LoadTplFiles("template", ".tp")
	if err != nil {
		log.Fatal(err)
	}

	//解析命令行参数
	flag.Parse()

	api, err := process(ramlFile)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Processing API spec for", ramlFile)
	generate(api, genFile)
	log.Println("Created Laravel Project in ", genFile)
}

// Generate handler functions based on an API definition.
func generate(api *raml.APIDefinition, genFile string) {

	os.Remove(genFile)
	err := os.MkdirAll(genFile, 0755)
	if err != nil {
		log.Fatal(err)
	}
	os.Remove(genFile)
	f, err := os.Create(genFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//######### routes.php
	// Write the routes.php header.
	f.WriteString(templates["routeHead"].Content)
	// Start the route method.
	r := template.Must(template.New("routeEntry").Parse(templates["routeEntry"].Content))
	for uri, resource := range api.Resources {
		generateRoute("", uri, &resource, r, f)
	}
	//######### end of routes.php

	//######### Controllers/xxx/xxx.php
	// Now add the HTTP Contrllers.
	//t := template.Must(template.New("controllerText").Parse(templates["controllerText"].Content))
	//for uri, resource := range api.Resources {
	//	generateController("", uri, &resource, t, f)
	//}
	//######### end of Controllers/xxx/xxx.php
	format(f)
}

// format runs go fmt on a file.
func format(f *os.File) {
	// Run go fmt on the file.
	cmd := exec.Command("go", "fmt")
	cmd.Stdin = f
	_ = cmd.Run()
}

// generateRoute builds app/Http/routes.php
func generateRoute(parent, uri string, resource *raml.Resource, e *template.Template, f *os.File) {
	path := parent + uri
	err := e.Execute(f, Route{path, resource.UriParameters, resource.Methods(), resource.Description, resource.DisplayName})
	if err != nil {
		log.Println("executing template:", err)
	}

	// Get all children.
	for nestname, nested := range resource.Nested {
		generateRoute(path, nestname, nested, e, f)
	}
}

// generateController creates many Controller struct at app/Http/Controllers/xxx/xxx.php
func generateController(parent, name string, resource *raml.Resource, t *template.Template, f *os.File) string {
	path := parent + name

	for _, method := range resource.Methods() {
		err := t.Execute(f, ControllerInfo{method.DisplayName, method.Name, path, method.Description})
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	// Get all children.
	for nestname, nested := range resource.Nested {
		return generateController(path, nestname, nested, t, f)
	}
	return path
}
