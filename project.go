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
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type ProjectOptions struct {
	RemoveBeforeRender bool
	Hooks              []ProjectHook
}

type ProjectHook interface {
}

type ProjectOption func(*ProjectOptions) error

func ProjectRemoveBeforeRender(ok bool) ProjectOption {
	return func(options *ProjectOptions) error {
		options.RemoveBeforeRender = ok
		return nil
	}
}

func ProjectRenderHook(v ProjectHook) ProjectOption {
	return func(options *ProjectOptions) error {
		if v == nil {
			return fmt.Errorf("project hook is nil")
		}
		options.Hooks = append(options.Hooks, v)
		return nil
	}
}

func NewProject(rootDir string, opts ...ProjectOption) (p *Project, err error) {
	opt := &ProjectOptions{
		RemoveBeforeRender: false,
		Hooks:              make([]ProjectHook, 0, 1),
	}
	if opts != nil && len(opts) > 0 {
		for _, option := range opts {
			optErr := option(opt)
			if optErr != nil {
				err = fmt.Errorf("gcg: new project failed, %v", optErr)
				return
			}
		}
	}
	rootDir = strings.TrimSpace(rootDir)
	if rootDir == "" {
		err = fmt.Errorf("gcg: new project failed for root dir is empty")
		return
	}
	if !path.IsAbs(rootDir) {
		rootDir, err = filepath.Abs(rootDir)
		if err != nil {
			err = fmt.Errorf("gcg: new project failed for making absolute representation of %s", rootDir)
			return
		}
	}
	rootDir = strings.ReplaceAll(rootDir, "\\", "/")
	rootDir = filepath.Clean(rootDir)

	gopath, hasGoPathEnv := os.LookupEnv("GOPATH")
	if !hasGoPathEnv {
		err = fmt.Errorf("gcg: new project failed for GOPATH ENV was not found")
		return
	}
	if !path.IsAbs(gopath) {
		gopath, err = filepath.Abs(gopath)
		if err != nil {
			err = fmt.Errorf("gcg: new project failed for making absolute representation of %s", gopath)
			return
		}
	}
	gopath = strings.ReplaceAll(gopath, "\\", "/") + "/"
	if strings.Index(rootDir, gopath) != 0 {
		err = fmt.Errorf("gcg: new project failed for root dir is not in GOPATH")
		return
	}
	modPath := rootDir[len(gopath):]
	_, modName := filepath.Split(modPath)

	version := runtime.Version()
	versionItems := strings.Split(version, ".")
	version = versionItems[0][2:] + "." + versionItems[1]
	mod := &Module{
		Name:      modName,
		Path:      modPath,
		GoVersion: version,
		Requires:  make([]Require, 0, 1),
	}

	p = &Project{
		opt:       opt,
		path:      rootDir,
		mod:       mod,
		resources: make(map[string][]byte),
		files:     make(map[string]*File),
		folders:   make(map[string]*Folder),
	}
	return
}

type Project struct {
	opt       *ProjectOptions
	path      string
	mod       *Module
	resources map[string][]byte
	files     map[string]*File
	folders   map[string]*Folder
}

func (p *Project) SetGoVersion(v string) {
	p.mod.GoVersion = strings.TrimSpace(v)
}

func (p *Project) AddRequire(v Require) {
	p.mod.Requires = append(p.mod.Requires, v)
}

func (p *Project) AddResourceFile(filename string, content []byte) (err error) {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		err = fmt.Errorf("gcg: project add resource file failed for filename is empty")
		return
	}
	if content == nil {
		err = fmt.Errorf("gcg: project add resource file failed for content is empty")
		return
	}
	_, has := p.resources[filename]
	if has {
		err = fmt.Errorf("gcg: project add resource file failed for filename is exist")
		return
	}
	p.resources[filename] = content
	return
}

func (p *Project) AddFile(filename string, file *File) (err error) {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		err = fmt.Errorf("gcg: project add file failed for filename is empty")
		return
	}
	if file == nil {
		return
	}
	ext := filepath.Ext(filename)
	if ext == "" {
		filename = fmt.Sprintf("%s.go", filename)
	} else if ext != "go" {
		err = fmt.Errorf("gcg: project add file failed for filename is invalid")
		return
	}
	_, has := p.files[filename]
	if has {
		err = fmt.Errorf("gcg: project add file failed for filename is exist")
		return
	}
	p.files[filename] = file
	return
}

func (p *Project) AddFolder(name string) (v *Folder) {
	name = strings.TrimSpace(name)
	if name == "" {
		panic("gcg: project add folder failed for name is empty")
		return
	}
	_, has := p.folders[name]
	if has {
		panic("gcg: project add folder failed for name is exist")
		return
	}

	v = newFolder(name)
	p.folders[name] = v
	return
}

func (p *Project) Render() (err error) {
	err = p.renderDir()
	if err != nil {
		return
	}

	return
}

func (p *Project) renderDir() (err error) {
	if FileExist(p.path) {
		if !p.opt.RemoveBeforeRender {
			err = fmt.Errorf("gcg: project render failed for dir is exist")
			return
		}
		rmErr := os.RemoveAll(p.path)
		if rmErr != nil {
			err = fmt.Errorf("gcg: project render failed for remove dir failed, %v", rmErr)
			return
		}
	}
	createErr := os.MkdirAll(p.path, os.ModeDir)
	if createErr != nil {
		err = fmt.Errorf("gcg: project render failed for create dir failed, %v", createErr)
		return
	}
	return
}

func (p *Project) renderMod() (err error) {

	return
}

func newFolder(name string) *Folder {
	name = strings.TrimSpace(name)
	if name == "" {
		panic("gcg: new folder failed for name is empty")
	}
	return &Folder{
		name:     name,
		files:    make(map[string]*File),
		children: make(map[string]*Folder),
	}
}

type Folder struct {
	name     string
	files    map[string]*File
	children map[string]*Folder
}

func (f *Folder) AddFile(filename string, file *File) {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		panic("gcg: folder add file failed for filename is empty")
		return
	}
	if file == nil {
		return
	}
	_, has := f.files[filename]
	if has {
		panic("gcg: folder add file failed for filename is exist")
		return
	}
	f.files[filename] = file
	return
}

func (f *Folder) AddFolder(name string) (v *Folder) {
	name = strings.TrimSpace(name)
	if name == "" {
		panic("gcg: folder add sub folder failed for name is empty")
		return
	}
	_, has := f.children[name]
	if has {
		panic("gcg: folder add sub file failed for name is exist")
		return
	}
	v = newFolder(name)
	f.children[name] = v
	return
}
