# error-packer

Go struct packer, specifically for handling errors, works similarly as https://github.com/zenpk/gin-error-handler but more flexible.

## Why

The way Go handles errors will always end up with a bunch of `if err!=nil`, it's ok itself but when we have to return a
struct every time an error occurs, it can take up a lot of lines which significantly decreases the readability and make
typing work more tedious.

Take Gin framework as an example, the usual way to handle errors is something like this

```go
package whygolandforcesmetouseapackageinreadme

func handler(c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusOK, SomeStruct{
			Resp: Resp{
				Code: -1,
				Msg:  err.Error(),
			},
			Data: someData,
		})
		return
	}
}
```

After introducing the error-packer (ep), it will become something like this

```go
package whygolandforcesmetouseapackageinreadme

func handler(c *gin.Context) {
	packer := ep.Packer{V: SomeStruct{}}
	if err != nil {
		c.JSON(http.StatusOK, packer.Pack(err))
		return
	}
}
```

Much cleaner right? Let's get into it.

## What does it do?

### Packer

Packer is for "packing a struct". It will pack an ErrPack struct into any struct and fill the fields with tagged values.

### ErrPack struct

ErrPack struct is an error information struct which implements the `error` type, it contains a code field and a message
field. You can customize these fields as you want.

If you don't need a code field, you can also just pass the err (type error) into Packer, it accepts both types.

## Usage

Copy the `packer.go` and `err_pack.go` file to wherever you want.

Define your struct with "ep" tags.

There are 3 meaningful tags, others stand for the default values of the fields:

`ep:"err.code"` - the field with this tag will be filled with ErrPack.Code, it must be int type
`ep:"err.msg"` - the field with this tag will be filled with ErrPack.Msg, it must be string type
`ep:""` - the field with an empty tag will remain the default value of the type, you can also omit the tag

## Example

```go
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
	Code int64  `json:"code" ep:"err.code"`   // will eventually be ErrPack.Code
	Msg  string `json:"msg" ep:"err.msg"`     // will eventually be ErrPack.Msg
	User *User  `json:"user,omitempty" ep:""` // will omit this field
}

func handler(c *gin.Context) {
	// create a new packer your JSON interface
	packer := ep.Packer{V: UserLoginResp{}}
	err := ep.ErrInputBody      // assume that an input body error happened
	resp := packer.Pack(err)    // pack the response struct with the error
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
```

### Output

```text
{"seq":-1,"code":4102,"msg":"input body error"}
```
