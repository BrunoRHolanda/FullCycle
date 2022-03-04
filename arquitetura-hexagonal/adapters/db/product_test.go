package db_test

import (
	"database/sql"
	"log"
	"testing"

	db_adapter "github.com/BrunoRHolanda/FullCycle/go-hexagonal/adapters/db"
	"github.com/BrunoRHolanda/FullCycle/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := "create table if not exists products(id string, name string, price float, status string)"
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, "disabled")`
	stmt, err := Db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db_adapter.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db_adapter.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.0

	result, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetName(), result.GetName())
	require.Equal(t, product.GetPrice(), result.GetPrice())
	require.Equal(t, product.GetStatus(), result.GetStatus())

	product.Status = "enabled"

	result, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.GetStatus(), result.GetStatus())
}
