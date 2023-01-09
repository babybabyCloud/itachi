package configuration

type bodyType string

const (
	JSON       bodyType = "JSON"
	TEXT       bodyType = "TEXT"
	RAW        bodyType = "RAW"
	BINARY     bodyType = "BINARY"
	FORM       bodyType = "FORM"
	URLENCODED bodyType = "URLENCODED"
	XML        bodyType = "XML"
)

// RequestConf is used to configure the request
type RequestConf struct {

	// Reference is the key "reference"
	Reference string `yaml:"reference"`

	// Name is the Requeset name, this can be used in reference
	Name string `yaml:"name"`

	// Method specify HTTP method
	Method string `yaml:"method"`

	// Path is the path of this request
	Path string `yaml:"path"`

	// Parameters is the request parameters
	Parameters Parameter `yaml:"parameters"`

	// Body is the request body
	Body Body `yaml:"body"`

	// Headers is the headers of the request
	Headers Header `yaml:"headers"`
}

// RequestConfs is the alias of map of RequestConf
type RequestConfs map[string]RequestConf

func (r RequestConfs) merge(other RequestConfs) {
	for key, value := range other {
		r[key] = value
	}
}

// Body is the rueset body
type Body struct {

	// Type specifies the type of body
	Type bodyType `yaml:"type"`

	// JsonValue is the json value of this body
	JsonValue string `yaml:"jsonValue"`

	// TextValue is the text value of this body
	TextValue string `yaml:"textValue"`

	// RawValue is the raw valule of this body
	RawValue string `yaml:"rawValue"`

	// BinaryValue is the binary value of this body
	BinaryValue string `yaml:"binaryValue"`

	// FormValue is the multipart/form-data of this body
	FormValue KeyPairValue `yaml:"formValue"`

	// XMLValue is the XML value of this body
	XMLValue string `yaml:"xmlValue"`

	// URLEncodedValue is the application/x-www-form-urlencoded of this body
	URLEncodedValue KeyPairValue `yaml:"urlencodedValue"`
}

// KeyPairValue is the alias of map[string]interface{}
type KeyPairValue map[string]any
