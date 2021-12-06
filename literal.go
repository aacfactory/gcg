/*
 * Copyright 2021 Wang Min Xiang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gcg

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Literal(v interface{}) (code Code) {
	if v == nil {
		code = &literalCode{
			v: "nil",
		}
		return
	}
	out := ""
	switch v.(type) {
	case string:
		x := v.(string)
		if strconv.CanBackquote(x) {
			out = "`" + x + "`"
		} else {
			out = `"` + x + `"`
		}
	case bool, int, complex128:
		out = fmt.Sprintf("%#v", v)
	case float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		out = fmt.Sprintf("%T(%#v)", v, v)
	case float64:
		out = fmt.Sprintf("%#v", v)
		if !strings.Contains(out, ".") && !strings.Contains(out, "e") {
			out += ".0"
		}
	case complex64:
		out = fmt.Sprintf("%T%#v", v, v)
	case []byte:
		x := string(v.([]byte))
		if strconv.CanBackquote(x) {
			out = "[]byte(" + "`" + x + "`" + ")"
		} else {
			out = "[]byte(" + `"` + x + `"` + ")"
		}
	default:
		panic(fmt.Sprintf("gcg: unsupported type for literal: %T", v))
	}
	code = &literalCode{
		v: out,
	}
	return
}

func LiteralByte(v byte) (code Code) {
	code = &literalCode{
		v: fmt.Sprintf("byte('%s')", string(v)),
	}
	return
}

type literalCode struct {
	v string
}

func (l literalCode) Render(w io.Writer) (err error) {
	_, err = w.Write([]byte(l.v))
	return
}

func (l literalCode) imports() (imports Imports) {
	imports = emptyImports
	return
}
