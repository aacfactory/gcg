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

func TestVar(t *testing.T) {

	x := gcg.Var("hello", gcg.String())
	fmt.Println(x.Build().Render(os.Stdout))

	x.Comments("foo")
	x.Value(gcg.Literal("world"))
	fmt.Println(x.Build().Render(os.Stdout))

}

func TestVars(t *testing.T) {

	x := gcg.Vars()
	a := gcg.Var("hello", gcg.String())
	a.Comments("foo")
	a.Value(gcg.Literal("world"))
	x.Add(a)
	b := gcg.Var("b", gcg.String())
	b.Comments("foo")
	b.Value(gcg.Literal("world"))
	x.Add(b)
	fmt.Println(x.Build().Render(os.Stdout))

}
