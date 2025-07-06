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

func TestReadTodo(t *testing.T) {
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

	id := "1" // TODO : try using dummy data for test or something

	tododb, err := dbmodels.Todos(qm.Where("todo_id=?", id)).One(context.Background(), con)

	require.NoError(t, err)
	require.NotEmpty(t, tododb)

	todo, err := service.GetInstance().ReadTodo(context.Background(), id)

	if tododb.DelFlag {
		require.Error(t, err)
		require.Empty(t, todo)
	} else {
		require.NoError(t, err)
		require.NotEmpty(t, todo)

		require.Equal(t, id, fmt.Sprintf("%d", tododb.TodoID))
		require.Equal(t, todo.Title, tododb.Title)
		require.Equal(t, todo.Description, tododb.Description.String)
		require.Equal(t, todo.StatusId, tododb.StatusID)
		require.Equal(t, todo.DueDate, tododb.DueDate.Time.Unix())
	}

}
