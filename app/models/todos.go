package models

import (
	"fmt"
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	Status    int
	UserID    int
	CreatedAt time.Time
}

const (
	StatusTodo = iota
	StatusInProgress
	StatusDone
)

// create Userのメソッドとして作る
func (u *User) CreateTodo(content string, status int) (err error) {
	cmd := `insert into todos (
		content,
		status,
		user_id,
		created_at) values (?, ?, ?, ?);`
	_, err = Db.Exec(cmd, content, status, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, status, user_id, created_at from todos where id =?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.Status,
		&todo.UserID,
		&todo.CreatedAt)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, status, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Status,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	fmt.Println(todos)
	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id , content, status, user_id, created_at from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Status,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

// ソート用
func GetTodosWithStatus(id int, status int) (todos []Todo, err error) {
	cmd := `select id, content, status, user_id, created_at from todos where user_id = ? and status = ?`
	rows, err := Db.Query(cmd, id, status)
	if err != nil {
		log.Fatalln("実行時エラー", err, cmd, id, status)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Status,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln("値アサート時エラー", err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?, status = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.Status, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err := Db.Exec(cmd, t.ID)
	if err != nil {
		log.Println(err)
	}
	return err
}
