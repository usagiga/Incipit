package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Load from environment variables through reading struct key.
// It compatible with Int, String, Bool only
func Load(out interface{}) (err error) {
	t := reflect.TypeOf(out).Elem()
	v := reflect.ValueOf(out).Elem()
	fields := t.NumField()

	for i := 0; i < fields; i++ {
		fType := t.Field(i)
		fTypeKind := fType.Type.Kind()
		fVal := v.Field(i)

		// Look up `env` struct tag
		envKey, ok := fType.Tag.Lookup("env")
		if !ok || envKey == "" {
			continue
		}

		// Look up env
		envStrVal, ok := os.LookupEnv(envKey)
		if !ok || envKey == "" {
			errMsg := fmt.Sprint("[Config] environment variable is not set or something went wrong.")
			return errors.New(errMsg)
		}

		// Set the value
		switch fTypeKind {
		case reflect.String:
			envRefVal := reflect.ValueOf(envStrVal)
			fVal.Set(envRefVal)
		case reflect.Int:
			// Cast
			envIntVal, err := strconv.Atoi(envStrVal)
			if err != nil {
				return err
			}
			envRefVal := reflect.ValueOf(envIntVal)
			fVal.Set(envRefVal)
		case reflect.Bool:
			envBoolVal := envStrVal == "true"
			envRefVal := reflect.ValueOf(envBoolVal)
			fVal.Set(envRefVal)
		default:
			errMsg := fmt.Sprint("[Config] Not compatible with this type", fTypeKind)
			return errors.New(errMsg)
		}
	}

	return nil
}
