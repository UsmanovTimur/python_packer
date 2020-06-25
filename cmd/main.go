package main

import (
	"fmt"
	"log"
	"os"
	"python_packer/internal/config"
	"python_packer/internal/folder_tree"
	"python_packer/internal/utils"
)

func main() {
	config.Config()
	folder, err := folder_tree.CreateFolderTree(config.GlobalConfig.ProjectPath)
	if err != nil {
		log.Fatal(err)
	}
	a := folder.FilesArray()
	b := utils.CreateFolderStruct(a)
	if err := os.MkdirAll(config.GlobalConfig.OutPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(b); i++ {
		nf := b[i]
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", config.GlobalConfig.OutPath, nf.ShortPath), os.ModePerm); err != nil {
			log.Fatal(err)
		}
		utils.CopyFile(nf.OriginalPath, fmt.Sprintf("%s/%s/%s", config.GlobalConfig.OutPath, nf.ShortPath, nf.FileName))
	}
}
