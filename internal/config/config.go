package config

import "flag"

type ConfigStruct struct {
	ProjectPath string
	OutPath     string
}

var GlobalConfig ConfigStruct

//Настройка проекта
func Config() {
	projectPath := flag.String("project.path", ".", "Путь до проекта")
	outPath := flag.String("out.path", "out_folder", "Путь до итогового проекта")
	flag.Parse()
	GlobalConfig.ProjectPath = *projectPath
	GlobalConfig.OutPath = *outPath
}
