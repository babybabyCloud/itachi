package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (

	// FILE is the constant of include configuration with type "file"
	FILE = "file"

	// HTTP is the constant of include configuration with type "http"
	HTTP = "http"
)

// IncludeConf is the model of includes in a configuration file
type IncludeConf struct {

	// Type is the key "type"
	Type string `yaml: "type"`

	// Path is the key "path"
	Path string `yaml: "path"`
}

// Includer
type Includer interface {
	Read() []byte
}

// FileInclude
type FileInclude struct {
	IncludeConf
}

// Read reads a file configuration from local directory
func (inc *FileInclude) Read() []byte {
	file, err := os.Open(inc.Path)
	if err != nil {
		// TODO repalce panic with self-defined error message and exit
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		// TODO repalce panic with self-defined error message and exit
		panic(err)
	}

	return data
}

// HTTPInclude
type HTTPInclude struct {
	IncludeConf
}

func (inc *HTTPInclude) Read() []byte {
	// TODO
	return []byte{}
}

// NewInclude creates a new Includer according to IncludeConf
func NewInclude(conf IncludeConf) Includer {
	switch conf.Type {
	case FILE:
		return &FileInclude{
			IncludeConf: conf,
		}
	case HTTP:
		return &HTTPInclude{
			IncludeConf: conf,
		}
	default:
		// TODO repalce panic with self-defined error message and exit
		panic(fmt.Sprintf("Unrecognized type of includes, possible values are '%s', '%s'", FILE, HTTP))
	}
}
