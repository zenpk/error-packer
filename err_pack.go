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
	Code int32
	Msg  string
}

// Error implement error type
func (e ErrPack) Error() string {
	return fmt.Sprintf("error: code=%d, Msg=%s", e.Code, e.Msg)
}

// Error Code
// A B CD
// A: error side [2: success, 4: client, 5: server]
// B: error position (layer)
// 		[0: client]
// 		[1: handler, 2: middleware, 3: rpc, 4: message queue, 5: service, 6: cache, 7: database, 9: unknown]
// CD:  [00: -, 01~99: specified error]

var (
	ErrOK               = ErrPack{2000, "success"}
	ErrUnknown          = ErrPack{5900, "unknown error"}
	ErrNotLogin         = ErrPack{4001, "user not logged in error"}
	ErrNoPermission     = ErrPack{4002, "no permission error"}
	ErrInputHeader      = ErrPack{4003, "input header error"}
	ErrInputBody        = ErrPack{4004, "input body error"}
	ErrParseToken       = ErrPack{5201, "parse token error"}
	ErrServiceConn      = ErrPack{5301, "service communication error"}
	ErrLogic            = ErrPack{5501, "logical error"}
	ErrCacheConn        = ErrPack{5601, "cache connection error"}
	ErrNoCache          = ErrPack{5602, "no cache error"}
	ErrDBConn           = ErrPack{5701, "database connection error"}
	ErrNoRecord         = ErrPack{5702, "database no record error"}
	ErrDuplicatedRecord = ErrPack{5703, "database duplicated record error"}
	ErrTypeConv         = ErrPack{5901, "type conversion error"}
	ErrGenJWT           = ErrPack{5903, "generate JWT error"}
	ErrParseJWT         = ErrPack{5903, "parse JWT error"}
)
