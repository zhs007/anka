package base

import (
	"io/ioutil"
	"os"
	"path"
	"sync"

	"gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	WebServAddr  string
	GrpcServAddr string
	LogPath      string
	ErrPath      string
	LogLevel     string
	DbPath       string
	GuestDbName  string
}

var cfg Config
var onceCfg sync.Once

// LoadConfig - load config
func load(filename string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer fi.Close()
	fd, err1 := ioutil.ReadAll(fi)
	if err1 != nil {
		return err1
	}

	err2 := yaml.Unmarshal(fd, &cfg)
	if err2 != nil {
		return err2
	}

	return nil
}

func procWithBaseDir(basedir string) {
	cfg.DbPath = path.Join(basedir, cfg.DbPath)
	cfg.ErrPath = path.Join(basedir, cfg.ErrPath)
	cfg.LogPath = path.Join(basedir, cfg.LogPath)
}

// LoadConfig - load config
func LoadConfig(filename string, basedir string) (err error) {
	onceCfg.Do(func() {
		err = load(filename)
		procWithBaseDir(basedir)

		MakePath(cfg.DbPath)
	})

	return
}

// GetConfig - get Config
func GetConfig() *Config {
	return &cfg
}
