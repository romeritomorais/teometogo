package dirutils

import "path/filepath"

type ProjectDir struct {
	Base   string
	Cmd    string
	Data   string
	Models string
}

func (p *ProjectDir) GetDirs(filename string) {
	abspath, err := filepath.Abs(filename)
	if err != nil {
		panic("Impossible to get abspath")
	}

	p.Cmd = filepath.Dir(abspath)
	p.Base = filepath.Dir(p.Cmd)
	p.Data = filepath.Join(p.Base, "data")
	p.Models = filepath.Join(p.Base, "models")
}
