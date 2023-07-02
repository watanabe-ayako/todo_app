package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"todo_go/app/models"
	"todo_go/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// この関数でアクセス制限をかける
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// URLの正規表現のパターンをコンパイルする
var validPath = regexp.MustCompile("^/todos/(edit|update|delete|todo|inprogress|done)/([0-9]+)$")

// ハンドラ関数を返す関数
// 引数でtodoEdit関数などを受け取り、パースしたIDを渡して関数を実行
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1 こんなURLからIDを受け取りたい
		// validPathとマッチした部分をスライスで取得
		q := validPath.FindStringSubmatch(r.URL.Path)
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi) // 引数で受け取った関数を実行
	}
}

// ハンドラに登録したviewをここでURLに登録
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLとコントローラのハンドラ関数を渡す
	// 末尾に/がついていないと完全一致。/終わりの場合は先頭が登録されたURLと一致するかどうか。
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit)) // parseURL経由でtodoEditを実行
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	http.HandleFunc("/todos/", parseURL(sortTodo))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
