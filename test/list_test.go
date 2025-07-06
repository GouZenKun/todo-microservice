package testing

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	dbmodels "todo_module/internal/repository/db/models"
	"todo_module/internal/service"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/stretchr/testify/require"
)

func TestListTodo(t *testing.T) {
	const (
		dbDriver = "mysql"
		dbSource = "root:root@tcp(localhost:3306)/db_todo?charset=utf8mb4&parseTime=true"
	)
	con, err := sql.Open(dbDriver, dbSource) //DBコネクションを設定
	if err != nil {
		log.Println(err)
	}
	defer con.Close()
	boil.SetDB(con)

	filterByStatusId := uint64(0) // Keep it 0 for no filtering
	isSortByDue := false

	todos, err := service.GetInstance().ListTodo(context.Background(), filterByStatusId, isSortByDue)

	require.NoError(t, err)
	require.NotEmpty(t, todos)

	todosdb, err := dbmodels.Todos(qm.Where("del_flag=?", false)).All(context.Background(), con)

	require.NoError(t, err)
	require.NotEmpty(t, todosdb)

	require.Equal(t, len(todos), len(todosdb))
	for i, todo := range todos {
		require.Equal(t, todo.TodoId, fmt.Sprintf("%v", todosdb[i].TodoID))
		require.Equal(t, todo.Title, todosdb[i].Title)
		require.Equal(t, todo.Description, todosdb[i].Description.String)
		require.Equal(t, todo.StatusId, todosdb[i].StatusID)
		require.Equal(t, todo.DueDate, todosdb[i].DueDate.Time.Unix())
	}
}

func TestListTodo_FilterByStatusId(t *testing.T) {
	const (
		dbDriver = "mysql"
		dbSource = "root:root@tcp(localhost:3306)/db_todo?charset=utf8mb4&parseTime=true"
	)
	con, err := sql.Open(dbDriver, dbSource) //DBコネクションを設定
	if err != nil {
		log.Println(err)
	}
	defer con.Close()
	boil.SetDB(con)

	filterByStatusId := uint64(1) // Manually config
	isSortByDue := false

	todos, err := service.GetInstance().ListTodo(context.Background(), filterByStatusId, isSortByDue)

	require.NoError(t, err)
	require.NotEmpty(t, todos)

	todosdb_filtered, err := dbmodels.Todos(qm.Where("status_id=?", filterByStatusId)).All(context.Background(), con)

	require.NoError(t, err)
	require.NotEmpty(t, todosdb_filtered)

	require.Equal(t, len(todos), len(todosdb_filtered))
	for i, todo := range todos {
		require.Equal(t, todo.TodoId, fmt.Sprintf("%v", todosdb_filtered[i].TodoID))
		require.Equal(t, todo.Title, todosdb_filtered[i].Title)
		require.Equal(t, todo.Description, todosdb_filtered[i].Description.String)
		require.Equal(t, todo.StatusId, todosdb_filtered[i].StatusID)
		require.Equal(t, todo.DueDate, todosdb_filtered[i].DueDate.Time.Unix())
	}
}

func TestListTodo_SortByDue(t *testing.T) {
	const (
		dbDriver = "mysql"
		dbSource = "root:root@tcp(localhost:3306)/db_todo?charset=utf8mb4&parseTime=true"
	)
	con, err := sql.Open(dbDriver, dbSource) //DBコネクションを設定
	if err != nil {
		log.Println(err)
	}
	defer con.Close()
	boil.SetDB(con)

	filterByStatusId := uint64(0) // Keep it 0 for no filtering
	isSortByDue := true

	todos, err := service.GetInstance().ListTodo(context.Background(), filterByStatusId, isSortByDue)

	require.NoError(t, err)
	require.NotEmpty(t, todos)

	todosdb_sorted, err := dbmodels.Todos(qm.Where("del_flag=?", false), qm.OrderBy("due_date")).All(context.Background(), con)

	require.NoError(t, err)
	require.NotEmpty(t, todosdb_sorted)

	require.Equal(t, len(todos), len(todosdb_sorted))
	for i, todo := range todos {
		require.Equal(t, todo.TodoId, fmt.Sprintf("%v", todosdb_sorted[i].TodoID))
		require.Equal(t, todo.Title, todosdb_sorted[i].Title)
		require.Equal(t, todo.Description, todosdb_sorted[i].Description.String)
		require.Equal(t, todo.StatusId, todosdb_sorted[i].StatusID)
		require.Equal(t, todo.DueDate, todosdb_sorted[i].DueDate.Time.Unix())
	}
}
