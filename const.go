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
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func NewConstant(name string, value TypeExpr) (v *Constant, err error) {
	name = strings.TrimSpace(name)
	if name == "" {
		err = fmt.Errorf("gcg: new contant failed for name is empty")
		return
	}
	if value == nil {
		err = fmt.Errorf("gcg: new contant failed for value is empty")
		return
	}
	code, codeErr := value.Type().ValueStatement().Code()
	if codeErr != nil {
		err = fmt.Errorf("gcg: new contant failed for %v", codeErr)
		return
	}
	cv := ""
	switch value.Type().Kind {
	case "string":
		cv = code
	case "bool":
		cv = code
	case "int", "int8", "int16", "int32", "int64":
		cv = code
	case "uint", "uint8", "uint16", "uint32", "uint64":
		cv = code
	case "float32", "float64":
		cv = code
	default:
		err = fmt.Errorf("gcg: new contant failed for type of value is not supported")
		return
	}
	v = &Constant{
		name:  name,
		value: cv,
	}
	return
}

type Constant struct {
	name  string
	value string
}

func (c Constant) Code() (code string, err error) {
	code = fmt.Sprintf("const %s = %s", c.name, c.value)
	return
}

func (c Constant) Imports() (imports Imports) {
	imports = NewImports()
	return
}

func (c Constant) Build() (imports Imports, code string, err error) {
	imports = NewImports()
	code, err = c.Code()
	return
}

func NewConstants() (v Constants) {
	v = make([]*Constant, 0, 1)
	return
}

type Constants []*Constant

func (c *Constants) Add(v *Constant) (err error) {
	added := false
	for _, constant := range *c {
		if constant.name == v.name {
			added = true
			break
		}
	}
	if added {
		err = fmt.Errorf("gcg: contants add failed for %s is added", v.name)
		return
	}
	*c = append(*c, v)
	return
}

func (c Constants) Code() (code string, err error) {
	sort.Slice(c, func(i, j int) bool {
		return c[i].name < c[j].name
	})
	buf := bytes.NewBufferString("")
	buf.WriteString("const (")
	for _, constant := range c {
		buf.WriteString(fmt.Sprintf("\t%s = %s", constant.name, constant.value))
	}
	buf.WriteString(")")
	code = buf.String()
	return
}

func (c Constants) Build() (imports Imports, code string, err error) {
	imports = NewImports()
	code, err = c.Code()
	return
}
