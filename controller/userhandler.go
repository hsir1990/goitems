package controller

import (
	"encoding/json"
	"fmt"
	"goitems/dao"
	"html/template"
	"net/http"
	"strconv"
)

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	age, err := strconv.Atoi(r.PostFormValue("age"))
	if err != nil {
		fmt.Println("err-----", err)
	}
	//调用userdao中验证用户名和年龄的方法
	user, _ := dao.CheckUserNameAndPassword(username, age)
	fmt.Println("获取的User是: ", user)
	if user.Id > 0 {
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/pages/page.html"))

		t.Execute(w, user)
	} else {
		//用户名和密码不正确
		//用户名和密码正确
		t := template.Must(template.ParseFiles("views/index.html"))

		t.Execute(w, "查讯不到")
	}
}

//checkUserName通过ajax来进行
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user, err := dao.CheckUserName(username)
	if err != nil {
		fmt.Println("err----", err)
	}

	if user.Id > 0 {
		result, _ := json.Marshal(user) //序列化成字符串以后,会变成小写开头的
		w.Write([]byte(result))
	} else {
		w.Write([]byte("用户名不存在"))
	}

}
