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

func FmtPackage() *Package {
	return &Package{
		Name:  "fmt",
		Alias: "",
		Path:  "fmt",
	}
}

func IoPackage() *Package {
	return &Package{
		Name:  "io",
		Alias: "",
		Path:  "io",
	}
}

func IoUtilPackage() *Package {
	return &Package{
		Name:  "ioutil",
		Alias: "",
		Path:  "io/ioutil",
	}
}

func BytesPackage() *Package {
	return &Package{
		Name:  "bytes",
		Alias: "",
		Path:  "bytes",
	}
}

func StringsPackage() *Package {
	return &Package{
		Name:  "strings",
		Alias: "",
		Path:  "strings",
	}
}

func StrconvPackage() *Package {
	return &Package{
		Name:  "strconv",
		Alias: "",
		Path:  "strconv",
	}
}

func SyncPackage() *Package {
	return &Package{
		Name:  "sync",
		Alias: "",
		Path:  "sync",
	}
}

func AtomicPackage() *Package {
	return &Package{
		Name:  "atomic",
		Alias: "",
		Path:  "sync/atomic",
	}
}

func JsonPackage() *Package {
	return &Package{
		Name:  "json",
		Alias: "",
		Path:  "encoding/json",
	}
}
