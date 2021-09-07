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


type osFiles []fs.FileInfo

func (a osFiles) Len() int           { return len(a) }
func (a osFiles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a osFiles) Less(i, j int) bool { return a[i].Name() < a[j].Name() }

func dirTreeHelper(out io.Writer, path string, printFiles bool, prefix string) error{
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Something went wrong, %v",err)
	}
	var dir osFiles
	dir, err = f.Readdir(0)
	var folders []fs.FileInfo
	for _, val := range dir{
		if val.IsDir(){
			folders = append( folders, val)
		}
	}
	if !printFiles{
		dir = folders
	}
	sort.Sort(dir)
	for idx, val := range dir{
		{
			if idx == dir.Len()-1{
				fmt.Fprintf(os.Stdout, prefix+ "└───" + val.Name()+ printSize(val) + "\n")
				if(val.IsDir()){
					dirTreeHelper(out, path+"/"+val.Name(), printFiles, prefix+"   ")
				}
			}else{
				fmt.Fprintf(os.Stdout, prefix+"├───"+ val.Name() + printSize(val) + "\n")
				if(val.IsDir()){
					dirTreeHelper(out, path+"/"+val.Name(), printFiles, prefix+"│   ")
				}
			}
		}	
	}
	return nil
}
func printSize(file fs.FileInfo) string{
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


