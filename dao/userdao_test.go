package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	t.Run("验证用户名字和年龄", testLogin)
	t.Run("验证用户名字", testRegist)
	// t.Run("插入用户名字和年龄", testSave)
}

func testLogin(t *testing.T) {
	fmt.Println(CheckUserNameAndPassword("admin", 1))
	// CheckUserNameAndPassword("admin", 1)
}

func testRegist(t *testing.T) {
	fmt.Println(CheckUserName("admin"))
	// CheckUserName("admin")
}

// func testSave(t *testing.T) {
// 	fmt.Println(SaveUser("admin1", 2, "男"))
// 	// SaveUser("admin1", 2, "男")
// }
