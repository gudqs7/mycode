package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	text "text/template"
	"io"
	"os"
)

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w,"Hello gudqs!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w,nil)
	} else {
		r.ParseForm()
		fmt.Println("userName:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
//		temp := template.HTMLEscapeString(r.Form.Get("username"))
		text.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
//		t.ExecuteTemplate(w, "T", r.Form.Get("username"))
		fmt.Fprintf(w, r.Form.Get("username"))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}
