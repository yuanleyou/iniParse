package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// struct -> []byte
func INIMarshal(v interface{}) ([]byte, error) {
	// 获取反射类型
	typeInfo := reflect.TypeOf(v)
	valueInfo := reflect.ValueOf(v)
	// 判断是否是结构体
	if typeInfo.Kind() != reflect.Struct {
		return []byte{}, errors.New("the 🔠input type should struct")
	}
	var conf []string
	// 获取结构体字段并逐个处理
	for i := 0; i < typeInfo.NumField(); i++ {
		//logger.Println(typeInfo.Field(i).Type.Kind())
		//logger.Println(valueInfo.Field(i))
		labelField := typeInfo.Field(i)
		labelVal := valueInfo.Field(i)
		if labelField.Type.Kind() != reflect.Struct {
			continue
		}
		// 制作[SERVER]和[CLIENT]
		// 从结构体的tag中获取别名
		tagVal := labelField.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = labelField.Name // 如果没有别名则将字段名赋值给tag
		}
		// 拼接配置文件的section字符串
		section := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, section)
		//logger.Println(conf)
		// 拼接配置文件的字段值 k = v
		for j := 0; j < labelField.Type.NumField(); j++ {
			//logger.Println(labelField.Type.Field(j).Name)
			// 取key
			keyField := labelField.Type.Field(j).Tag.Get("ini")
			if len(keyField) == 0 {
				keyField = labelField.Type.Field(j).Name
			}
			// 取value
			valueField := labelVal.Field(j) //这是reflect类型的值
			//logger.Println(valueField.Interface()) //这是interface类型的值
			item := fmt.Sprintf("%s = %v\n", keyField, valueField.Interface())
			//logger.Printf(item)
			conf = append(conf, item)
			//logger.Println(conf)
		}
	}
	var result []byte
	//将字符串转换为切片
	for _, val := range conf {
		byteVal := []byte(val)
		result = append(result, byteVal...)
	}
	return result, nil
}

// []byte -> struct
func INIUnMarshal(data []byte, v interface{}) error {
	// 判断v是不是指针
	typeInfo := reflect.TypeOf(v)
	if typeInfo.Kind() != reflect.Ptr {
		return nil
	}
	// 判断是不是结构体
	if typeInfo.Elem().Kind() != reflect.Struct {
		return nil
	}
	// 定义全局变量fieldName
	var fieldName string

	// 按行切割
	lineSlice := strings.Split(string(data), "\n")
	for _, lineData := range lineSlice {
		// 排除注释空行等行数据
		line := strings.TrimSpace(lineData)
		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		// 获取section对应的结构体字段名
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			fieldName, _ = getFieldName(line, typeInfo.Elem())
			continue
		}
		// 处理其他行，将k = v 对设置到结构体对应的属性上
		err := setField(fieldName, line, v)
		if err != nil {
			logger.Println(err)
			return err
		}
	}
	return nil
}
