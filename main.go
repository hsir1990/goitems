package main

import (
	"html/template"
	"net/http" //获取链接请求
)

//IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("./views/index.html"))
	//执行
	t.Execute(w, "")
}
func main() {

	//需要设置静态文件路径,如果不设置 会拒绝加载静态文件
	//http.StripPrefix(prefix string, h Handler) Handler
	//StripPrefix返回一个处理器,该处理器会将请求的URL.Path字段中给定前缀prefix去除后,再交有h处理,StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static"))))
	// //直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./views/pages"))))

	http.HandleFunc("/main", IndexHandler)

	http.ListenAndServe(":8080", nil)
}

// must
// 英 [mʌst , məst]   美 [mʌst , məst]
// modal
// (表示必要或很重要)必须;(表示很可能或符合逻辑)一定;(提出建议)应该，得
// n.
// 必须做(或看、买等)的事
// 复数： musts过去式： must过去分词： must
