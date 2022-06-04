package handleRequests

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var autoIncrement int

//Set is a function that can be called to set key and value in database
func Set(k string, v interface{}, m map[string]interface{}) {
	m[k] = v
}

//Get value of key in map
func Get(k string, m map[string]interface{}) interface{} {
	return m[k]
}

// Del Delete key and value from map
func Del(k string, m map[string]interface{}) {
	delete(m, k)
}

//Reg Query on keys over custom map
func Reg(reg string, m map[string]interface{}) []string {
	var output []string
	for s, _ := range m {
		found, _ := regexp.Match(reg, []byte(s))
		if found {
			output = append(output, s)
		}
	}
	return output
}

//Use create new database
func Use(s string, list map[string]map[string]interface{}) (map[string]interface{}, bool) {
	if list[s] != nil {
		return list[s], true
	}
	return make(map[string]interface{}), false

}

//List print list of databases
func List(m map[string]map[string]interface{}) {
	fmt.Printf("> [")
	length := len(m) - 1
	temp := 0
	for i, _ := range m {
		if temp != length {
			fmt.Printf("%q, ", i)
		} else {
			fmt.Printf("%q", i)
		}
		temp++
	}
	fmt.Printf("]\n")
}

//Dump write a database data into .csv file
func Dump(path string, m map[string]interface{}) {
	temp := strings.Split(path, "/")
	newPath := temp[:len(temp)-1]
	if _, err := os.Stat(strings.Join(newPath, "/")); err != nil {
		os.Mkdir(strings.Join(newPath, "/"), 0755)
		file, _ := os.Create(path + "_" + strconv.Itoa(autoIncrement/100) + ".csv")
		w := csv.NewWriter(file)
		defer w.Flush()
	}
	file, _ := os.OpenFile(path+"_"+strconv.Itoa(autoIncrement/100)+".csv", os.O_CREATE|os.O_RDWR, 0755)
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	var data [][]string

	for k, _ := range m {
		row := []string{k, m[k].(string)}
		row[0] = strconv.Itoa(autoIncrement) + "$" + row[0]
		autoIncrement++
		data = append(data, row)
	}
	err := w.WriteAll(data)
	if err != nil {
		log.Println(err)
	}

}

//Load read data from .csv and load it into Ram
func Load(path string, m map[string]interface{}) {
	for i := 0; i < autoIncrement/100+1; i++ {
		file, err := os.Open(path + "_" + strconv.Itoa(autoIncrement) + ".csv")
		if err != nil {
			fmt.Println(err)
		}
		reader := csv.NewReader(file)
		records, _ := reader.ReadAll()
		for _, record := range records {
			m[record[0]] = record[1]
		}
	}

}
