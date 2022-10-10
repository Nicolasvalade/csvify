package csvify

import (
	"fmt"
	"os"
	"reflect"
)

const fieldsep string = ";"
const kvsep string = "="

// CsvifyLine writes a new line in a .csv file with your data.
func CsvifyLine(path string, data any) error {

	// Create it if it doesn't exist, append data otherwise
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = _csvifyLine(file, data)
	if err != nil {
		return err
	}

	// end line
	file.WriteString("\n")
	return nil
}

func _csvifyLine(file *os.File, data any) error {

	val := reflect.ValueOf(data)
	switch val.Kind() {

	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			_csvifyLine(file, val.Index(i).Interface())
		}

	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			_csvifyLine(file, val.Index(i).Interface())
		}

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			_csvifyLine(file, val.Field(i).Interface())
		}

	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			file.WriteString(fmt.Sprintf("%v%s", k, kvsep))
			_csvifyLine(file, v.Interface())
		}

	default:
		_, err := file.WriteString(fmt.Sprintf("%v%s", val, fieldsep))
		if err != nil {
			return err
		}
	}

	return nil
}
