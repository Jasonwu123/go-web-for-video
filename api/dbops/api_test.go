package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// Test user
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("jasonwu", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v\n", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("jasonwu")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("jasonwu", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v\n", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("jasonwu")
	if err != nil {
		t.Errorf("Error of RegetUser: %V\n", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed!")
	}
}

// Test video_info
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideoInfo", testAddNewVideo)
	t.Run("GetVideoInfo", testGetVideoInfo)
	t.Run("DeleteVideoInfo", testDeleteVideoInfo)
	t.Run("RegetVideoInfo", testRegetVideoInfo)
}

var tempvid string

func testAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "my_video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v\n", err)
	}
	tempvid = video.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVidoInfo: %v\n", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v\n", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempvid)
	if err != nil || video != nil {
		t.Errorf("Error of RegetVidoInfo: %v\n", err)
	}
}

// Test comments API
func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddNewComments)
	t.Run("ListComments", testListComments)
}

func testAddNewComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video!"

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddNewComments: %v\n", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v\n", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v\n", i, ele)
	}
}
