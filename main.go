package main

import (
	"bufio"
	"fmt"
	h "github.com/mrezaj79/redisak/handleRequests"
	"os"
	"strings"
)

func customPrint() {
	fmt.Printf("> \n")
}

func main() {
	var input, key, value string
	activeDatabase := make(map[string]interface{})
	databaseList := make(map[string]map[string]interface{})
	databaseList["default"] = activeDatabase

	reader := bufio.NewReader(os.Stdin)
	for input != "exit" {
		fmt.Scanf("%s", &input)
		switch input {
		case "set":
			fmt.Scanf("%s", &key)
			value, _ = reader.ReadString('\n')
			value = strings.TrimSuffix(value, "\n")
			h.Set(key, value, activeDatabase)
			customPrint()
		case "get":
			fmt.Scanf("%s", &key)
			output := h.Get(key, activeDatabase)
			if output != nil {
				fmt.Printf("> %s \n", output)
			} else {
				customPrint()
			}

		case "del":
			fmt.Scanf("%s", &key)
			h.Del(key, activeDatabase)
			customPrint()
		case "keys":
			var reg string
			fmt.Scanf("%s", &reg)
			p := h.Reg(reg, activeDatabase)
			fmt.Printf("> [")
			for i, v := range p {
				if i != len(p)-1 {
					fmt.Printf("%q ,", v)
				} else {
					fmt.Printf("%q", v)
				}
			}
			fmt.Printf("]\n")

		case "use":
			fmt.Scanf("%s", &key)
			newMap, isFound := h.Use(key, databaseList)
			if isFound {
				activeDatabase = databaseList[key]
			} else {
				activeDatabase = newMap
				databaseList[key] = newMap
			}
			customPrint()
		case "list":
			h.List(databaseList)
		case "dump":
			fmt.Scanf("%s", &key)
			value, _ = reader.ReadString('\n')
			value = strings.TrimSuffix(value, "\n")
			h.Dump(value, databaseList[key])
			customPrint()
		case "load":
			fmt.Scanf("%s", &key)
			value, _ = reader.ReadString('\n')
			value = strings.TrimSuffix(value, "\n")
			activeDatabase = databaseList[value]
			h.Load(key, activeDatabase)
			customPrint()
		case "exit":
			customPrint()
			return
		}

	}
}
