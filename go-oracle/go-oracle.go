package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-goracle/goracle"
)

type QueryResults struct {
	Results []map[string]string `json:"results"`
}

func getDBConnection() *sql.DB {
	db_username := os.Getenv("UCAT_ORACLE_USERNAME")
	db_password := os.Getenv("UCAT_ORACLE_PASSWORD")
	db_host := os.Getenv("UCAT_ORACLE_HOST")
	db_port := os.Getenv("UCAT_ORACLE_PORT")
	db_sid := os.Getenv("UCAT_ORACLE_SID")
	connectStr := fmt.Sprintf("%s/%s@%s:%s/%s", db_username, db_password, db_host, db_port, db_sid)

	db, err := sql.Open("goracle", connectStr)
	if err != nil {
		panic(err)
	}
	return db
}

func rowMapString(columnNames []string, rows *sql.Rows) (map[string]string, error) {
	lenCN := len(columnNames)
	result := make(map[string]string, lenCN)
	columnPointers := make([]interface{}, lenCN)
	for i := 0; i < lenCN; i++ {
		columnPointers[i] = new(sql.RawBytes)
	}
	if err := rows.Scan(columnPointers...); err != nil {
		return nil, err
	}
	for i := 0; i < lenCN; i++ {
		if rb, ok := columnPointers[i].(*sql.RawBytes); ok {
			result[columnNames[i]] = string(*rb)
		} else {
			return nil, fmt.Errorf("Cannot convert index %d column %s to type *sql.RawBytes", i, columnNames[i])
		}
	}
	return result, nil
}

func QueryDB(query string) QueryResults {
	db := getDBConnection()
	defer db.Close()
	var queryResults QueryResults
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	columnNames, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		cv, err := rowMapString(columnNames, rows)
		if err != nil {
			panic(err)
		}
		queryResults.Results = append(queryResults.Results, cv)
	}
	return queryResults
}

func main() {
	results := QueryDB("select * from uc_product where ROWNUM < 10")
	for i, result := range results.Results {
		fmt.Println(i, result)
	}
}
