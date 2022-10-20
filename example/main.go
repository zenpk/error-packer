package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ep "github.com/zenpk/error-packer"
	"net/http"
	"net/http/httptest"
)

type User struct {
	Name string `json:"name,omitempty"`
}

type UserLoginResp struct {
	Seq  int64  `json:"seq" ep:"-1"`          // will be -1
	Code int64  `json:"code" ep:"err.code"`   // will eventually be errPack.Code
	Msg  string `json:"msg" ep:"err.msg"`     // will eventually be errPack.Msg
	User *User  `json:"user,omitempty" ep:""` // will omit this field
}

func handler(c *gin.Context) {
	// create a new packer your JSON interface
	packer := ep.Packer{V: UserLoginResp{}}
	err := ep.ErrInputBody      // assume that an input body error happened
	resp := packer.Pack(err)    // pack the response interface with the error
	c.JSON(http.StatusOK, resp) // return with resp
}

func main() {
	req, _ := http.NewRequest(http.MethodGet, "/err", nil) // make a mock request
	rec := httptest.NewRecorder()                          // record the mock request
	// use Gin to handle the request
	r := gin.Default()
	r.GET("/err", handler)
	r.ServeHTTP(rec, req)
	fmt.Println(rec.Body.String())
}
