package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Unknwon/goconfig"
)

var (
	ConfigFile *goconfig.ConfigFile

	ROOT string

	TemplateDir string
)

const mainIniPath = "/config/env.ini"

func init() {
	curFilename := os.Args[0]
	binaryPath, err := exec.LookPath(curFilename)
	if err != nil {
		panic(err)
	}

	binaryPath, err = filepath.Abs(binaryPath)
	if err != nil {
		panic(err)
	}

	ROOT = filepath.Dir(filepath.Dir(binaryPath))

	configPath := ROOT + mainIniPath

	if !fileExist(configPath) {
		curDir, _ := os.Getwd()
		pos := strings.LastIndex(curDir, "src")
		if pos == -1 {
			// panic("can't find " + mainIniPath)
			fmt.Println("can't find " + mainIniPath)
		} else {
			ROOT = curDir[:pos]

			configPath = ROOT + mainIniPath
		}
	}

	TemplateDir = ROOT + "/template/"

	ConfigFile, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		// panic(err)
		fmt.Println("load config file error:", err)
		ConfigFile, _ = goconfig.LoadFromData([]byte(""))
	}

	if err = loadIncludeFiles(); err != nil {
		panic("load include files error:" + err.Error())
	}

	go signalReload()
}

func ReloadConfigFile() {
	var err error
	configPath := ROOT + mainIniPath
	ConfigFile, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		fmt.Println("reload config file, error:", err)
		return
	}

	if err = loadIncludeFiles(); err != nil {
		fmt.Println("reload files include files error:", err)
		return
	}
	fmt.Println("reload config file successfully！")
}

func SaveConfigFile() error {
	err := goconfig.SaveConfigFile(ConfigFile, ROOT+mainIniPath)
	if err != nil {
		fmt.Println("save config file error:", err)
		return err
	}

	fmt.Println("save config file successfully!")
	return nil
}

func loadIncludeFiles() error {
	includeFile := ConfigFile.MustValue("include_files", "path", "")
	if includeFile != "" {
		includeFiles := strings.Split(includeFile, ",")
		return ConfigFile.AppendFiles(includeFiles...)
	}

	return nil
}

// fileExist 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
