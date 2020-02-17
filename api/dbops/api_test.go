package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/Chenxi97/GiliGili/api/defs"
	"github.com/jinzhu/gorm"
)

var (
	tempvid string
	tempaid uint
)

func clearTables() {
	dbConn.Delete(&defs.User{})
	dbConn.Delete(&defs.VideoInfo{})
	dbConn.Delete(&defs.Comment{})
	dbConn.Delete(&defs.Session{})
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUser(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	user, err := AddUser("root", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
	tempaid = user.ID
}

func testGetUser(t *testing.T) {
	user, err := GetUser("root")
	if user.Pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("root", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	user, err := GetUser("root")
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if user != nil && user.Pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func TestVideo(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.ID
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddCommnets", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := tempaid
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
