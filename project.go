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
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type ProjectOptions struct {
	RemoveBeforeRender bool
	GoProxy            string
	Hooks              []*projectHookInstance
}

type projectHookInstance struct {
	hook           ProjectHook
	commandAndArgs []string
}

type ProjectHook func(ctx context.Context, commandAndArgs ...string) error

func ProjectGoGenerateHook(ctx context.Context, commandAndArgs ...string) (err error) {
	rootDir := ctx.Value("rootDir").(string)
	out := bytes.NewBufferString("")
	errOut := bytes.NewBufferString("")
	args := make([]string, 0, 1)
	args = append(args, "generate")
	if commandAndArgs != nil && len(commandAndArgs) > 0 {
		args = append(args, commandAndArgs...)
	}
	cmd := exec.Command("go", args...)
	cmd.Stdout = out
	cmd.Stderr = errOut
	cmd.Dir = rootDir
	gopath, _ := os.LookupEnv("GOPATH")
	if gopath == "" {
		err = fmt.Errorf("do go generate need $GOPATH")
		return
	}
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOPATH=%s", gopath))
	cmdErr := cmd.Run()
	if cmdErr != nil {
		err = fmt.Errorf("do go generate in %s failed, %v", rootDir, errOut.String())
		return
	}
	if errOut.String() != "" {
		err = fmt.Errorf("do go generate in %s failed, %s", rootDir, errOut.String())
		return
	}
	return
}

type ProjectOption func(*ProjectOptions) error

func ProjectRemoveBeforeRender(ok bool) ProjectOption {
	return func(options *ProjectOptions) error {
		options.RemoveBeforeRender = ok
		return nil
	}
}

func ProjectGoProxy(v string) ProjectOption {
	return func(options *ProjectOptions) error {
		options.GoProxy = v
		return nil
	}
}

func ProjectRenderHook(v ProjectHook, commandAndArgs ...string) ProjectOption {
	return func(options *ProjectOptions) error {
		if v == nil {
			return fmt.Errorf("project hook is nil")
		}
		options.Hooks = append(options.Hooks, &projectHookInstance{
			hook:           v,
			commandAndArgs: commandAndArgs,
		})
		return nil
	}
}

func NewProject(rootDir string, opts ...ProjectOption) (p *Project, err error) {
	opt := &ProjectOptions{
		RemoveBeforeRender: false,
		Hooks:              make([]*projectHookInstance, 0, 1),
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

func (p *Project) AddFile(filename string, file *File) {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		panic(fmt.Errorf("gcg: project add file failed for filename is empty"))
		return
	}
	if file == nil {
		return
	}
	ext := filepath.Ext(filename)
	if ext == "" {
		filename = fmt.Sprintf("%s.go", filename)
	} else if ext != "go" {
		filename = fmt.Sprintf("%s.go", filename)
	}
	_, has := p.files[filename]
	if has {
		panic(fmt.Errorf("gcg: project add file failed for filename is exist"))
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
	v = newFolder(filepath.Join(p.path, name))
	p.folders[name] = v
	return
}

func (p *Project) Render() (err error) {
	// dir
	err = p.renderDir()
	if err != nil {
		return
	}
	// mod
	err = p.renderMod()
	if err != nil {
		return
	}
	// resource
	err = p.renderResourceFiles()
	if err != nil {
		return
	}
	// code file
	err = p.renderCodeFiles()
	if err != nil {
		return
	}
	// go mod tidy
	err = p.doGoModTidy()
	if err != nil {
		return
	}
	// hook
	err = p.doHooks()
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
	createErr := os.MkdirAll(p.path, 0666)
	if createErr != nil {
		err = fmt.Errorf("gcg: project render failed for create dir failed, %v", createErr)
		return
	}
	return
}

func (p *Project) renderMod() (err error) {
	modContent := p.mod.String()
	err = ioutil.WriteFile(filepath.Join(p.path, "go.mod"), []byte(modContent), 0666)
	if err != nil {
		err = fmt.Errorf("gcg: project render failed for create go.mod, %v", err)
		return
	}
	err = ioutil.WriteFile(filepath.Join(p.path, "go.sum"), []byte(""), 0666)
	if err != nil {
		err = fmt.Errorf("gcg: project render failed for create go.sum, %v", err)
		return
	}
	return
}

func (p *Project) renderResourceFiles() (err error) {
	if len(p.resources) == 0 {
		return
	}
	for name, content := range p.resources {
		err = ioutil.WriteFile(filepath.Join(p.path, name), content, 0666)
		if err != nil {
			err = fmt.Errorf("gcg: project render failed for write %s, %v", name, err)
			return
		}
	}
	return
}

func (p *Project) renderCodeFiles() (err error) {
	if len(p.files) > 0 {
		for name, file := range p.files {
			name = filepath.Join(p.path, name)
			w := FileRender(name, true)
			wErr := file.Render(w)
			if wErr != nil {
				err = fmt.Errorf("gcg: project render failed for write %s, %v", name, wErr)
				return
			}
			closeErr := w.Close()
			if closeErr != nil {
				err = fmt.Errorf("gcg: project render failed for write %s, %v", name, closeErr)
				return
			}
		}
	}
	if len(p.folders) > 0 {
		for _, folder := range p.folders {
			folderPath := folder.name
			wErr := folder.Write()
			if wErr != nil {
				err = fmt.Errorf("gcg: project render failed for write %s, %v", folderPath, wErr)
				return
			}
		}
	}
	return
}

func (p *Project) doGoModTidy() (err error) {
	out := bytes.NewBufferString("")
	errOut := bytes.NewBufferString("")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = out
	cmd.Stderr = errOut
	cmd.Dir = p.path
	gopath, _ := os.LookupEnv("GOPATH")
	if gopath == "" {
		err = fmt.Errorf("do go mod tidy need $GOPATH")
		return
	}
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOPATH=%s", gopath))
	if p.opt.GoProxy != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("GOPROXY=%s", p.opt.GoProxy))
	}
	cmdErr := cmd.Run()
	if cmdErr != nil {
		err = fmt.Errorf("do go mod tidy in %s failed, %v", p.path, errOut.String())
		return
	}
	if errOut.String() != "" {
		err = fmt.Errorf("do go mod tidy in %s failed, %s", p.path, errOut.String())
		return
	}
	return
}

func (p *Project) doHooks() (err error) {
	if len(p.opt.Hooks) == 0 {
		return
	}
	ctx := context.TODO()
	ctx = context.WithValue(ctx, "rootDir", p.path)
	for _, hook := range p.opt.Hooks {
		err = hook.hook(ctx, hook.commandAndArgs...)
		if err != nil {
			err = fmt.Errorf("do hook failed, %v", err)
			return
		}
	}
	return
}

func newFolder(name string) *Folder {
	name = strings.TrimSpace(name)
	if name == "" {
		panic("gcg: new folder failed for name is empty")
	}
	return &Folder{
		name:      name,
		files:     make(map[string]*File),
		resources: make(map[string][]byte),
		children:  make(map[string]*Folder),
	}
}

type Folder struct {
	name      string
	files     map[string]*File
	resources map[string][]byte
	children  map[string]*Folder
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
	ext := filepath.Ext(filename)
	if ext == "" {
		filename = fmt.Sprintf("%s.go", filename)
	} else if ext != "go" {
		filename = fmt.Sprintf("%s.go", filename)
	}
	_, has := f.files[filename]
	if has {
		panic("gcg: folder add file failed for filename is exist")
		return
	}
	f.files[filename] = file
	return
}

func (f *Folder) AddResourceFile(filename string, content []byte) (err error) {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		err = fmt.Errorf("gcg: folder add resource file failed for filename is empty")
		return
	}
	if content == nil {
		err = fmt.Errorf("gcg: folder add resource file failed for content is empty")
		return
	}
	_, has := f.resources[filename]
	if has {
		err = fmt.Errorf("gcg: folder add resource file failed for filename is exist")
		return
	}
	f.resources[filename] = content
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
	v = newFolder(filepath.Join(f.name, name))
	f.children[name] = v
	return
}

func (f *Folder) Write() (err error) {
	createErr := os.MkdirAll(f.name, 0666)
	if createErr != nil {
		err = fmt.Errorf("write %s failed, %v", f.name, createErr)
		return
	}
	if len(f.files) > 0 {
		for name, file := range f.files {
			name = filepath.Join(f.name, name)
			w := FileRender(name, true)
			wErr := file.Render(w)
			if wErr != nil {
				err = fmt.Errorf("write %s failed, %v", name, wErr)
				return
			}
			closeErr := w.Close()
			if closeErr != nil {
				err = fmt.Errorf("write %s failed, %v", name, closeErr)
				return
			}
		}
	}
	if len(f.resources) > 0 {
		for name, content := range f.resources {
			name = filepath.Join(f.name, name)
			err = ioutil.WriteFile(name, content, 0666)
			if err != nil {
				err = fmt.Errorf("write %s failed, %v", name, err)
				return
			}
		}
	}
	if len(f.children) > 0 {
		for _, child := range f.children {
			wErr := child.Write()
			if wErr != nil {
				err = fmt.Errorf("gcg: project render failed for write %v", wErr)
				return
			}
		}
	}
	return
}
