package dbutils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tobgu/qframe"
	qsql "github.com/tobgu/qframe/config/sql"

	// driver de sqlite3
	_ "github.com/mattn/go-sqlite3"
)

//OpenSQLite open connection to sqlite3 database and returns a conection *sql.DB type
func OpenSQLite(dbpath string) *sql.DB {
	con, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		fmt.Println(err)
		panic("Impossível abrir conexão")
	}
	return con
}

//ImportQuery imports a query from file (queryname) and returns a string content
func ImportQuery(queryname string) string {
	content, err := ioutil.ReadFile(queryname)
	if err != nil {
		fmt.Println(err)
		panic("Could not read query")
	}
	return string(content)
}

//FormatQuery makes a placeholder in a query (string) using key/value map
func FormatQuery(query string, params map[string]string) string {
	for key, value := range params {
		query = strings.ReplaceAll(query, key, value)
	}
	return query
}

//ExecSQLFile executes a query from file directly to a database and returns a possible error
func ExecSQLFile(filename string, con *sql.DB, params map[string]string) error {
	query := FormatQuery(ImportQuery(filename), params)
	con.Exec(query)
	return nil
}

//ExecQuery executes a query to database and returns a Dataframe struct (qframe)
func ExecQuery(query string, con *sql.DB) qframe.QFrame {
	tx, err := con.Begin()
	if err != nil {
		panic("Error to use con.Begin()")
	}
	return qframe.ReadSQL(tx, qsql.Query(query))
}

//ExecQueryFile reads a query from file, executes in into database and return a dataframe result from query
func ExecQueryFile(filename string, con *sql.DB, params map[string]string) qframe.QFrame {
	query := FormatQuery(ImportQuery(filename), params)
	return ExecQuery(query, con)
}
