package config

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func MarshalFile(filename string, data interface{}) (err error) {
	result, err := Marshal(data)
	if err != nil {
		return
	}
	return ioutil.WriteFile(filename, result, 0666)
}

// Marshal 传入interface{}
func Marshal(data interface{}) (result []byte, err error) {
	// get type and value
	typeInfo := reflect.TypeOf(data)
	valueInfo := reflect.ValueOf(data)
	if typeInfo.Kind() != reflect.Struct {
		return
	}

	var conf []string
	// 可以获取所有属性 获取结构体字段个数：t.NumField()
	for i := 0; i < typeInfo.NumField(); i++ {
		// get field and value
		labelField := typeInfo.Field(i)
		labelValue := valueInfo.Field(i)
		fieldType := labelField.Type
		// 判断类型
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := labelField.Tag.Get("json")
		label := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, label)
		// 拼接 k-v
		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTagVal := keyField.Tag.Get("json")

			valField := labelValue.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTagVal, valField.Interface())
			conf = append(conf, item)
		}
	}
	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return
}

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return UnMarshal(data, result)
}

func UnMarshal(data []byte, result interface{}) (err error) {
	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		return
	}
	if typeInfo.Elem().Kind() != reflect.Struct {
		return
	}
	// 转类型，按行分割
	lineArr := strings.Split(string(data), "\n")
	var myFieldName string

	for _, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			myFieldName, err = myLabel(line, typeInfo.Elem())
			if err != nil {
				return
			}
			continue
		}
		err = myField(myFieldName, line, result)
		if err != nil {
			return
		}
	}
	return
}

func myLabel(line string, typeInfo reflect.Type) (fieldName string, err error) {
	labelName := line[1 : len(line)-1]
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		tagValue := field.Tag.Get("json")
		if labelName == tagValue {
			fieldName = field.Name
			break
		}
	}
	return
}

func myField(fieldName string, line string, result interface{}) (err error) {
	//fmt.Println(line)
	key := strings.TrimSpace(line[0:strings.Index(line, "=")])
	val := strings.TrimSpace(line[strings.Index(line, "=")+1:])
	// 解析到结构体
	resultValue := reflect.ValueOf(result)
	labelValue := resultValue.Elem().FieldByName(fieldName)
	//fmt.Println(labelValue)
	labelType := labelValue.Type()
	var keyName string
	for i := 0; i < labelType.NumField(); i++ {
		// 获取结构体字段
		field := labelType.Field(i)
		tagVal := field.Tag.Get("json")
		if tagVal == key {
			keyName = field.Name
			break
		}
	}
	fmt.Println(keyName)
	// 赋值
	fieldValue := labelValue.FieldByName(keyName)
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fieldValue.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fieldValue.SetUint(i)
	case reflect.Float64:
		f, _ := strconv.ParseFloat(val, 64)
		fieldValue.SetFloat(f)
	}
	return
}
