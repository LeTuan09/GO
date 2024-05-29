# GO

`gin.Context` cung cấp các phương thức để lấy dữ liệu từ yêu cầu HTTP như tham số, query string, body và headers.

```go```
func handler(c *gin.Context) {
    name := c.Query("name")          // Lấy query parameter 'name'
    age := c.Param("age")            // Lấy URL parameter 'age'
    body := c.PostForm("body")       // Lấy form data 'body'
    header := c.GetHeader("token")   // Lấy giá trị của header 'token'

    c.JSON(200, gin.H{
        "name": name,
        "age": age,
        "body": body,
        "header": header,
    })
}
