package main

import (
	"fmt"
	"os"
)

type ManageData interface {
	Save(key, value string) error
	Get(key string) (string, error)
}

type Memory struct {
	Data map[string]string
}

func (d *Memory) Save(key, value string) error {
	d.Data[key] = value
	return nil
}

func (d *Memory) Get(key string) (string, error) {
	return d.Data[key], nil
}

type File struct {}

func (f *File) Save(key, value string) error {
	// key to be the file name
	return os.WriteFile(key, []byte(value), 0644)
}

func(f *File) Get(key string) (string, error) {
	val, err := os.ReadFile(key)
	return string(val), err
}

type DataProcessing struct {
	Mode ManageData
	Key string
	Value string
}

func InitProcessing(m ManageData, key, value string) *DataProcessing {
	return &DataProcessing{
		Mode: m,
		Key: key,
		Value: value,
	}
}

func(dp *DataProcessing) Process() {

	if err := dp.Mode.Save(dp.Key, dp.Value); err!=nil {
		fmt.Println("Data Not saved for Key:", dp.Key, "Error:", err.Error())
		return
	}
	fmt.Println("Read Saved")

	val,err := dp.Mode.Get(dp.Key)
	if err!=nil {
		fmt.Println("Data could not fetched for Key:", dp.Key, "Error:", err.Error())
		return
	}

	fmt.Println("Read Data", val)
}

func InterfaceDIImplementation2(){

	memo := Memory{Data: make(map[string]string)}
	memoP := InitProcessing(&memo, "name", "Harsh Lamba")
	memoP.Process()

	fil := File{}
	filP := InitProcessing(&fil, "name.txt", "Hi my name is harsh lamba")
	filP.Process()

}
