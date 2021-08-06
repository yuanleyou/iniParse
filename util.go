package main

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var logger *log.Logger

func init(){
	logger = log.New(os.Stdout, "[reflect] ", log.LstdFlags|log.Lshortfile)
}

func setField(lableName string, line string, v interface{}) error {
	// 分离line中的key 和 value
	key := strings.TrimSpace(line[0:strings.Index(line, "=")])
	value := strings.TrimSpace(line[strings.Index(line, "=") + 1:])
	// 设置到结构体字段中
	valueInfo := reflect.ValueOf(v)
	lableValue := valueInfo.Elem().FieldByName(lableName)
	lableType := lableValue.Type()
	var keyName string
	for i := 0; i < lableType.NumField(); i++ {
		field := lableType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyName = field.Name
			break
		}
	}
	// 给结构体字段赋值
	fieldValue := lableValue.FieldByName(keyName)
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int:
		parseInt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Println(err)
			return err
		}
		fieldValue.SetInt(parseInt)
	case reflect.Uint:
		parseUint, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			logger.Println(err)
			return err
		}
		fieldValue.SetUint(parseUint)
	case reflect.Float32:
		parseFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			logger.Println(err)
			return err
		}
		fieldValue.SetFloat(parseFloat)
	default:
		logger.Println("类型不支持设置到结构体")
	}
	return nil
}

func getFieldName(data string, typeInfo reflect.Type) (string, error) {
	// 去掉字符串两端的方括号
	section := strings.TrimSuffix(strings.TrimPrefix(data, "["), "]")
	//logger.Println(section)
	// 返回结构体tag匹配section的字段名
	for i := 0; i < typeInfo.NumField(); i++ {
		field := typeInfo.Field(i)
		tagVal := field.Tag.Get("ini")
		//logger.Println(field.Name)
		if tagVal == section {
			return field.Name, nil
		}
	}
	return "", nil
}