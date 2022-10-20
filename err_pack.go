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
// ABC
// A:  [0: success,
//		(client) 1: input error,
//		(server) 2: data type error, 3: database error, 4: communication error, 5: logical error, 6: unknown error]
// BC: [00: -, 01~99: specified error]

var (
	ErrOK              = ErrPack{0, "success"}
	ErrUnknown         = ErrPack{600, "unknown error"}
	ErrInputHeader     = ErrPack{101, "input header error"}
	ErrInputBody       = ErrPack{102, "input body error"}
	ErrInputToken      = ErrPack{103, "input token error"}
	ErrInputCookie     = ErrPack{104, "input cookie error"}
	ErrTypeConv        = ErrPack{201, "type conversion error"}
	ErrParseToken      = ErrPack{202, "parse token error"}
	ErrParseCookie     = ErrPack{203, "parse token error"}
	ErrSetToken        = ErrPack{204, "set token error"}
	ErrSetCookie       = ErrPack{205, "set cookie error"}
	ErrDBConn          = ErrPack{301, "database connection error"}
	ErrNoRecord        = ErrPack{302, "database no record error"}
	ErrDuplicateRecord = ErrPack{303, "database duplicate record error"}
	ErrServiceConn     = ErrPack{401, "service communication error"}
)
