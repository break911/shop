// furniture_shop project main.go
package main

import (
	"fmt"
	//	"html/template"
	"log"
	"net/http"
	//"encoding/json"
	//	"io"
	"os"
	//"strings"
	"errors"
)

type WebServerSettings struct {
	rootWWW_path string
	css_path     string
	tmpl_path    string
	data_path    string
}

var ws_settings = &WebServerSettings{}

func initWebServerSettings(path string) *WebServerSettings {
	return &WebServerSettings{rootWWW_path: path}
}

func (s *WebServerSettings) getWWWPath() string {
	return s.rootWWW_path
}

func (s *WebServerSettings) setCssPath(path string) {
	s.css_path = s.rootWWW_path + path
}

func (s *WebServerSettings) getCssPath() string {
	return s.css_path
}

func (s *WebServerSettings) setTmplPath(path string) {
	s.tmpl_path = s.rootWWW_path + path
}

func (s *WebServerSettings) getTmplPath() string {
	return s.tmpl_path
}

func (s *WebServerSettings) setDataPath(path string) {
	s.data_path = s.rootWWW_path + path
}

func (s *WebServerSettings) getDataPath() string {
	return s.data_path
}

func __getAllFSFromDir(path string, isdirectory bool) ([]string, error) {
	var files []string

	f, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}

	files_attr, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return []string{}, err
	}

	for _, file := range files_attr {
		if file.IsDir() == isdirectory {
			files = append(files, path+file.Name())
		}
	}

	return files, nil
}

func getAllFilesFromDir(path string) ([]string, error) {
	return __getAllFSFromDir(path, false)
}

func getAllSubDirFromDir(path string) ([]string, error) {
	return __getAllFSFromDir(path, true)
}

const (
	GET_CSS  = 0
	GET_TMPL = 1
	GET_DATA = 2
)

func getData(typeofdata int, namedata string) ([]byte, error) {

	fullpathtoresource := ""

	switch typeofdata {
	case GET_CSS:
		fullpathtoresource = ws_settings.getCssPath() + namedata
		break
	case GET_TMPL:
		fullpathtoresource = ws_settings.getTmplPath() + namedata
		break
	case GET_DATA:
		fullpathtoresource = ws_settings.getDataPath() + namedata
		break
	default:
		return []byte{}, errors.New("func getData(typeofdata int, namedata string) unknow typeofdata")
	}

	f, err := os.Open(fullpathtoresource)
	if err != nil {
		return []byte{}, err
	}
	info, err := f.Stat()
	if err != nil {
		return []byte{}, err
	}

	if info.IsDir() == true {
		return []byte{}, errors.New("func getData(typeofdata int, namedata string) request resource is directory")
	}
	data := make([]byte, info.Size())
	readsbyte, err := f.Read(data)
	if int64(readsbyte) != info.Size() {
		return []byte{}, errors.New("func getData(typeofdata int, namedata string) request resource is directory")
	}

	return data, nil
}

func WWWRoot(w http.ResponseWriter, r *http.Request) {
	subdir, err := getAllSubDirFromDir(ws_settings.getDataPath())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	lenpath := len(ws_settings.getDataPath())
	for _, s := range subdir {
		fmt.Fprintln(w, s)
		title, _ := getData(GET_DATA, s[lenpath:]+"/title.txt")
		fmt.Fprintln(w, string(title))
	}
}

func main() {
	ws_settings = initWebServerSettings("/home/break/go/src/www/")
	ws_settings.setTmplPath("tmpl/")
	ws_settings.setCssPath("css/")
	ws_settings.setDataPath("data/")

	http.HandleFunc("/", WWWRoot)
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
		template_files, err := getAllFilesFromDir(ws_settings.getTmplPath())
		if err != nil {
			log.Fatal(err)
		}

		templates = template.Must(template.

		for _, s := range template_files {
			fmt.Println(s)
		}*/
}
