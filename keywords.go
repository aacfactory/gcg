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
	s.Token(ident)
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

func (s *Statement) Goto() *Statement {
	s.Keyword("goto")
	return s
}

func Goto() (code Code) {
	code = newStatement().Goto()
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

func (s *Statement) Continue() *Statement {
	s.Keyword("continue")
	return s
}

func Continue() (code Code) {
	code = newStatement().Continue()
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
