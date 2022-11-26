package file

import (
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"runtime"
)

type Path struct {
	ProjectDir  string
	ApiDir      string
	BuildDir    string
	CmdDir      string
	ConfigsDir  string
	InternalDir string
}

func NewPath() *Path {
	_, file, _, _ := runtime.Caller(1)
	cmdDir := filepath.Dir(filepath.Dir(file))
	ProjectDir := filepath.Dir(cmdDir)
	color.Red(cmdDir)
	sep := string(os.PathSeparator)
	path := &Path{
		ProjectDir:  ProjectDir + sep,
		ApiDir:      ProjectDir + sep + "api" + sep,
		BuildDir:    ProjectDir + sep + "build" + sep,
		CmdDir:      ProjectDir + sep + "cmd" + sep,
		ConfigsDir:  ProjectDir + sep + "configs" + sep,
		InternalDir: ProjectDir + sep + "internal" + sep,
	}
	return path
}
