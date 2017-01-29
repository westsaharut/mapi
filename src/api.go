package main

// รวมๆหน้านี้เรียกว่า Service

import (
  "net/http"

  "github.com/labstack/echo"

  "github.com/labstack/echo/middleware"
)
// หรือ
// import "net/http"
// import "github.com/labstack/echo"

type User struct {
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Username string `json:"username,omitempty"`  // omitempty คือถ้าไม่มีข้อมูลใน value จะไม่แสดง key ออกมา
  Password string `json:"password,omitempty"`  // json:"password" จะตั้งเป็นชื่ออะไรก็ได้ แล้วจะไปแสดงเป็นแบบนั้น
}

func index(c echo.Context) error {
  return c.JSON(http.StatusOK, "Hello World")
}

func getUserByID(c echo.Context) error {
  id := c.Param("id") // id := c.Param("id") ชื่อตรง id ต้องเหมือนกัน e.GET("/users/:id", getUserByID)
  return c.JSON(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
  user := new(User)
  if err := c.Bind(user); err != nil{
    return c.JSON(http.StatusBadRequest, nil)
  }
  return c.JSON(http.StatusOK, user)
}

func getUsers(c echo.Context) error {
  west := User{
    Firstname:  "saharut",
    Lastname: "suntiwarawit",
    Username: "west",
    // "1234",
    // ต้องมี , ตัวสุดท้าย เพราะ go จะเติม ; ให้เอง เมื่อเติมตรงนี้จะเกิด error เลยต้องใส่ ,
  }
  // หรือ var firstname string = "West"
  // หรือ var firstname = "West"
  return c.JSON(http.StatusOK, west)
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.GET("/", index) // / เป็น default
  e.GET("/users", getUsers) // 1 / = 1 Service
  e.GET("/users/:id", getUserByID)
  e.POST("/users", saveUser)
  e.Logger.Fatal(e.Start(":1323"))
}
