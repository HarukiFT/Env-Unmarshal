package envunmarshal

import (
	"bufio"
	"env-unmarshal/internal/features"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Env struct {
	path string
	data *map[string]interface{}
}

func (e *Env) mapData(file *os.File) *map[string]interface{} {
	s := bufio.NewScanner(file)
	result := make(map[string]interface{})

	for s.Scan() {
		line := s.Text()
		splitted := strings.Split(line, "=")
		if len(splitted) != 2 {
			panic(fmt.Sprintf("Bad line %s", line))
		}

		splitted[0] = strings.Trim(splitted[0], " ")
		splitted[1] = strings.Trim(splitted[1], " ")

		result[splitted[0]] = features.ConvType(splitted[1])
	}

	return &result
}

func (e *Env) ToEnv() {
	for key, value := range *e.data {
		os.Setenv(key, fmt.Sprint(value))
	}
}

func (e *Env) Unmarshal(target interface{}) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.IsNil() {
		panic("Target must be a non-nil pointer")
	}

	targetValue = targetValue.Elem()
	targetType := targetValue.Type()

	for key, value := range *e.data {
		for i := 0; i < targetType.NumField(); i++ {
			field := targetType.Field(i)
			if field.Tag.Get("env") == key {
				fieldValue := targetValue.Field(i)

				if !fieldValue.CanSet() {
					continue
				}

				switch fieldValue.Kind() {
				case reflect.String:
					fieldValue.SetString(value.(string))
				case reflect.Int:
					fieldValue.SetInt(value.(int64))
				default:
					panic(fmt.Sprintf("Unsupported field type: %s", fieldValue.Kind()))
				}
			}
		}
	}
}

func LoadEnv(path string) (*Env, error) {
	env := Env{path: path}

	file, err := features.LoadFileWD(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := env.mapData(file)
	env.data = data

	return &env, nil
}
