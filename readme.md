# Tree util clone

### Description
This is a simple implementation of unix tree command that prints a tree of files and folders of a specified directory.

### Example:
```
└───testdata
    ├───project
    ├───static
    │    ├───a_lorem
    │    │    └───ipsum
    │    ├───css
    │    ├───html
    │    ├───js
    │    └───z_lorem
    │        └───ipsum
    └───zline
        └───lorem
            └───ipsum
```
### Required arguments
This program requires one argument - path to the folder, i.e. "." would be the path to the current folder

### Optional arguments
This program also can take one flag, -f to print files, otherwise it would print only directories.

### Remarks
This is an implementation for hmwk 1 of Разработка веб-сервисов на Go - основы языка on Coursera