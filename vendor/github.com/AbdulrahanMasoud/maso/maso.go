package maso

const version = "1.0.0"

type Maso struct {
	AppName string
	Debug   bool
	Version string
}

func (m *Maso) New(rootPath string) error {
	pathConfig := initPath{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := m.Init(pathConfig)
	if err != nil {
		return err
	}
	return nil
}

func (m *Maso) Init(p initPath) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		//create folder if it does not exist
		err := m.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
