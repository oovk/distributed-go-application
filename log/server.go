package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

//takes the content of http requests coming in and write the content of http requests to the log

var log *stlog.Logger //handle the logging for our application

type fileLog string //handle actual writing to the file system

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) //path where we gonna write log to
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) { //function for logger to point to
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

func write(message string) {
	log.Printf("%v\n", message)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg, err := ioutil.ReadAll(r.Body)
		if err != nil || len(msg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}
