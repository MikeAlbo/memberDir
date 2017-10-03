package dbService

import (
	"os"
)

const demoDbPath  = "../data/demoJSON.json"

func SetData(data []byte) (string, error) {
	n, err := writeToFile(data)
	if err != nil {
		return n, err
	}
	return  n, nil
}

//func openFile()  {
//
//}

func writeToFile(data []byte) (string, error)  {
	jsonFile, err := os.Create(demoDbPath)
	if err != nil {
		return jsonFile.Name(), err
	}
	defer jsonFile.Close()
	jsonFile.Write(data)
	jsonFile.Close()
	return jsonFile.Name(), nil
}
