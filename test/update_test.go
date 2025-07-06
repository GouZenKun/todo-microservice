package testing

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"
	"testing"
	"time"
	todov1 "todo_module/internal/controller/gen/proto/v1"
	dbmodels "todo_module/internal/repository/db/models"
	"todo_module/internal/service"
	"todo_module/util"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	"github.com/stretchr/testify/require"
)

func TestUpdateTodo(t *testing.T) {
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

	arg := &todov1.Todo{
		TodoId:      id,
		Title:       util.GenerateRandomString(10),
		Description: util.GenerateRandomString(10),
		StatusId:    uint64(rand.IntN(3) + 1),
		DueDate:     time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Unix(),
	}

	tododb, err := dbmodels.Todos(qm.Where("todo_id=?", id)).One(context.Background(), con)

	require.NoError(t, err)
	require.NotEmpty(t, tododb)

	err = service.GetInstance().UpdateTodo(context.Background(), arg)

	if tododb.DelFlag {
		require.Error(t, err)
		require.Equal(t, tododb.DelFlag, true)
	} else {
		require.NoError(t, err)

		tododb2, err := dbmodels.Todos(qm.Where("todo_id=?", id)).One(context.Background(), con)

		require.NoError(t, err)
		require.NotEmpty(t, tododb2)

		require.Equal(t, id, fmt.Sprintf("%d", tododb2.TodoID))
		require.Equal(t, arg.Title, tododb2.Title)
		require.Equal(t, arg.Description, tododb2.Description.String)
		require.Equal(t, arg.StatusId, tododb2.StatusID)
		require.Equal(t, arg.DueDate, tododb2.DueDate.Time.Unix())

		require.NotZero(t, tododb2.CreateAt)
		require.NotZero(t, tododb2.UpdateAt)
		require.NotEqual(t, tododb.UpdateAt, tododb2.UpdateAt)
	}
}
