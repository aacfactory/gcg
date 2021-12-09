/*
 * Copyright 2021 Wang Min Xiang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package gcg

func (s *Statement) Symbol(v string) *Statement {
	s.Add(newToken(v))
	return s
}

func Symbol(s string) (stmt *Statement) {
	stmt = newStatement().Symbol(s)
	return
}

func (s *Statement) Space() *Statement {
	s.Add(newToken(" "))
	return s
}

func Space() (stmt *Statement) {
	stmt = newStatement().Space()
	return
}

func (s *Statement) Line() *Statement {
	s.Add(newToken("\n"))
	return s
}

func Line() (stmt *Statement) {
	stmt = newStatement().Line()
	return
}

func (s *Statement) Tab() *Statement {
	s.Add(newToken("\t"))
	return s
}

func Tab() (stmt *Statement) {
	stmt = newStatement().Tab()
	return
}

func (s *Statement) Tabs(n int) *Statement {
	for i := 0; i < n; i++ {
		s.Add(newToken("\t"))
	}
	return s
}

func Tabs(n int) (stmt *Statement) {
	stmt = newStatement().Tabs(n)
	return
}

func (s *Statement) Dot() *Statement {
	s.Symbol(".")
	return s
}

func Dot() (stmt *Statement) {
	stmt = newStatement().Dot()
	return
}

func (s *Statement) Equal() *Statement {
	s.Symbol("=")
	return s
}

func Equal() (stmt *Statement) {
	stmt = newStatement().Equal()
	return
}

func (s *Statement) ColonEqual() *Statement {
	s.Symbol(":=")
	return s
}

func ColonEqual() (stmt *Statement) {
	stmt = newStatement().ColonEqual()
	return
}

func (s *Statement) Colon() *Statement {
	s.Symbol(":")
	return s
}

func Colon() (stmt *Statement) {
	stmt = newStatement().Colon()
	return
}

func (s *Statement) LT() *Statement {
	s.Symbol("<")
	return s
}

func LT() (stmt *Statement) {
	stmt = newStatement().LT()
	return
}

func (s *Statement) LTE() *Statement {
	s.Symbol("<=")
	return s
}

func LTE() (stmt *Statement) {
	stmt = newStatement().LTE()
	return
}

func (s *Statement) GT() *Statement {
	s.Symbol(">")
	return s
}

func GT() (stmt *Statement) {
	stmt = newStatement().GT()
	return
}

func (s *Statement) GTE() *Statement {
	s.Symbol(">=")
	return s
}

func GTE() (stmt *Statement) {
	stmt = newStatement().GTE()
	return
}

func (s *Statement) Star() *Statement {
	s.Symbol("*")
	return s
}

func Star() (stmt *Statement) {
	stmt = newStatement().Star()
	return
}
