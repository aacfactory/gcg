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
	"strconv"
	"strings"
)

func (s *Statement) Literal(v interface{}) *Statement {
	if v == nil {
		s.Add(newToken("nil"))
		return s
	}
	out := ""
	switch v.(type) {
	case string:
		x := v.(string)
		if strconv.CanBackquote(x) {
			if strings.Contains(x, "\\") {
				out = `"` + x + `"`
			} else {
				out = "`" + x + "`"
			}
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
	s.Add(newToken(out))
	return s
}

func Literal(v interface{}) (code Code) {
	code = newStatement().Literal(v)
	return
}

func (s *Statement) LiteralByte(v byte) *Statement {
	s.Add(newToken(fmt.Sprintf("byte('%s')", string(v))))
	return s
}

func LiteralByte(v byte) (code Code) {
	code = newStatement().LiteralByte(v)
	return
}

func (s *Statement) LiteralRune(v rune) *Statement {
	s.Add(newToken(fmt.Sprintf(strconv.QuoteRune(v))))
	return s
}

func LiteralRune(v rune) (code Code) {
	code = newStatement().LiteralRune(v)
	return
}
