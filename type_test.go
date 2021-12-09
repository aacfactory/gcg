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

package gcg_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aacfactory/gcg"
)

func TestType(t *testing.T) {
	// struct
	s := gcg.Struct()
	s.AddField("Name", gcg.Ident("string"), gcg.StructFieldTag("json", "name", "xml", "name"), "name", "string")
	s.AddField("Age", gcg.Ident("int"), gcg.StructFieldTag("json", "age", "xml", "age"))
	s.AddField("World", gcg.QualifiedIdent(gcg.NewPackage("foo/bar"), "World"), nil)
	ts := gcg.Type("Hello", s.Build())
	fmt.Println(ts.Render(os.Stdout))
	// func
	fnb := gcg.Func()
	fnb.AddParam("p1", gcg.Ident("string"))
	fnb.AddParam("p2", gcg.Star().QualifiedIdent(gcg.NewPackage("foo/bar"), "G"))
	fnb.AddResult("r1", gcg.Ident("string"))
	fnb.AddResult("err", gcg.Ident("error"))
	tf := gcg.Type("Hello", fnb.Build())
	fmt.Println(tf.Render(os.Stdout))
	// ident
	ti := gcg.Type("Hello", gcg.Ident("string"))
	fmt.Println(ti.Render(os.Stdout))
}