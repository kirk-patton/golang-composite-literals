package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghodss/yaml"
)

type Qzu struct {
	Name string
	DOB  string
}

type Baz struct {
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Children []Qzu
}

type Foo struct {
	Bar []Baz
}

func main() {
	f := Foo{
		Bar: []Baz{
			{Name: "Joe", Age: 45, Children: []Qzu{
				{Name: "Bobby", DOB: "01/09/2018"},
			}},
			{Name: "Sue", Age: 20, Children: []Qzu{
				{Name: "Kim", DOB: "05/03/2015"},
			}},
			{Name: "Bob", Age: 24},
		}}

	ymlData, err := yaml.Marshal(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", ymlData)

	u := &Foo{}

	err = yaml.Unmarshal(ymlData, u)
	if err != nil {
		panic(err)
	}

	spew.Dump(u)

}
