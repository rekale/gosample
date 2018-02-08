package counts

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var countView int

func init() {
	var err error

	countView, err = readFromFile()
	if err != nil {
		log.Println(err)
		writeToFile(0)
	}
}

//IncrementsMiddleware increment the count view
func IncrementsMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	countView++
	writeToFile(countView)
	next(rw, r)
}

// Get the count view value
func Get() int {
	return countView
}

func readFromFile() (int, error) {
	data, err := ioutil.ReadFile("/tmp/count")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	countFile := string(data)

	count, err := strconv.Atoi(countFile)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func writeToFile(count int) {
	c := fmt.Sprintf("%v", count)
	err := ioutil.WriteFile("/tmp/count", []byte(c), 0644)
	if err != nil {
		log.Println(err)
	}
}
