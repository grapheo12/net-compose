package parser

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parse the Yaml file given by fname
func Parse(fname string) map[interface{}]interface{} {
	config := make(map[interface{}]interface{})
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		panic(err)
	}

	return config
}
