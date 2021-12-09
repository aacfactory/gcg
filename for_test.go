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

func TestForRangeArray(t *testing.T) {

	x := gcg.ForRangeArray(gcg.Ident("xx"), "i", "x", gcg.Return())
	fmt.Println(x.Render(os.Stdout))

}

func TestFor(t *testing.T) {
	x := gcg.For(gcg.Token("i := 0"), gcg.Token("i < 10"), gcg.Token("i ++"), gcg.Return())
	fmt.Println(x.Render(os.Stdout))

}

func TestForCond(t *testing.T) {
	x := gcg.ForCond(gcg.Token("i == 0"), gcg.Return())
	fmt.Println(x.Render(os.Stdout))
}

func TestForDead(t *testing.T) {
	x := gcg.ForDead(gcg.Return())
	fmt.Println(x.Render(os.Stdout))
}
