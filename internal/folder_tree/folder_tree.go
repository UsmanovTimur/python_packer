package folder_tree

import (
	"fmt"
	"io/ioutil"
)

type Folder struct {
	Path       string   //Путь до директории
	Files      []string //Список файлов
	SubFolders []Folder //Список подкаталогов
}

//Построение дерева каталогов для указанного пути
func CreateFolderTree(path string) (*Folder, error) {
	//TODO
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	outFolder := Folder{
		Path:       path,
		Files:      []string{},
		SubFolders: []Folder{},
	}
	for _, file := range files {
		if file.IsDir() {
			if subFolder, err := CreateFolderTree(path + "/" + file.Name()); err == nil {
				outFolder.SubFolders = append(outFolder.SubFolders, *subFolder)
			}
		} else {
			outFolder.Files = append(outFolder.Files, file.Name())
		}
	}
	return &outFolder, err
}

//Вывод всех файлов в виде строкового массива
func (f *Folder) FilesArray() []string {
	outArray := []string{}
	for i := 0; i < len(f.Files); i++ {
		outArray = append(outArray, fmt.Sprintf("%s/%s", f.Path, f.Files[i]))
	}
	for i := 0; i < len(f.SubFolders); i++ {
		fArray := f.SubFolders[i].FilesArray()
		for j := 0; j < len(fArray); j++ {
			outArray = append(outArray, fArray[j])
		}
	}
	return outArray
}
