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

func Comments(texts []string) (code Code) {
	code = &comments{
		texts: texts,
	}
	return
}

type comments struct {
	texts []string
}

func (c comments) Render(w io.Writer) (err error) {
	if c.texts == nil || len(c.texts) == 0 {
		return
	}
	buf := bytes.NewBufferString("")
	for _, text := range c.texts {
		buf.WriteString(fmt.Sprintf("// %s\n", strings.TrimSpace(text)))
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		err = fmt.Errorf("render comments failed, %v", err)
		return
	}
	return
}

func (c comments) imports() (imports Imports) {
	imports = emptyImports
	return
}
