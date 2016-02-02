package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var output = "tmp_gen_%d/routes.php"

func TestGenerate(t *testing.T) {
	TestLoadTplFiles(t)
	api, _ := process("raml/valid.raml")
	currentOutput := fmt.Sprintf(output, int32(time.Now().Unix()))
	generate(api, currentOutput)
	_, err := os.Open(currentOutput)
	if err != nil {
		t.Fatalf("Expected output file to exist, got %v\n", err)
	}
	//os.Remove(currentOutput)
}

func TestLoadTplFiles(t *testing.T) {
	//加载文件夹中的templates模板文件
	err := LoadTplFiles("template", ".tp")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Templates Loaded:")
	for p, v := range templates {
		fmt.Println("Template:", p)
		fmt.Println("Path:", v.Path)
	}
}
