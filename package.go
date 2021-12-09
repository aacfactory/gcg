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
	"path"
	"strings"
)

var emptyPackages = NewPackages()

func NewPackages() Packages {
	return make([]*Package, 0, 1)
}

type Packages []*Package

func (p Packages) Empty() (v bool) {
	v = len(p) == 0
	return
}

func (p *Packages) Add(v *Package) *Packages {
	added := false
	for _, import0 := range *p {
		if import0.Path == v.Path {
			added = true
			break
		}
	}
	if added {
		return p
	}
	*p = append(*p, v)
	return p
}

func (p *Packages) Merge(x Packages) {
	if x == nil || len(x) == 0 {
		return
	}
	for _, i := range x {
		p.Add(i)
	}
}

func (p Packages) MapToCode() (code Code) {
	stmt := newStatement()
	if len(p) == 0 {
		code = stmt
		return
	}

	group := Group("(", ")", "\n")
	for _, i := range p {
		if i.Alias != "" {
			group.Add(Ident(i.Alias).Space().Symbol("\"").Token(i.Path).Symbol("\""))
		} else {
			group.Add(Symbol("\"").Token(i.Path).Symbol("\""))
		}
	}
	stmt.Keyword("import").Space().Add(group).Line()
	code = stmt
	return
}

func NewPackage(pkgPath string) (v *Package) {
	v = NewPackageWithAlias(pkgPath, "")
	return
}

func NewPackageWithAlias(pkgPath string, alias string) (v *Package) {
	name := ""
	alias = strings.TrimSpace(alias)
	if alias != "" {
		name = alias
	} else {
		_, name = path.Split(pkgPath)
	}
	v = &Package{
		Name:  name,
		Alias: alias,
		Path:  pkgPath,
	}
	return
}

type Package struct {
	Name  string
	Alias string
	Path  string
}
