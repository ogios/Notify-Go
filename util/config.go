package util

import (
	"flag"
	"fmt"
)

type ConnectionOpt struct {
	Network string
	Address string
}

type SocketOpt struct {
	BufferSize int
}

type ServerOpt struct {
	Connection ConnectionOpt
	Socket     SocketOpt
}

type TempfileOpt struct {
	Name string
}

type Config struct {
	Server   ServerOpt
	Tempfile TempfileOpt
	Debug    bool
}

var GlobalConfig Config = Config{
	Debug: false,
	Server: ServerOpt{
		Connection: ConnectionOpt{
			Network: "tcp",
			Address: ":9977",
		},
		Socket: SocketOpt{
			BufferSize: 2048,
		},
	},
	Tempfile: TempfileOpt{
		Name: "android-notify-*",
	},
}

func init() {
	debug := flag.Bool("debug", false, "start with debug logging")
	host := flag.String("h", "", "listen host, default empty string")
	port := flag.String("p", "9977", "listen port, default 9977")
	flag.Parse()
	fmt.Println("Flag parsed: ", *debug, *host, *port)
	GlobalConfig.Server.Connection.Address = (*host + ":" + *port)
	GlobalConfig.Debug = *debug
}

/* func UnmarshalConfig() {
	// read file
	path, _ := os.Getwd()
	ymlb, err := os.ReadFile(path + "/config.yml")
	if err != nil {
		panic(err)
	}

	// Unmarshal
	fmt.Println("Unmarshaling...")
	err = yaml.Unmarshal(ymlb, &YMLConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", YMLConfig)
}

func getFieldByTag(s interface{}, tag string) (interface{}, error) {
	// Get Value under interface
	v := reflect.ValueOf(s)
	// fmt.Println(v)

	// Transform pointer to value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Check if type is struct
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expect struct")
	}

	// Get Value type to match tag
	typeField := v.Type()

	// Split Tags
	tags := strings.Split(tag, ".")
	// fmt.Println(tags)

	// go through fields
	for i := 0; i < typeField.NumField(); i++ {

		// get one field's type
		f := typeField.Field(i)
		// fmt.Println("inner - ", f)

		// match tag
		if f.Tag.Get("yaml") == tags[0] {
			if len(tags) > 1 {
				// match fields under current field
				// fmt.Println("next")
				return getFieldByTag(v.Field(i).Interface(), strings.Join(tags[1:], "."))
			} else {

				// return field if there's no tags left
				fmt.Println("gotten")
				return v.Field(i).Interface(), nil
			}
		}
	}

	// return nil if not match
	return nil, fmt.Errorf("field %s not found", strings.Join(tags[0:], "."))
}

func GetConfig(path string) (interface{}, error) {
	return getFieldByTag(YMLConfig, path)
} */
