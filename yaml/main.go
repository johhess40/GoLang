
/*
TODO: Make this functional for use with terraform/pulumi etc....
This code extracts values from YAML files and can be used to perform config tasks
	-Simply provide the path to your yaml file to make it work
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Persons struct {
	People []struct {
		Person struct {
			Name string `yaml:"name"`
			Age  int    `yaml:"age"`
			Job  string `yaml:"job"`
		} `yaml:"person"`
	} `yaml:"people"`
}

func main(){
	c, err := readFile("./test.yml")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%T\n", c.People)
	//fmt.Printf("%v\n", c.People[0])
	//See below for a procedural way to grab all data from a map[string]interface{}...lots of heavy lifting
	//fmt.Printf("%v\n", c["people"].([]interface{})[0].(map[interface{}]interface{})["person"].(map[interface{}]interface{})["name"])
	//fmt.Printf("%v\n", c["people"].([]interface{})[0].(map[interface{}]interface{})["person"].(map[interface{}]interface{})["job"])
	//fmt.Printf("%v\n", c["people"].([]interface{})[0].(map[interface{}]interface{})["person"].(map[interface{}]interface{})["age"])
	//fmt.Printf("%v\n", c["people"].([]interface{})[1].(map[interface{}]interface{})["person"].(map[interface{}]interface{})["name"])
	//fmt.Printf("%v\n", c["people"].([]interface{})[1].(map[interface{}]interface{})["person"].(map[interface{}]interface{})["job"])
	//fmt.Printf("%v\n", c["people"].([]interface{})[1].(map[interface{}]interface{})["person"].(map[interface{}]interface{})["age"])

	//his is how we would loop through our map of interface
	for _, v := range c.People {
		fmt.Printf("Suspect Name: %v\n Suspect Job: %v\n Suspect Age: %v\n",
			v.Person.Name,
			v.Person.Job,
			v.Person.Age,
		)
	}


}

//func readFile reads our YAML file and creates a map of string to interface
func readFile(x string) (*Persons, error){

	file, err := ioutil.ReadFile(x)
	if err != nil {
		log.Fatalf("ioutil.ReadFile failed with error %v", err)
	}


	ppl := Persons{}

	err = yaml.Unmarshal(file, &ppl)
	if err != nil {
		log.Fatalf("Error in file %v", err)
	}


	return &ppl, nil
}