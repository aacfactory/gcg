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

import (
	"fmt"
	"io"
	"strings"
)

func Keyword(ident string) (code Code) {
	code = &keyword{
		ident: strings.TrimSpace(ident),
	}
	return
}

type keyword struct {
	ident string
}

func (k keyword) Render(w io.Writer) (err error) {
	_, err = w.Write([]byte(k.ident + " "))
	if err != nil {
		err = fmt.Errorf("render keywork %s failed, %v", k.ident, err)
		return
	}
	return
}

func (k keyword) imports() (imports Imports) {
	imports = emptyImports
	return
}
