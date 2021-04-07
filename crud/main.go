package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// "encoding/json" //json형태로 출력해주는 패키지
)

type TODO struct {
	userId    string
	startDate string
	endDate   string
	title     string
	status    string
}

type USER struct {
	userId   string
	password string
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	var num int

	for {
		fmt.Println("")
		fmt.Println("**********************")
		fmt.Println("*TODO 테이블 사용하기*")
		fmt.Println("**********************")

		fmt.Println("1. Insert")
		fmt.Println("2. Select")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("")

		fmt.Print("번호 입력 : ")
		fmt.Scanln(&num)
		fmt.Println("")
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
	todo := TODO{}

	fmt.Print("user id : ") //user_id
	fmt.Scanln(&todo.userId)
	fmt.Print("start date: ") //password
	fmt.Scanln(&todo.startDate)
	fmt.Print("end date: ") //password
	fmt.Scanln(&todo.endDate)
	fmt.Print("title: ") //password
	fmt.Scanln(&todo.title)
	fmt.Print("status: ") //password
	fmt.Scanln(&todo.status)

	// INSERT 문 실행
	result, err := db_obj.Exec(`INSERT INTO todo 
								VALUES (?, ?, ?, ?, ?)`,
		todo.userId, todo.startDate, todo.endDate, todo.title, todo.status)

	checkError(err)

	n, err := result.RowsAffected()
	checkError(err)
	fmt.Println(n, "row inserted.")

	return
}

func data_select(db_obj *sql.DB) {

	var uid string
	todo := TODO{}

	fmt.Print("Find user by id : ")
	fmt.Scanln(&uid)
	fmt.Println("")

	rows, err := db_obj.Query(`SELECT * FROM todo where user_id like ?`, "%"+uid+"%")
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&todo.userId, &todo.startDate, &todo.endDate, &todo.title, &todo.status)
		checkError(err)

		fmt.Println(todo.userId, todo.startDate, todo.endDate, todo.title, todo.status)
	}

}
func data_update(db_obj *sql.DB) {
	var uid string
	new_data := TODO{}

	fmt.Println("Search by user id: ")
	fmt.Scanln(&uid)

	fmt.Print("\nNew user id : ")
	fmt.Scanln(&new_data.userId)
	fmt.Print("New start date: ")
	fmt.Scanln(&new_data.startDate)
	fmt.Print("New end date: ")
	fmt.Scanln(&new_data.endDate)
	fmt.Print("New title: ")
	fmt.Scanln(&new_data.title)
	fmt.Print("New status: ")
	fmt.Scanln(&new_data.status)

	stmt, err := db_obj.Prepare("UPDATE todo SET user_id=?, start_date=?, end_date=?, title=?, status=? WHERE user_id=?;")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(new_data.userId, new_data.startDate, new_data.endDate, new_data.title, new_data.status, uid)
	checkError(err)
	defer stmt.Close()

	fmt.Print("new info : ")
	fmt.Println(new_data.userId, new_data.startDate, new_data.endDate, new_data.title, new_data.status)
}

func data_delete(db_obj *sql.DB) {
	var uid string

	fmt.Print("Search by user id to delete from list : ")
	fmt.Scanln(&uid)
	fmt.Println("")

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
