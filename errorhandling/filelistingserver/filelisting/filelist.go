package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (u userError) Error() string {
	return "for system: " + u.Message()
}

func (u userError) Message() string {
	return "for user: " + string(u)
}

const api = "/list/"

// Want to isolate this business logic, and throw error out to centralize error handling
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	a := 0
	fmt.Println(6 / a) // unknown error
	if strings.Index(request.URL.Path, api) != 0 {
		return userError("path must start with " + api)
	}
	path := request.URL.Path[len(api):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}
