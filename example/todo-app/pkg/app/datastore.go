package app

import (
	"database/sql"

	"github.com/fossildev/todo-app/pkg/database"
)

// TodoDatastore interface
type TodoDatastore interface {
	Get(id int64) (*Todo, error)
	GetAll() ([]*Todo, error)
	Insert(todo *Todo) (int64, error)
	Update(id int64, todo *Todo) error
	Delete(id int64) error
}

type datastore struct {
	DB *sql.DB
}

// NewTodoDatastore create new todo datastore instance
func NewTodoDatastore() TodoDatastore {
	ds := new(datastore)
	ds.DB = database.NewDatabase().Connect()
	return ds
}

func (ds *datastore) Get(id int64) (*Todo, error) {
	todo := new(Todo)

	if err := ds.DB.QueryRow("SELECT id, title, completed, createtime, updatetime FROM todos WHERE id=?", id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreateTime, &todo.UpdateTime); err != nil {
		return nil, err
	}

	return todo, nil
}

func (ds *datastore) GetAll() ([]*Todo, error) {
	var todos []*Todo

	rows, err := ds.DB.Query("SELECT id, title, completed, createtime, updatetime FROM todos")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreateTime, &todo.UpdateTime); err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}

func (ds *datastore) Insert(todos *Todo) (int64, error) {
	result, err := ds.DB.Exec("INSERT INTO todos (title) VALUES (?)", todos.Title)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (ds *datastore) Update(id int64, todos *Todo) error {
	if _, err := ds.DB.Exec("UPDATE todos SET completed=? WHERE id=?", todos.Completed, id); err != nil {
		return err
	}
	return nil
}

func (ds *datastore) Delete(id int64) error {
	if _, err := ds.DB.Exec("DELETE FROM todos WHERE id=?", id); err != nil {
		return err
	}

	return nil
}
