package main

import (
	"io/ioutil"
	"regexp"
	"strings"

	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	_ "os"
)

const (
	path           = "/"
	defaultLogFile = "/tmp/webhook-log.txt"
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {

	// Set variables from environment
	logFileName := getEnv("LOGS_FILE", defaultLogFile)
	log.Println("Using Log file [", logFileName, "]")

	// Open Log file
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	checkErr(err)

	// Send logs to STDOUT and Log file
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		// headers := make([]string, len(r.Header))

		// i := 0
		// for header := range r.Header {
		// 	headers[i] = header
		// 	i++
		// 	fmt.Printf("%s: %s\n", header, r.Header[header][0])
		// }

		if r.Body != nil {
			bodyBytes, err = ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Body reading error: %v", err)
				return
			}
			defer r.Body.Close()
		}
		if len(bodyBytes) > 0 {
			if bodyJSON, err := json.Marshal(bodyBytes); err != nil {
				fmt.Printf("JSON parse error: %v", err)
				return
			} else {
				str1 := string(bodyJSON[1 : len(bodyJSON)-1]) // body is double qouted
				x, _ := b64.StdEncoding.DecodeString(str1)
				// fmt.Printf("%s\n", string(x))
				s := strings.ReplaceAll(string(x), "\"labels\"", "\n{\"labels\"")
				s2 := strings.ReplaceAll(s, "}}", "}}\n")
				//fmt.Println(s2)
				re := regexp.MustCompile(`{"labels":.*\n`)
				s3 := re.FindAll([]byte(s2), -1)
				for _, value := range s3 {
					fmt.Printf("%s", value)
				}

			}
		} else {
			fmt.Printf("Body: No Body Supplied\n")
		}
	})
	http.ListenAndServe("0.0.0.0:4000", nil)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
