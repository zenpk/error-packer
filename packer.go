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

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

type Packer struct {
	V interface{}
}

// Pack return the given interface with tagged values and ErrPack
func (p *Packer) Pack(err error) interface{} {
	errPack := convertErrPack(err)
	return p.packCore(errPack)
}

// PackWithInfo will do the pack and additionally log the error with INFO level
func (p *Packer) PackWithInfo(err error) interface{} {
	errPack := convertErrPack(err)
	log.Printf("INFO: %s", errPack.Msg) // use the logger you prefer
	return p.packCore(errPack)
}

// PackWithWarn will do the pack and additionally log the error with WARN level
func (p *Packer) PackWithWarn(err error) interface{} {
	errPack := convertErrPack(err)
	log.Printf("WARN: %s", errPack.Msg) // use the logger you prefer
	return p.packCore(errPack)
}

// PackWithError will do the pack and additionally log the error with Error level
func (p *Packer) PackWithError(err error) interface{} {
	errPack := convertErrPack(err)
	log.Printf("ERROR: %s", errPack.Msg) // use the logger you prefer
	return p.packCore(errPack)
}

// convertErrPack convert error to ErrPack
func convertErrPack(err error) (errPack ErrPack) {
	if err == nil {
		errPack = ErrOK
	} else {
		if !errors.As(err, &errPack) {
			errPack = ErrUnknown
			errPack.Msg = err.Error()
		}
	}
	return
}

// packCore does the core functions
func (p *Packer) packCore(errPack ErrPack) interface{} {
	ref := reflect.ValueOf(&p.V).Elem()
	refCopy := reflect.New(ref.Elem().Type()).Elem()
	for i := 0; i < ref.Elem().NumField(); i++ {
		tag := ref.Elem().Type().Field(i).Tag.Get("ep")
		if tag == "-" {
			continue
		}
		// embedded struct, must be exported
		if ref.Elem().Field(i).Kind() == reflect.Struct {
			embeddedPack := &Packer{V: ref.Elem().Field(i).Interface()}
			refCopy.Field(i).Set(reflect.ValueOf(embeddedPack.packCore(errPack)))
			continue
		}
		if len(tag) <= 0 {
			continue
		}
		if tag == "err.msg" {
			refCopy.Field(i).SetString(errPack.Msg)
			continue
		}
		if tag == "err.code" {
			refCopy.Field(i).SetInt(int64(errPack.Code))
			continue
		}
		switch ref.Elem().Field(i).Interface().(type) {
		case string:
			refCopy.Field(i).SetString(tag)
		case int64, int32, int:
			val, _ := strconv.ParseInt(tag, 10, 64)
			refCopy.Field(i).SetInt(val)
		case uint64, uint32, uint:
			val, _ := strconv.ParseUint(tag, 10, 64)
			refCopy.Field(i).SetUint(val)
		case bool:
			val, _ := strconv.ParseBool(tag)
			refCopy.Field(i).SetBool(val)
		case float64, float32:
			val, _ := strconv.ParseFloat(tag, 64)
			refCopy.Field(i).SetFloat(val)
		default:
			// other implementations...
		}
	}
	ref.Set(refCopy)
	return p.V
}
