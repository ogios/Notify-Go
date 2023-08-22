package tcpt_test

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Network string `yaml:"network"`
		Address string `yaml:"address"`
	} `yaml:"server"`
}

var YMLConfig Config = Config{}

func UnmarshalConfig() {
	// 读取文件
	path, _ := os.Getwd()
	ymlb, err := os.ReadFile(path + "/config.yml")
	if err != nil {
		panic(err)
	}

	// 解集
	fmt.Println("Unmarshal...")
	err = yaml.Unmarshal(ymlb, &YMLConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", YMLConfig)
}

func getFieldByTag(s interface{}, tag string) (interface{}, error) {
	v := reflect.ValueOf(s)
	fmt.Println(v)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expect struct")
	}

	typeField := v.Type()

	tags := strings.Split(tag, ".")
	fmt.Println(tags)
	for i := 0; i < typeField.NumField(); i++ {
		f := typeField.Field(i)
		fmt.Println("inner - ", f)

		if f.Tag.Get("yaml") == tags[0] {
			if len(tags) > 1 {
				fmt.Println("next")
				return getFieldByTag(v.Field(i).Interface(), strings.Join(tags[1:], "."))
			} else {
				fmt.Println("gotten")
				return v.Field(i).Interface(), nil
			}
		}
	}
	return nil, fmt.Errorf("field %s not found", strings.Join(tags[0:], "."))
}

func GetConfig(path string) (interface{}, error) {
	return getFieldByTag(YMLConfig, path)
}
