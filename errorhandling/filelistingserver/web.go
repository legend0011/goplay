package main

import (
	"fmt"
	"net/http"
	"os"

	"hello/errorhandling/filelistingserver/filelisting"
)

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		// some error we can know
		if userErr, ok := err.(userError); ok { // type assertion: https://go.dev/tour/methods/15
			http.Error(w,
				userErr.Message(),
				http.StatusBadRequest)
			return
		}

		code := http.StatusOK
		switch {
		case os.IsNotExist(err):
			code = http.StatusNotFound
		case os.IsPermission(err):
			// http://localhost:8090/list/errorhandling/filelistingserver/filelisting/staticfiles/notpermitted.txt
			code = http.StatusForbidden
		default:
			code = http.StatusInternalServerError
		}
		http.Error(w, http.StatusText(code), code)
	}
	return
}

type requestHandler func(http.ResponseWriter, *http.Request) error

// first think how this been used. then think input and output
func wrapper(handler requestHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			// even error occured in handler(), the outside defer have change to recover from panic
			r := recover()
			if r != nil {
				if myErr, ok := r.(userError); ok { // type assertion: https://go.dev/tour/methods/15
					fmt.Printf("Warning! %s", myErr.Error())
					http.Error(writer,
						fmt.Sprintf("wrapper Handling: %v", myErr.Message()),
						http.StatusBadRequest)
				} else {
					fmt.Println("Fatal error!")
					panic("unknown error")
				}
			}
		}()
		err := handler(writer, request)

		//a := 0
		//fmt.Println(6 / a) // unknown error

		handleError(writer, err)
	}
}

type userError interface {
	error            // for system
	Message() string // for user
}

func main() {
	http.HandleFunc("/",
		wrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
