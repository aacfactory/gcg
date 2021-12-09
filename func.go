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

import "strings"

func (s *Statement) Call(fnName string, params ...Code) *Statement {
	fnName = strings.TrimSpace(fnName)
	if fnName != "" {
		s.Token(fnName)
	}
	s.Token("(")
	if params != nil {
		for i, param := range params {
			s.Add(param)
			if i > 0 && i < len(params)-1 {
				s.Symbol(",").Space()
			}
		}
	}
	s.Token(")")
	return s
}

func Call(fnName string, params ...Code) (stmt *Statement) {
	stmt = newStatement().Call(fnName, params...)
	return
}
