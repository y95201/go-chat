package main
import (
    "net/http"
)
func main(){
    euler()
}
func userLogin(writer http.ResponseWriter,request *http.Request)  {
    // io.WriteString(writer,"hello,world!123456")
    request.ParseForm()
    // moblie := request.PostForm.Get("name")
    passwd := request.PostForm.Get("passwd")
    loginok := false
    if passwd == "123456" {
        loginok = true
    }
    str := `{"code":0,"data":{"id":1,"token":"test"}}`
    if !loginok {
        str = `{"code":-1,"msg":"失败"}`
    }
    writer.Header().Set("Content-Type", "application/json")
    writer.WriteHeader(http.StatusOK)
    writer.Write([]byte(str))
}
func euler(){
    http.HandleFunc("/user/login",userLogin)
    http.ListenAndServe(":8080", nil)
}


