package muXhttp
import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	

)

type ChatRoom struct  {
	users map[string]User
	queue chan string
}


type User struct {
	Username string
	conn *websocket.Conn
	fromRoom *ChatRoom

}

var upgrader = websocket.Upgrader {
	ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func (cr *ChatRoom) Init() {
	cr.users = make(map[string]User)
	cr.queue = make(chan string,5)

	go func() {
		for {
			cr.monitorMessages()
			time.Sleep(100*time.Millisecond)
		}
	}()
}

func (cr *ChatRoom) monitorMessages() {
	message:=""
	infLoop:
		for {
			select {
			 case m:= <- cr.queue:
			 	message+= m
			default:
				break infLoop
			}
		}

		if len(message) >0 {
			for _,user:= range cr.users {
				user.Send(message)
			}
		}
}

func (user *User) Send(message string) {
	user.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (cr *ChatRoom) Join(name string, conn *websocket.Conn) *User {
	newUser := User {
		Username: name,
		conn: conn,
		fromRoom: cr,
	}
	if _,exists:=cr.users[name]; exists {
		return nil
	}
	cr.users[name] = newUser
	cr.AddMsg("<B>"+ name + " just joined the chat..." +"</B>")
	return &newUser
}

func (user *User) Leave() {
	user.fromRoom.Exit(user.Username)
}

func (cr *ChatRoom) Exit(name string) {
	delete(cr.users, name)
	cr.AddMsg(name + " has left!")

}

func (cr *ChatRoom) AddMsg(msg string) {
	cr.queue <- msg
}

func (user *User) NewMsg(msg string) {
	user.fromRoom.AddMsg(user.Username + ": " + msg)
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w,r,nil)
	if err!=nil {
		conn.Close()
		return
	}

	go func() {
		nUser:= chat.Join(user.Username, conn)
		if nUser==nil {
			conn.Close()
			return
		}

		for {
			_,msg,err := conn.ReadMessage()
			if err!=nil {
				conn.Close()
				nUser.Leave()
				return
			}
			nUser.NewMsg(string(msg))
		}
	}()
}	