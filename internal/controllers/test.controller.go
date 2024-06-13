package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Test struct {
	age int
	name string
}

// TODO: Khởi tạo hàm có kiểu là con trỏ trỏ tới struct và trả về con trỏ trỏ tới đối tượng Test
func NewTest () *Test {
	return &Test{}
}


func (t *Test) Testc(c *gin.Context) {
	name := c.DefaultQuery("TuanZin", "Test")
	// m := "TuanZin" 

	// TODO: Khởi tạo con trỏ p có kiểu dữ liệu chính là kiểu dữ liệu của biến m đồng thời gán giá trị con trỏ p là địa chỉ của biến m, để lấy giá trị của con trỏ p dùng *p => giá trị của biến m
	// p := &m 
	
	// TODO: Khởi tạo con trỏ k có kiểu float32 nhưng chưa gán giá trị thì mặc định nó là nil
	// var k *float32
	// TODO: Khởi tạo con trỏ x có kiểu Test, nhưng chưa gán giá trị thì mặc định nó là nil
	var x *Test
	// Gán địa chỉ của một đối tượng Test mới cho p
	x = &Test{10, "tunazin"} // println: &{value}

	// fmt.Println(x)
	// fmt.Println(p)
	fmt.Println("Giá trị của ", x)
	fmt.Println("Giá trị của ", *x)


	c.JSON(200, gin.H{
		"name": name,
	})
}
/**
 * !controller -> service -> repo -> model -> dbs
 */