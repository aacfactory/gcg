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
	"sort"
)

func (s *Statement) Constant(ident string, lit interface{}) *Statement {
	s.Keyword("const").Space().Ident(ident).Space().Symbol("=").Space().Literal(lit).Line()
	return s
}

func Constant(ident string, lit interface{}) (stmt *Statement) {
	stmt = newStatement().Constant(ident, lit)
	return
}

func Constants(items map[string]interface{}) (stmt *Statement) {
	stmt = newStatement()
	stmt.Keyword("const").Space()
	group := Group("(", ")", "\n")
	idents := make([]string, 0, 1)
	for ident := range items {
		if ident == "" {
			continue
		}
		idents = append(idents, ident)
	}
	sort.Strings(idents)
	for _, ident := range idents {
		lit := items[ident]
		group.Add(Ident(ident).Space().Equal().Space().Literal(lit))
	}
	stmt.Add(group).Line()
	return
}
