package provider

import (
	"fmt"
	Db "gin/common/database"
	_ "gin/common/database/schema"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
	Country  string `db:"country"`
	City     string `db:"city"`
	Telcode  int    `db:"telcode"`
}

func (p *Person) CreatePerson(personData map[string]interface{}) string {
	sqlStr := "insert into person(username, password, sex, email, country, city, telcode)values(?, ?, ?, ?, ?, ?, ?)"

	r, err := Db.Database.Exec(sqlStr, "root", "admin", "male", "root@163.com", "China", "Beijing", "123456")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return "exec failed"
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return "LastInsertId failed"
	}

	fmt.Println("insert succ:", id)
	return "success"
}

func (p *Person) GetPerson(username string) *Person {
	var person []*Person
	sqlStr := "select username, password, sex, email, country, city, telcode from person where username=?"
	err := Db.Database.Select(&person, sqlStr, username)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return nil
	}
	fmt.Println("select succ:", person)

	return person[0]
}
