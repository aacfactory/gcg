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

func (s *Statement) Ident(name string) *Statement {
	s.Token(name)
	return s
}

func Ident(name string) (stmt *Statement) {
	stmt = newStatement().Ident(name)
	return
}

func (s *Statement) QualifiedIdent(pkg *Package, name string) *Statement {
	s.Token(pkg.Name).Dot().Token(name, pkg)
	return s
}

func QualifiedIdent(pkg *Package, name string) (stmt *Statement) {
	stmt = newStatement().QualifiedIdent(pkg, name)
	return
}
