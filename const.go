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
	"io"
	"strings"
)

func Constant(name string, lit interface{}, comments ...string) (code Code) {
	name = strings.TrimSpace(name)
	if name == "" {
		panic(fmt.Errorf("gcg: new contant failed for name is empty"))
		return
	}
	if lit == nil {
		panic(fmt.Errorf("gcg: new contant failed for lit is empty"))
		return
	}
	var litCode Code
	switch lit.(type) {
	case string:
		litCode = Literal(lit)
	case bool:
		litCode = Literal(lit)
	case int, int8, int16, int32, int64:
		litCode = Literal(lit)
	case uint, uint8, uint16, uint32, uint64:
		litCode = Literal(lit)
	case float32, float64:
		litCode = Literal(lit)
	default:
		panic(fmt.Sprintf("gcg: unsupported type for contant: %T", lit))
		return
	}
	code = &constant{
		name:     name,
		lit:      litCode,
		comments: Comments(comments),
	}
	return
}

type constant struct {
	name     string
	lit      Code
	comments Code
}

func (c constant) Render(w io.Writer) (err error) {
	buf := bytes.NewBufferString("")
	_ = c.comments.Render(w)
	buf.WriteString(fmt.Sprintf("const %s = ", c.name))
	err = c.lit.Render(buf)
	if err != nil {
		err = fmt.Errorf("render constant %s failed, %v", c.name, err)
		return
	}
	buf.WriteString("\n")
	p, fmtErr := reformat(buf.Bytes())
	if fmtErr != nil {
		err = fmt.Errorf("render constant %s failed, %v", c.name, err)
		return
	}
	_, err = w.Write(p)
	if err != nil {
		err = fmt.Errorf("render constant %s failed, %v", c.name, err)
		return
	}
	return
}

func (c constant) imports() (imports Imports) {
	imports = emptyImports
	return
}

func Constants(v ...Code) (code Code) {
	cc := constants(make([]*constant, 0, 1))
	if v == nil || len(v) == 0 {
		return
	}
	for _, c := range v {
		x, ok := c.(*constant)
		if !ok {
			panic(fmt.Sprintf("gcg: unsupported type for contant: %T", c))
			return
		}
		cc.add(x)
	}
	code = cc
	return
}

type constants []*constant

func (c *constants) add(v *constant) {
	*c = append(*c, v)
	return
}

func (c constants) Render(w io.Writer) (err error) {
	if len(c) == 0 {
		return
	}
	buf := bytes.NewBufferString("")
	buf.WriteString("const (\n")
	for _, v := range c {
		cc := bytes.NewBufferString("")
		_ = v.comments.Render(cc)
		if cc.Len() > 0 {
			buf.WriteString(fmt.Sprintf("\t%s", cc.String()))
		}
		bb := bytes.NewBufferString("")
		err = v.lit.Render(bb)
		if err != nil {
			err = fmt.Errorf("render constant %s failed, %v", v.name, err)
			return
		}
		buf.WriteString(fmt.Sprintf("\t%s = %s\n", v.name, bb.String()))
	}
	buf.WriteString(")\n")
	p, fmtErr := reformat(buf.Bytes())
	if fmtErr != nil {
		err = fmt.Errorf("render constants failed, %v", err)
		return
	}
	_, err = w.Write(p)
	if err != nil {
		err = fmt.Errorf("render constants failed, %v", err)
		return
	}
	return
}

func (c constants) imports() (imports Imports) {
	imports = emptyImports
	return
}
