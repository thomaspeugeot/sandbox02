package walk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func createSingleFileInNgTargetPath(filename string) (file *os.File) {
	var err error

	path := filepath.Join(MatTargetPath, filename)

	log.Println("generating : " + path)

	file, err = os.Create(path)
	if err != nil {
		log.Panic(err)
	}

	return file
}

//
// createSingleFileOneLevelsAboveNgTargetPath enable creation of files
// two levels above de src/lib path
func createSingleFileOneLevelsAboveNgTargetPath(filename string) (file *os.File) {
	var err error

	pathToLibDir := filepath.Join(MatTargetPath, "..")

	path := filepath.Join(pathToLibDir, filename)

	log.Println("generating : " + path)

	file, err = os.Create(path)
	if err != nil {
		log.Panic(err)
	}

	return file
}

func createDirAndTreeFilesInNgTargetPath(

	// "Editor"
	prefix string,
	// "s-table" for table
	// "" for sidebar
	suffix string) (fileTS *os.File, fileHTML *os.File, fileCSS *os.File) {

	var err error

	lowerCasePrefix := strings.ToLower(prefix)

	dirpath := filepath.Join(MatTargetPath, fmt.Sprintf("%s%s", lowerCasePrefix, suffix))

	errd := os.Mkdir(dirpath, os.ModePerm)
	if os.IsNotExist(errd) {
		log.Println("creating directory : " + dirpath)
	}
	if os.IsExist(errd) {
		log.Println("directory " + dirpath + " allready exists")
	}

	pathTS := filepath.Join(
		dirpath,
		fmt.Sprintf("%s%s.component.ts", lowerCasePrefix, suffix))

	log.Println("generating : " + pathTS)

	fileTS, err = os.Create(pathTS)
	if err != nil {
		log.Panic(err)
	}

	pathHTML := filepath.Join(
		dirpath,
		fmt.Sprintf("%s%s.component.html", lowerCasePrefix, suffix))

	log.Println("generating : " + pathHTML)

	fileHTML, err = os.Create(pathHTML)
	if err != nil {
		log.Panic(err)
	}

	pathCSS := filepath.Join(
		dirpath,
		fmt.Sprintf("%s%s.component.css", lowerCasePrefix, suffix))

	log.Println("generating : " + pathCSS)

	fileCSS, err = os.Create(pathCSS)
	if err != nil {
		log.Panic(err)
	}

	return fileTS, fileHTML, fileCSS
}
