package muXhttp

import (
	"github.com/gorilla/mux"
	"net/http"
	"models"
	"fmt"
)


type Env struct {
	db models.DataOperations
}

var chat ChatRoom 

func InitAll() *mux.Router {
	fmt.Println("...Init DB...")
	dBname := "mysql"
	dBoptions := "root:password@tcp(127.0.0.1:port)/database" //a.k.a 3306
	db := models.InitDB(dBname, dBoptions)
	env := &Env{db:db}
	fmt.Println("...Init muXhttp...")
	chat.Init()
	myRouter:=mux.NewRouter()
 	myRouter.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/", myRouter)
	myRouter.HandleFunc("/login",env.loginHandler)
	myRouter.HandleFunc("/register", env.registerHandler)
	myRouter.HandleFunc("/AllUsers", env.allUsersHandler)
	myRouter.HandleFunc("/home", homeHandler)
	myRouter.HandleFunc("/news", newsHandler)
	myRouter.HandleFunc("/newpost", newPostHandler)
	myRouter.HandleFunc("/posts", postsHandler)
	myRouter.HandleFunc("/chat", chatHandler)
	myRouter.HandleFunc("/ws", socketHandler)
	return myRouter
}