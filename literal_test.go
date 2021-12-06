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

package gcg_test

import (
	"fmt"
	"github.com/aacfactory/gcg"
	"os"
	"testing"
)

func TestLiteral(t *testing.T) {
	n := gcg.Literal(nil)
	fmt.Println(n.Render(os.Stdout))
	s := gcg.Literal("string")
	fmt.Println(s.Render(os.Stdout))
	i := gcg.Literal(1)
	fmt.Println(i.Render(os.Stdout))
	i8 := gcg.Literal(int8(2))
	fmt.Println(i8.Render(os.Stdout))
	u := gcg.Literal(uint(3))
	fmt.Println(u.Render(os.Stdout))
	u64 := gcg.Literal(uint64(64))
	fmt.Println(u64.Render(os.Stdout))
	cp := gcg.Literal(complex64(65.44))
	fmt.Println(cp.Render(os.Stdout))
	p := gcg.Literal([]byte(`xx`))
	fmt.Println(p.Render(os.Stdout))
	b := gcg.LiteralByte(byte('b'))
	fmt.Println(b.Render(os.Stdout))
}
