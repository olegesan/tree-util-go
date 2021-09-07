package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error{
	return dirTreeHelper(out, path, printFiles, "")
}

// Make a new type that will implement sort.Sort interface
type osFiles []fs.FileInfo

func (a osFiles) Len() int           { return len(a) }
func (a osFiles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a osFiles) Less(i, j int) bool { return a[i].Name() < a[j].Name() }

func dirTreeHelper(out io.Writer, path string, printFiles bool, prefix string) error{
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("something went wrong, %v",err)
	}
	var dir osFiles
	dir, _ = f.Readdir(0)
	
	// Filter all files leaving only folders
	if !printFiles{
		var folders []fs.FileInfo
		for _, val := range dir{
			if val.IsDir(){
				folders = append( folders, val)
			}
		}
		dir = folders
	}

	sort.Sort(dir)
	for idx, val := range dir{
		{
			if idx == dir.Len()-1{
				fmt.Fprintf(out, prefix+ "└───" + val.Name()+ getFileSize(val) + "\n")
				if(val.IsDir()){
					dirTreeHelper(out, path+"/"+val.Name(), printFiles, prefix+"\t")
				}
			}else{
				fmt.Fprintf(out, prefix+"├───"+ val.Name() + getFileSize(val) + "\n")
				if(val.IsDir()){
					dirTreeHelper(out, path+"/"+val.Name(), printFiles, prefix+"│\t")
				}
			}
		}	
	}
	return nil
}

func getFileSize(file fs.FileInfo) string{
	if !file.IsDir(){
		var size = file.Size()
		if size == 0{
			return (" (empty)")
		}else{
			return fmt.Sprintf(" (%db)",file.Size())
		}
	}
	return ""
}


