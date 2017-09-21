package api_test

import (
	"api"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	currentUser int
	allUsers int
	deleteUrl string
	updateUrl string
	usersUrl string
	userUrl string
	addUrl string

)

func init() {
	currentUser = 1
	allUsers = 1
	server = httptest.NewServer(api.Handlers())
	fmt.Println(server.URL)
	usersUrl = fmt.Sprintf("%s/users", server.URL)
	userUrl = fmt.Sprintf("%s/users/1", server.URL)
	updateUrl = fmt.Sprintf("%s/update/1", server.URL)
	fmt.Println(updateUrl)
	addUrl = fmt.Sprintf("%s/adduser/%v", server.URL, currentUser)
	deleteUrl = fmt.Sprintf("%s/delete/%v", server.URL, currentUser)
}


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






//Test Юзера
func TestShowUsers(t *testing.T) {
	fmt.Println("---Testing Show all users--- ")

	reader = strings.NewReader("")
	request, err := http.NewRequest("GET", usersUrl, reader)

	res, err:= http.DefaultClient.Do(request)

	if err!=nil {
		t.Error(err)
	}

	checkResponseCode(t, http.StatusOK, res.StatusCode)
}


//Тест всего списка
func TestShowUser(t *testing.T) {
	fmt.Println("---Testing Show user---")

	reader = strings.NewReader("")
	request,err := http.NewRequest("GET", userUrl,reader)

	if err!=nil {
		t.Error(err)
	}

	res, err:= http.DefaultClient.Do(request)

	checkResponseCode(t, http.StatusOK, res.StatusCode)

}

//Test апдейта
func TestUpdateUser(t *testing.T) {


	userJson:= `{"Id":"1", "Username":"admin", "OS_family":"Windows", 
	"OS":"OS", "Shell":"?!","Kernel:Kernel", "CPU":"Wow!","Terminal" : "Terminal"}`

	reader = strings.NewReader(userJson)

	request, err :=  http.NewRequest("PUT", updateUrl, reader)

	res, err := http.DefaultClient.Do(request)

	fmt.Println("---Testing User Update---")
	if err!=nil {
		t.Error(err)
	}
	checkResponseCode(t, http.StatusOK, res.StatusCode)

}

//Тест добавления
func TestAddUser(t *testing.T) {
	
	userJson:= `{"Id":"1", "Username":"User", "OS_family":"OS_family", 
	"OS":"OS", "Shell":"Shell","Kernel:Kernel", "CPU":"CPU","Terminal" : "Terminal"}`

	reader = strings.NewReader(userJson)

	request, err :=  http.NewRequest("POST", addUrl, reader)

	res, err := http.DefaultClient.Do(request)

	fmt.Println("Testing Add a new User")
	if err!=nil {
		t.Error(err)
	}

	isAdd:=checkResponseCode(t, http.StatusOK, res.StatusCode)
	if isAdd!=true {
		t.Error("Mistake adding new user")
	} 
	
}

//Test удаления юзера, который находится в списке
func TestDeleteUser(t *testing.T) {
	fmt.Println("---Testing Delete---")

	reader = strings.NewReader("")
	request,err := http.NewRequest("DELETE", deleteUrl,reader)

	if err!=nil {
		t.Error(err)
	}

	res, err:= http.DefaultClient.Do(request)

	isDelete:=checkResponseCode(t, http.StatusOK, res.StatusCode)
	if isDelete!=true {
		t.Error("Delete error")
	} else {
		currentUser++;
	}
}

//Test апдейта
func TestUpdateUserThatNotExists(t *testing.T) {


	userJson:= `{"Id":"1", "Username":"admin", "OS_family":"Windows", 
	"OS":"OS", "Shell":"?!","Kernel:Kernel", "CPU":"Wow!","Terminal" : "Terminal"}`

	reader = strings.NewReader(userJson)

	request, err :=  http.NewRequest("PUT", updateUrl, reader)

	res, err := http.DefaultClient.Do(request)

	fmt.Println("---Testing User Update---")
	if err!=nil {
		t.Error(err)
	}

	if http.StatusOK != res.StatusCode {
			fmt.Println("      ---OK---     ")
			fmt.Println()
	}

}


//Test удаления юзера, которого не существует
func TestDeleteUserThatNotExists(t *testing.T) {
	fmt.Println("---Testing Delete (not exists)---")

	reader = strings.NewReader("")
	request,err := http.NewRequest("PUT", deleteUrl,reader)

	if err!=nil {
		t.Error(err)
	}

	res, err:= http.DefaultClient.Do(request)

	if http.StatusOK != res.StatusCode {
			fmt.Println("      ---OK---     ")
			fmt.Println()
	}
}




func checkResponseCode(t *testing.T, expected, actual int) bool{
	if expected!=actual {
		t.Error("Expected is not equals actual")
		return false
	} else {
		fmt.Println("      ---OK---     ")
		fmt.Println()
		return true
	}
}