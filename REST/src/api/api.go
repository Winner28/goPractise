package api

import (
	"github.com/gorilla/mux"
	"html/template"
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	
)
const (
	idNotFoundTrouble = 1
	userAlreadyExists = 2
	userNotExists = 3
)

type User struct {
	Id          string `json:id`
	Username    string `json:Username`
	OS_family   string `json:OS_family`
	OS 			string `json:OS`
	Shell 		string `json:Shell`
	Kernel 		string `json:Kernel`
	CPU 		string `json:CPU`
	Terminal 	string `json:Terminal`
}

func init () {
	fmt.Println("init")
	appendCustomUsers()
}

var users []User

func homepage(w http.ResponseWriter, r *http.Request) {
	t,err:=template.ParseFiles("local/homepage.html")
	if (err!=nil) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
		return
	}
	t.Execute(w,nil)
}

func showUsers(w http.ResponseWriter, r *http.Request) {
	encodeFullList(w)
}


func getUser(w http.ResponseWriter, r *http.Request) {
	params:=mux.Vars(r)
	isFound:=false 
	for _, user:=range users {
		if user.Id == params["id"] {
			isFound = true
			su:=json.NewEncoder(w)
			su.SetIndent("","  ")
			err:=su.Encode(user)
			if (err!=nil) {
				panic(err)
				return
			}
		}
	}
	if (isFound==false) {
		errorHandler(w,r,http.StatusNotFound,idNotFoundTrouble)
	}
}




//Функция добавления пользователя
func addUser(w http.ResponseWriter, r *http.Request) {
	id:=r.URL.Path[len("/adduser/"):]
	for _,user := range users {
		if user.Id == id {
			errorHandler(w,r,0,userAlreadyExists)
			return
		}
	}
	t,err:=template.ParseFiles("local/addUser.html")
	if (err!=nil) {
		panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.Execute(w,id)
	//w.WriteHeader(http.StatusCreated)

}

func saveNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id:=r.URL.Path[len("/save/"):]
		r.ParseForm()
		Username := strings.Join(r.Form["Username"], " ")
		OS_family := strings.Join(r.Form["OS_family"], " ")
		OS := strings.Join(r.Form["OS"], " ")
		Shell := strings.Join(r.Form["Shell"], " ")
		Kernel := strings.Join(r.Form["Kernel"], " ")
		Terminal := strings.Join(r.Form["Terminal"], " ")
		CPU := strings.Join(r.Form["CPU"], " ")
		users = append(users, User{Id: id,Username:Username, OS_family:OS_family, OS:OS, Shell:Shell,
		Kernel:Kernel, CPU:CPU, Terminal:Terminal})
		http.Redirect(w,r, "/users", http.StatusFound)
	} else {
		return
	}
}


//Сохранение юзера после апдейта
func saveUpdatedUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id:=r.URL.Path[len("/saveUpdatedUser/"):]
		r.ParseForm()
		for i:=0; i<len(users); i++ {
			user:= &users[i]
			if user.Id == id {
				Username := strings.Join(r.Form["Username"], " ")
				OS_family := strings.Join(r.Form["OS_family"], " ")
				OS := strings.Join(r.Form["OS"], " ")
				Shell := strings.Join(r.Form["Shell"], " ")
				Kernel := strings.Join(r.Form["Kernel"], " ")
				Terminal := strings.Join(r.Form["Terminal"], " ")
				CPU := strings.Join(r.Form["CPU"], " ")
				checkAndSave(user, "Username",  Username)
				checkAndSave(user, "OS_family",  OS_family)
				checkAndSave(user, "OS",OS)
				checkAndSave(user, "Shell",Shell)
				checkAndSave(user, "Kernel",Kernel)
				checkAndSave(user, "CPU",CPU)
				checkAndSave(user, "Terminal",Terminal)
				http.Redirect(w,r, "/users", http.StatusFound)
				return
			}
		}
	}
}

func checkAndSave(user *User, param string, value string) {
	if len(value) > 0 {
		switch param {
		case "Username":
			user.Username = value
		case "OS_family":
			user.OS_family = value
		case "OS":
			user.OS = value
		case "Shell":
			user.Shell = value
		case "Kernel":
			user.Kernel = value
		case "Terminal":
			user.Terminal = value
		case "CPU":
			user.CPU = value
		}
	} else {
		return
	}
}


func updateUser(w http.ResponseWriter, r * http.Request) {
	
	id:=r.URL.Path[len("/update/"):]
	params:=mux.Vars(r)
	isInList:=false
	for _, users := range users {
		
		if users.Id == params["id"] {
			isInList = true
			break
		}
	}
	if isInList==false {
		w.WriteHeader(404)
		errorHandler(w,r,http.StatusNotFound,userNotExists)
		return
	} else {
		t,err:=template.ParseFiles("local/updateUser.html")
		if (err!=nil) {
			panic(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		t.Execute(w,id)
	}

}

//Функция удаления пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	
	params:=mux.Vars(r)
	for index,user:=range users {
		if user.Id == params["id"] {
			users = append(users[:index], users[index+1:]...)
			http.Redirect(w,r, "/users", http.StatusFound)
			return
		}
	}

	w.WriteHeader(404)
	errorHandler(w, r,http.StatusNotFound,userNotExists)
	

}

//Обработчик ошибок
func errorHandler(w http.ResponseWriter, r *http.Request, status int, trouble int) {
	switch trouble {
		case idNotFoundTrouble:
			id:=r.URL.Path[len("/users/"):]
			w.WriteHeader(status)
			if status==http.StatusNotFound {
				t,err:=template.ParseFiles("local/auError.html")
				if (err!=nil) {
					panic(err)
					return
				}
				t.Execute(w,id)
				}
		case userAlreadyExists:
			id:=r.URL.Path[len("/adduser/"):]
			t,err:=template.ParseFiles("local/existstError.html")
			if (err!=nil) {
				panic(err)
				return
			}
		
			t.Execute(w,id)
		case userNotExists:
			id:=r.URL.Path[len("/adduser/"):]
			t,err:=template.ParseFiles("local/notExistsError.html")
			if (err!=nil) {
				panic(err)
				return
			}
			
			t.Execute(w,id)
	}
}

// beaty json
func encodeFullList(w http.ResponseWriter) {
	var linuxUsers []User
	for _,user:= range users {
		if strings.ToLower(user.OS_family) == "linux" {
			linuxUsers = append(linuxUsers,user)
		}
	}
	su:=json.NewEncoder(w)
	su.SetIndent("","  ")
	err:=su.Encode(linuxUsers)
	if(err!=nil) {
		panic(err)
		return
	}
}

func appendCustomUsers() {
	users = append(users, User{Id: "1",Username:"admin", OS_family:"Linux", OS:"Ubuntu", Shell:"bash 4.3.46",
	 Kernel:"4.4.0-34-generic", CPU:"Intel Core i5 @2.4GHz", Terminal: "gnome-terminal"})
}


func Handlers() *mux.Router {
	myRouter:=mux.NewRouter()
	sub := myRouter.PathPrefix("/").Subrouter()
	sub.Methods("GET").Path("/").HandlerFunc(homepage)
	sub.Methods("GET").Path("/users").HandlerFunc(showUsers)
	sub.Methods("GET").Path("/users/{id}").HandlerFunc(getUser)
	sub.Methods("PUT").Path("/update/{id}").HandlerFunc(updateUser)
	sub.Methods("POST").Path("/adduser/{id}").HandlerFunc(addUser)
	sub.Methods("DELETE").Path("/delete/{id}").HandlerFunc(deleteUser)
	myRouter.HandleFunc("/update/{id}", updateUser)
	myRouter.HandleFunc("/adduser/{id}", addUser)
	myRouter.HandleFunc("/save/{id}", saveNewUser)
	myRouter.HandleFunc("/delete/{id}", deleteUser)
	myRouter.HandleFunc("/saveUpdatedUser/{id}",saveUpdatedUser)
	return myRouter

}

	
