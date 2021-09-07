package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
	// "path/filepath"
	// "strings"
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
	return dirTreeHelper(out, path, printFiles, 0)
}


type osFiles []fs.FileInfo

func (a osFiles) Len() int           { return len(a) }
func (a osFiles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a osFiles) Less(i, j int) bool { return a[i].Name() < a[j].Name() }

func dirTreeHelper(out io.Writer, path string, printFiles bool, level int) error{
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Something went wrong, %v",err)
	}
	var dir osFiles
	dir, err = f.Readdir(0)
	sort.Sort(dir)
	for idx, val := range dir{
		// if level != 0{
		// 	fmt.Print("|")
		// }
		if level > 0{
			fmt.Printf("   ")
		}
		for  i := 1; i < level; i++ {
			fmt.Printf("│   ")
		}
		if idx == dir.Len()-1{
			fmt.Printf("└")
			}else{
			fmt.Printf("├")
		}
		fmt.Fprintf(os.Stdout, "───"+val.Name() + " ")
		isDir :=  val.IsDir()
		if isDir{
			// fmt.Printf("D")
			fmt.Print("\n")
			dirTreeHelper(out, path+"/"+val.Name(), printFiles, level+1)
		}else{
			if printFiles{
				var size = val.Size()
				if size == 0{
					fmt.Printf("(empty)")
				}else{
					fmt.Printf("(%db)",val.Size())
				}
			}
			fmt.Print("\n")
		}
		
	}
	return nil
}

