package utils

import (
	"python_packer/internal/config"
	"strings"
)

type ParseFolder struct {
	OriginalPath string
	FullPath     string
	ShortPath    string
	FileName     string
}

//
func CreateFolderStruct(folders []string) []ParseFolder {
	outArray := []ParseFolder{}
	for i := 0; i < len(folders); i++ {
		str := folders[i]
		c := str[len(config.GlobalConfig.ProjectPath)+1:]
		if len(c) > 0 && c[0] != '.' {
			newFolder := ParseFolder{
				OriginalPath: str,
			}
			str = strings.Replace(str, "/__pycache__", "", -1)
			c = strings.Replace(c, "__pycache__/", "", -1)
			splitString := strings.Split(c, "/")
			filename := splitString[len(splitString)-1]
			splitFile := strings.Split(filename, ".")
			if len(splitFile) > 1 {
				//Пропускаем py файлы
				if splitFile[len(splitFile)-1] == "py" {
					continue
				}
				//Модифицируем pyc файлы для удачной подмены
				if splitFile[len(splitFile)-1] == "pyc" {
					filename = splitFile[0] + ".pyc"
				}
			}
			shortName := ""
			if len(splitString) > 1 {
				for j := 0; j < len(splitString)-1; j++ {
					shortName += "/" + splitString[j]
				}
				if len(shortName) > 0 {
					shortName = shortName[1:]
				}
			}
			newFolder.FullPath = config.GlobalConfig.ProjectPath + "/" + shortName
			newFolder.ShortPath = shortName
			newFolder.FileName = filename
			outArray = append(outArray, newFolder)
		}
	}
	return outArray
}
