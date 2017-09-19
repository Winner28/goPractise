package muXhttp


import (
	"html/template"
	"encoding/json"
	"net/http"
	"models"
	"fmt"
)


type News struct {
	Content []string
}


type State struct {
	IsGet bool
	IsPost bool
}


var user = models.CreateUser()


var posts map[int]*Post
var postId int

func init() {
	posts = make(map[int]*Post, 0)
	postId = 0
}



//Get All Users JSON
func (env *Env) allUsersHandler(w http.ResponseWriter, r *http.Request) {
	users:=env.db.GetAllUsers()
	su:=json.NewEncoder(w)
	su.SetIndent("","  ")	
	err:=su.Encode(users)
	if(err!=nil) {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Println(posts[1])
	fmt.Println(chat.users)

}
//////////////////////////////////////////////////////////////////////////

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if len(user.Username) <=0  {
		 http.Redirect(w, r,"/login" , 302)
	} else {
		t, err := template.ParseFiles("templates/mainpage.html")
		if err!=nil {
			fmt.Fprintf(w, "Some error occured, please reload page :/")
			}
		t.Execute(w,user)
	}
	
}


func newsHandler(w http.ResponseWriter, r *http.Request) {
	if len(user.Username) <=0  {
		 http.Redirect(w, r,"/login" , 302)
	} else {
		t, err := template.ParseFiles("templates/news.html")
		if err!=nil {
			fmt.Fprintf(w, "Some error occured, please reload page :/")
			return
			} 
		 news := &News {
			Content :[]string{"vk.com", "facebook.com", "abc.com"},
		}

		t.Execute(w,news)
	}
	
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
	  state:= &State{}
	  if len(user.Username) <=0  {
		 http.Redirect(w, r,"/login" , 302)
		 return
		}
	if r.Method == "GET" {
			if len(user.Username) <=0  {
		 http.Redirect(w, r,"/login" , 302)
		 } else {
		 	t, err := template.ParseFiles("templates/newpost.html")
		 	if err!=nil {
			fmt.Fprintf(w, "Some error occured, please reload page :/")
			return
			}
			state.IsGet = true
			state.IsPost = false
			t.Execute(w, state)

		 }
		} else {
			title := r.FormValue("title")
			content := r.FormValue("subject")
			t, err := template.ParseFiles("templates/newpost.html")
		 	if err!=nil {
			fmt.Fprintf(w, "Some error occured, please reload page :/")
			return
			}
			posts[postId] = NewPost(postId, title, content)
			postId++
			state.IsGet = false
			state.IsPost = true
			t.Execute(w, state)
		}
}


func postsHandler(w http.ResponseWriter, r *http.Request) {
	if len(user.Username) <=0  {
		 http.Redirect(w, r,"/login" , 302)
		 return
		} else {
			t, err := template.ParseFiles("templates/posts.html")
			if err!=nil {
				fmt.Println(err.Error())
				return
			}
			t.Execute(w, posts)
		}
}




//////////////////////////////////////////////////////////////////////////
func (env *Env) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	t, err := template.ParseFiles("templates/login.html")
	if err!=nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println(username, password)
		isUserExists := env.db.CheckUser(username)
		if !isUserExists {
			fmt.Fprintf(w, "Register, please!")
		} else {
			isEnterCorrect, cuser := env.db.Login(username, password)
			if !isEnterCorrect {
				fmt.Fprintf(w, "<h1> Password incorrect, try again! </h1>")
			} else {
				user = cuser
				http.Redirect(w, r,"/home" , http.StatusFound)
			}
		}
	}

}

func (env *Env) registerHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/register.html")
		if err!=nil {
			fmt.Fprintf(w, err.Error())
		}
		t.Execute(w, nil)
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		mail := r.FormValue("mail")
		isUserExists := env.db.CheckUser(username)
		if isUserExists {
			fmt.Fprintf(w,"User is already register")
		} else {
			isSuccess := env.db.RegisterUser(username, password, mail)
			if isSuccess {
				fmt.Fprintf(w, "Register success!")
			} else {
				fmt.Fprintf(w, "Register is not success")
			}
		}
	}
}



func chatHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/chat.html")
	if err!=nil {
		fmt.Println(err.Error())
	}

	t.Execute(w,nil)
}