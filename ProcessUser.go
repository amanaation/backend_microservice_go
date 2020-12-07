package main

import (
	"context"
	"strconv"
	"time"

	valid "github.com/asaskevich/govalidator"
)

type readOutput struct {
	uid           []int
	lname         []string
	name          []string
	creditBalance []int
	creditLimit   []int
}

func user(data map[string]string) {

	var values = make(map[string]string)
	var operation string
	var v string
	for k, v2 := range data {
		v = v2
		if k == "type" {
			operation = v
		} else {
			if !valid.IsInt(v) {
				v = "'" + v + "'"
			}
			values[k] = v
		}

	}
	switch operation {
	case "C":
		insertUser(values)

	case "U":
		updateUser(values)
	case "D":
		deleteUser(values)
	}

}

func getNewUserID() int {
	// defer db.Close()

	rows, err := db.Query("SELECT count(*) FROM User")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var id int
	for rows.Next() {

		// var lname, name string
		err = rows.Scan(&id)
		if err != nil {
			// handle this error
			panic(err)
		}
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return id + 1
}

func insertUser(data map[string]string) string {
	//data := map[string]string{"uid": "3", "lname": "Mishra", "name": "Akash", "creditLimit": "100000",
	//	"creditBalance": "100000"}

	println(data)
	var query string
	query = "Insert into User"

	columns := ""
	values := ""

	for k, v := range data {
		columns += k + ", "
		if !valid.IsInt(v) {
			v = "'" + v + "'"
		}
		values += v + ", "
	}

	columns = "(" + columns[:len(columns)-2] + ")"
	values = "(" + values[:len(values)-2] + ");"

	query += columns + " values " + values
	return insertUpdateDelete(query)
}

func insertUpdateDelete(query string) string {
	// println(query)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return "Error when preparing SQL statement"
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx)
	if err != nil {
		return "Error when performing operation on table"
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return "Error when finding rows affected : "
	}
	// println(strconv.Itoa(int(rows)) + " products created ")
	return strconv.Itoa(int(rows)) + " products created "
}

func updateUser(data map[string]string) string {
	// http://localhost:8080/user?type=U&setname=aditya&uid=2&setlname=mishra
	query := "Update User "
	setQuery := " set "
	whereClause := " where "
	for k, v := range data {
		if "set" == string(k)[:3] {
			setQuery += string(k)[3:] + " = " + string(v) + " , "
		} else {
			whereClause += string(k) + " = " + string(v) + " and "
		}

	}

	setQuery = setQuery[:len(setQuery)-2]
	whereClause = whereClause[:len(whereClause)-4]

	if setQuery != " set " {
		query += setQuery
	}
	if whereClause != " where " {
		query += whereClause
	}
	return insertUpdateDelete(query)

}

func deleteUser(data map[string]string) string {
	// http://localhost:8080/user?type=D&uid=2&lname=mishra
	query := " Delete from User "
	whereClause := " where "
	for k, v := range data {
		whereClause += string(k) + " = " + string(v) + " and "
	}

	if whereClause != " where " {
		whereClause = whereClause[:len(whereClause)-4]
		query += whereClause
	}
	return insertUpdateDelete(query)

}

// func readUser(data map[string]string) {

// 	defer db.Close()

// 	query := "SELECT * FROM User"
// 	if len(data) != 0 {
// 		query += " where "
// 	}

// 	for k, v := range data {
// 		v2 := string(v[0])
// 		if !valid.IsInt(v2) {
// 			v2 = "'" + v2 + "'"
// 		}
// 		query += string(k) + " = " + v2 + " and "
// 	}

// 	if query != "SELECT * FROM User" {
// 		query = query[:len(query)-4]
// 	}
// 	println(query)
// 	count_query := "Select count(*) from (" + query + ") as T"

// 	println(count_query)
// 	rows, err := db.Query(count_query)
// 	if err != nil {
// 		// handle this error better than this
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	var count int
// 	for rows.Next() {
// 		err = rows.Scan(&count)
// 	}
// 	println("...........", count)
// 	suid := make([]int, count)
// 	// println(" *****", suid)
// 	fmt.Printf("%v", suid)
// 	// slname := make([]string, count)
// 	// sname := make([]string, count)
// 	// screditLimit := make([]int, count)
// 	// screditBalance := make([]int, count)

// 	rows2, err2 := db.Query(query)
// 	if err != nil {
// 		// handle this error better than this
// 		panic(err2)
// 	}
// 	defer rows2.Close()

// 	for rows2.Next() {
// 		var id, climit, cbalance int
// 		var lname, name string
// 		err = rows.Scan(&id, &lname, &name, &climit, &cbalance)
// 		suid = append(suid, id)
// 		println(" ---- ", id)
// 		// slname := append(slname, lname)
// 		// sname := append(sname, name)
// 		// screditBalance := append(screditBalance, id)
// 		// screditLimit := append(screditLimit, id)

// 		if err != nil {
// 			// handle this error
// 			panic(err2)
// 		}
// 		fmt.Println(id, lname, name, cbalance, climit)
// 		//x[id] = append()
// 	}
// 	println(reflect.ValueOf(suid).Kind())
// 	o := readOutput{uid: suid}
// 	fmt.Printf("%v", o.uid)
// 	// get any error encountered during iteration
// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}

// }
