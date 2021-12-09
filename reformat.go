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
	"go/format"
	"io"
)

func reformat(p []byte) (b []byte, err error) {
	b, err = format.Source(p)
	if err != nil {
		err = fmt.Errorf("reformat code failed, %v", err)
		return
	}
	return
}

func newReformatWriter() *reformatWriter {
	return &reformatWriter{
		buf: bytes.NewBufferString(""),
	}
}

type reformatWriter struct {
	buf *bytes.Buffer
}

func (r *reformatWriter) WriteTo(w io.Writer) (n int64, err error) {
	b, fmtErr := reformat(r.buf.Bytes())
	if fmtErr != nil {
		err = fmtErr
		return
	}
	wn, wErr := w.Write(b)
	if wErr != nil {
		err = wErr
		return
	}
	n = int64(wn)
	return
}

func (r *reformatWriter) Write(p []byte) (n int, err error) {
	n, err = r.buf.Write(p)
	return
}
