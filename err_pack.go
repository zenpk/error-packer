// Copyright 2022 zenpk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ep

import "fmt"

type ErrPack struct {
	Code int16
	Msg  string
}

// Error implement error type
func (e ErrPack) Error() string {
	return fmt.Sprintf("error: code=%d, Msg=%s", e.Code, e.Msg)
}

// Error Code
// A B C D
// A:  [2: success, 4: client error, 5: server error]
// B:  [0: -, 1: input error, 2: data type error, 3: database error, 4: communication error, 5: logical error, 6: unknown error]
// CD: [00: -, 01~99: specified error]

var (
	ErrOK              = ErrPack{2000, "success"}
	ErrUnknown         = ErrPack{5600, "unknown error"}
	ErrInputHeader     = ErrPack{4101, "input header error"}
	ErrInputBody       = ErrPack{4102, "input body error"}
	ErrInputToken      = ErrPack{4103, "input token error"}
	ErrInputCookie     = ErrPack{4104, "input cookie error"}
	ErrTypeConv        = ErrPack{5201, "type conversion error"}
	ErrParseToken      = ErrPack{5202, "parse token error"}
	ErrParseCookie     = ErrPack{5203, "parse token error"}
	ErrSetToken        = ErrPack{5204, "set token error"}
	ErrSetCookie       = ErrPack{5205, "set cookie error"}
	ErrDBConn          = ErrPack{5301, "database connection error"}
	ErrNoRecord        = ErrPack{5302, "database no record error"}
	ErrDuplicateRecord = ErrPack{5303, "database duplicate record error"}
	ErrServiceConn     = ErrPack{5401, "service communication error"}
)
