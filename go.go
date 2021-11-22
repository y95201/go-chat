package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	euler()
}

//go get github.com/go-xorm/xorm
//go get github.com/go-sql-driver/mysql
type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var DbEngin *xorm.Engine

func init() {
	drivename := "mysql"
	DsName := "root:123456@(127.0.0.1:3306)/think?"
	DbEngin, err := xorm.NewEngine(drivename, DsName)
	if err != nil {
		log.Fatal(err.Error())
	}
	//是否显示sql语句
	DbEngin.ShowSQL(true)
	//数据库最大链接数
	DbEngin.SetMaxOpenConns(2)
	fmt.Println("init data base ok")
}
func userLogin(writer http.ResponseWriter, request *http.Request) {
	// io.WriteString(writer,"hello,world!123456")
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	loginok := false
	if username == "admin" && password == "123456" {
		loginok = true
	}
	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(writer, 0, data, "成功")
	} else {
		Resp(writer, -1, nil, "密码不正确")
	}
}

func RegisterView() {
	tpl, err := template.ParseGlob("views/**/*")
	if nil != err {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(rw http.ResponseWriter, r *http.Request) {
			tpl.ExecuteTemplate(rw, tplname, nil)
		})
	}
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write([]byte(ret))
}

func euler() {
	http.HandleFunc("/user/login", userLogin)
	//渲染静态资源目录
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	//加载静态文件
	RegisterView()
	http.ListenAndServe(":8080", nil)
}
