package models
import "fmt"

type User struct {
	Id 		 int
	Username string
	Password string
	Mail 	 string
	
}

func CreateUser() *User {
	user:=&User{}
	return user
}

func (db *DB) CheckUser(username string) bool {
	var name string
	err := db.QueryRow("SELECT username FROM users WHERE username=?",username).Scan(&name)
	if err!=nil {
		fmt.Println("Error: ", err.Error())
		return false
	}
	if name == username {
		return true
	}

	return false
}


func (db *DB) Login(username, password string) (bool, *User) {
	stmt, err := db.Prepare("SELECT * FROM users WHERE username=? && password=?")
	if err!=nil {
		return false, nil
	}
	rows, err := stmt.Query(username, password)
	if err!=nil {
		return false, nil
	}

	user:=&User{}
	for rows.Next() {

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Mail)

		if err!=nil {
			return false, nil
		}
		if username==user.Username && password==user.Password {
			return true, user
		} 
	}


	return false, nil
}


func (db *DB) RegisterUser(Username, Password, Mail string) bool {
	user:=new(User)
	user.Username = Username
	user.Password = Password
	user.Mail = Mail
	//DB
	stmt, err := db.Prepare("INSERT users SET username=?,password=?, mail=?")
	if err!=nil {
			fmt.Println("Error: ", err.Error())
			return false
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Mail)
	if err!=nil {
			fmt.Println("Error: ", err.Error())
			return false
	}
	
	
	return true
}


func (db *DB) GetAllUsers() []*User {
	rows, err := db.Query("SELECT * FROM users")
	if err!=nil {
			panic(err)
			return nil
		}	
	users:=make([]*User, 0 )
	for rows.Next() {
		user:=new(User)
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Mail)
		if err!=nil {
			panic(err)
			return nil
		}		
		users = append(users, user)
	}
	return users
}

