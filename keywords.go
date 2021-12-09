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

func (s *Statement) Keyword(ident string) *Statement {
	if ident == "default" {
		ident = ident + ":"
	}
	s.Add(newToken(ident))
	return s
}

func Keyword(ident string) (code Code) {
	code = newStatement().Keyword(ident)
	return
}

func (s *Statement) Break() *Statement {
	s.Keyword("break")
	return s
}

func Break() (code Code) {
	code = newStatement().Break()
	return
}

func (s *Statement) Default() *Statement {
	s.Keyword("default")
	return s
}

func Default() (code Code) {
	code = newStatement().Default()
	return
}

func (s *Statement) Func() *Statement {
	s.Keyword("func")
	return s
}

func Func() (code Code) {
	code = newStatement().Func()
	return
}

func (s *Statement) Interface() *Statement {
	s.Keyword("interface")
	return s
}

func Interface() (code Code) {
	code = newStatement().Interface()
	return
}

func (s *Statement) Select() *Statement {
	s.Keyword("select")
	return s
}

func Select() (code Code) {
	code = newStatement().Select()
	return
}

func (s *Statement) Case() *Statement {
	s.Keyword("case")
	return s
}

func Case() (code Code) {
	code = newStatement().Case()
	return
}

func (s *Statement) Defer() *Statement {
	s.Keyword("defer")
	return s
}

func Defer() (code Code) {
	code = newStatement().Defer()
	return
}

func (s *Statement) Go() *Statement {
	s.Keyword("go")
	return s
}

func Go() (code Code) {
	code = newStatement().Go()
	return
}

func (s *Statement) Map() *Statement {
	s.Keyword("map")
	return s
}

func Map() (code Code) {
	code = newStatement().Map()
	return
}

func (s *Statement) Chan() *Statement {
	s.Keyword("chan")
	return s
}

func Chan() (code Code) {
	code = newStatement().Chan()
	return
}

func (s *Statement) Else() *Statement {
	s.Keyword("else")
	return s
}

func Else() (code Code) {
	code = newStatement().Else()
	return
}

func (s *Statement) Goto() *Statement {
	s.Keyword("goto")
	return s
}

func Goto() (code Code) {
	code = newStatement().Goto()
	return
}

func (s *Statement) Switch() *Statement {
	s.Keyword("switch")
	return s
}

func Switch() (code Code) {
	code = newStatement().Switch()
	return
}

func (s *Statement) Fallthrough() *Statement {
	s.Keyword("fallthrough")
	return s
}

func Fallthrough() (code Code) {
	code = newStatement().Fallthrough()
	return
}

func (s *Statement) If() *Statement {
	s.Keyword("if")
	return s
}

func If() (code Code) {
	code = newStatement().If()
	return
}

func (s *Statement) Range() *Statement {
	s.Keyword("range")
	return s
}

func Range() (code Code) {
	code = newStatement().Range()
	return
}

func (s *Statement) Continue() *Statement {
	s.Keyword("continue")
	return s
}

func Continue() (code Code) {
	code = newStatement().Continue()
	return
}

func (s *Statement) For() *Statement {
	s.Keyword("for")
	return s
}

func For() (code Code) {
	code = newStatement().For()
	return
}

func (s *Statement) Return() *Statement {
	s.Keyword("return")
	return s
}

func Return() (code Code) {
	code = newStatement().Return()
	return
}
