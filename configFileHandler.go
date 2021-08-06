package main

import "io/ioutil"

func StructToFile(fileName string, v interface{}) error {
	bytes, err := INIMarshal(v)
	if err != nil {
		logger.Println(err)
		return err
	}
	err = ioutil.WriteFile(fileName, bytes, 0644)
	if err != nil {
		logger.Println(err)
		return err
	}
	return nil
}

func FileToStruct(fileName string, v interface{}) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Println(err)
		return err
	}
	err = INIUnMarshal(bytes, v)
	if err != nil {
		logger.Println(err)
		return err
	}
	return nil
}
