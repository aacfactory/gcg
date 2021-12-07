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

func newToken(content string, imports ...*Import) *token {
	t := &token{
		content:  strings.TrimSpace(content),
		_imports: NewImports(),
	}
	if imports != nil {
		for _, i := range imports {
			t._imports.Add(i)
		}
	}
	return t
}

type token struct {
	content  string
	_imports Imports
}

func (t token) Render(w io.Writer) (err error) {
	if t.content == "" {
		return
	}
	_, err = w.Write([]byte(t.content))
	if err != nil {
		err = fmt.Errorf("render token %s failed, %v", t.content, err)
		return
	}
	return
}

func (t token) imports() (imports Imports) {
	imports = t._imports
	return
}

func (t token) addImport(i *Import) {
	t._imports.Add(i)
	return
}
