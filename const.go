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

func (s *Statement) Constant(ident string, lit interface{}) *Statement {
	s.Keyword("const").Space().Ident(ident).Space().Symbol("=").Space().Literal(lit).Line()
	return s
}

func Constant(ident string, lit interface{}) (stmt *Statement) {
	stmt = newStatement().Constant(ident, lit)
	return
}

func Constants() (c *ConstantsBuilder) {
	c = &ConstantsBuilder{
		group: Group("(", ")", "\n"),
	}
	return
}

type ConstantsBuilder struct {
	group *StatementGroup
}

func (c *ConstantsBuilder) Add(ident string, lit interface{}) {
	if lit == nil {
		c.group.Add(Ident(ident))
	} else {
		c.group.Add(Ident(ident).Space().Equal().Space().Literal(lit))
	}
	return
}

func (c *ConstantsBuilder) Build() (code Code) {
	stmt := newStatement()
	stmt.Keyword("const").Space()
	stmt.Add(c.group).Line()
	code = stmt
	return
}
