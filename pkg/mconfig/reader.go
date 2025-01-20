package mconfig

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func ReadEnv(cfg any) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return readEnvRecursive(reflect.ValueOf(cfg).Elem(), "")
}
func readEnvRecursive(v reflect.Value, parentKey string) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		envKey := field.Tag.Get("env")
		if envKey == "" {
			envKey = strings.ToUpper(field.Name)
			if parentKey != "" {
				envKey = parentKey + "_" + envKey
			}
		}

		switch fieldValue.Kind() {
		case reflect.Struct:
			if err := readEnvRecursive(fieldValue, envKey); err != nil {
				return err
			}
		case reflect.String:
			if envValue := os.Getenv(envKey); envValue != "" {
				fieldValue.SetString(envValue)
			}
		case reflect.Int:
			if envValue := os.Getenv(envKey); envValue != "" {
				if intValue, err := strconv.Atoi(envValue); err == nil {
					fieldValue.SetInt(int64(intValue))
				}
			}
		}
	}
	return nil
}
func ReadYAML(path string, cfg any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); err == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
