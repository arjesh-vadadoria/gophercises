package demo

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func YamlHandler(filePath string) []componants {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	componants := parseUrlYAML(yamlFile)
	fmt.Printf("%v\n", componants)
	return componants
}

func parseUrlYAML(data []byte) []componants {
	var componants []componants
	err := yaml.Unmarshal(data, &componants)
	if err != nil {
		fmt.Printf("I AM DANGER! %v", err.Error())
		println()
		panic(err.Error())
	}
	return componants
}

type componants struct {
	Projects []project `yaml:"projects"`
	Date     string    `yaml:"date"`
}

type project struct {
	Name  string   `yaml:"name"`
	Time  string   `yaml:"time"`
	Tasks []string `yaml:"tasks"`
}
