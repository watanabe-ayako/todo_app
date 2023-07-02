package controllers

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"todo_go/app/models"
)

// func top(w http.ResponseWriter, r *http.Request) {
// 	// ログインしているかチェック
// 	_, err := session(w, r)
// 	if err != nil {
// 		generateHTML(w, "Hello", "layout", "public_navbar", "top")
// 	} else {
// 		http.Redirect(w, r, "/todos", 302)
// 	}
// }

func index(w http.ResponseWriter, r *http.Request) {
	// セッションがあるかチェック
	sess, err := session(w, r)
	if err != nil { // エラーがある場合はログインしていないのでloginにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession() // sessionからユーザーを特定
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser() // ユーザーのtodoを取得
		// Status順にソート
		sort.Slice(todos, func(i,j int)bool {
			return todos[i].Status < todos[j].Status
		})
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// ParseFormで解析後フォームから値を受け取り、各フィールドに当てはめて保存。
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		// strStatus := r.PostFormValue("status")
		status, err := strconv.Atoi(r.PostFormValue("status"))
		if err != nil {
			log.Println(err)
		}
		if err := user.CreateTodo(content, status); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		content := r.PostFormValue("content")
		strStatus := r.PostFormValue("status")
		status, err := strconv.Atoi(strStatus)
		if err != nil {
			log.Println(err)
		}

		t := &models.Todo{ID: id, Content: content, Status: status, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}

}

func sortTodo(w http.ResponseWriter, r *http.Request, status int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, err := models.GetTodosWithStatus(user.ID, status)
		if err != nil {
			log.Println(err)
		}
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
