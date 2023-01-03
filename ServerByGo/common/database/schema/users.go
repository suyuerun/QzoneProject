package schema

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
	Country  string `db:"country"`
	City     string `db:"city"`
	Telcode  int    `db:"telcode"`
}

var PersonSchema *Person
