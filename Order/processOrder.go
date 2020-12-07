package main

import (
	"context"
	"strconv"
	"time"

	valid "github.com/asaskevich/govalidator"
)

func order(data map[string]string) {

	var values = make(map[string]string)
	var v string
	for k, v2 := range data {
		v = v2
		if !valid.IsInt(v) {
			v = "'" + v + "'"
		}
		values[k] = v
	}

	initialProcessing(values)

}

func getNewID(idType string) int {
	//defer db.Close()

	query := "Select max(" + idType + ") FROM Order_details;"
	rows, err := db.Query(query)
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

func initialProcessing(data map[string]string) string {

	values := make(map[string]string)
	values["eventid"] = strconv.Itoa(getNewID("eventid"))
	values["orderid"] = strconv.Itoa(getNewID("orderid"))
	values["uid"] = data["uid"]
	values["order_notes"] = "Order initalised"
	values["order_status"] = "Initalised"
	values["order_amount"] = data["amount"]
	values["event_timestamp"] = time.Now().Format(time.RFC850)

	// println(values)
	insertOrder(values)
	query := "select credit_balance from User where uid = " + data["uid"]
	rows, err := db.Query(query)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	var creditBalance int
	for rows.Next() {
		err = rows.Scan(&creditBalance)
	}
	var status string
	amount, err := strconv.Atoi(data["amount"])
	if creditBalance < amount {
		status = "Insufficient balance. Connot Process Your Order!!"
		values["order_status"] = "Failed"
		values["order_notes"] = status
	} else {
		status = "Order Placed successfully!"
		values["order_status"] = "Completed"
		values["order_notes"] = status
	}
	values["eventid"] = strconv.Itoa(getNewID("eventid"))
	time.Sleep(3)
	insertOrder(values)
	return status

}

func insertOrder(data map[string]string) string {
	//data := map[string]string{"uid": "3", "lname": "Mishra", "name": "Akash", "creditLimit": "100000",
	//	"creditBalance": "100000"}

	var query string
	query = "Insert into Order_details"

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

func insertUpdateDelete2(query string) string {
	println(query)
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
	println(strconv.Itoa(int(rows)) + " products created ")
	return strconv.Itoa(int(rows)) + " products created "
}

func updateUser2(data map[string]string) string {
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
