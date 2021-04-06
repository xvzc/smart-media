package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	var num int
	fmt.Println("*TODO 테이블 사용하기")

	fmt.Println("1. Insert")
	fmt.Println("2. Select")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")

	for {
		fmt.Println("번호 앱력 : ")
		fmt.Scanln(&num)
		switch num {
		case 1:
			data_insert(db)
		case 2:
			data_select(db)
		case 3:
			data_update(db)
		case 4:
			data_delete(db)
		}
	}
}

func data_insert(db_obj *sql.DB) {
	var user_id, password string
	fmt.Println("user_id : ") //user_id
	fmt.Scanln(&user_id)
	fmt.Println("password : ") //password
	fmt.Scanln(&password)

	// INSERT 문 실행
	result, err := db_obj.Exec("INSERT INTO users VALUES (?, ?)", user_id, password)
	checkError(err)

	n, err := result.RowsAffected()
	checkError(err)
	fmt.Println(n, "row inserted.")

	return
}

func data_select(db_obj *sql.DB) {
	rows, err := db_obj.Query("SELECT * FROM users")
	checkError(err)
	defer rows.Close()

	var user_id, password string

	for rows.Next() {
		err = rows.Scan(&user_id, &password)
		checkError(err)

		fmt.Println(user_id, password)
	}

}
func data_update(db_obj *sql.DB) {
	var uid, new_user_id, new_password string

	fmt.Println("user_id to update : ")
	fmt.Scanln(&uid)

	fmt.Println("new_user_id : ") //user_id
	fmt.Scanln(&new_user_id)
	fmt.Println("new_password : ") //password
	fmt.Scanln(&new_password)

	stmt, err := db_obj.Prepare("UPDATE users SET user_id=?, password=? WHERE user_id=?")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(new_user_id, new_password, uid)
	checkError(err)
	defer stmt.Close()

	fmt.Println("new info")
	fmt.Println(new_user_id, new_password)
}
func data_delete(db_obj *sql.DB) {
	var uid string

	fmt.Println("user_id to be deleted: ")
	fmt.Scanln(&uid)

	stmt, err := db_obj.Prepare("Delete from users WHERE user_id=?")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(uid)
	checkError(err)
	defer stmt.Close()

}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
