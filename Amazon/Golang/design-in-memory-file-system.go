type FileSystem struct {
    root *Dir
}


type Dir struct {
    subDirs map[string]*Dir
    files   map[string]string
}


func Constructor() FileSystem {
    return FileSystem{
        &Dir{
            make(map[string]*Dir),
            make(map[string]string),
        },
    }
}


func (this *FileSystem) Ls(path string) []string {
    curr := this.root
    files := make([]string, 0)
    
    if path != "/" {
        dirPartitions := strings.Split(path, "/")
        for i := 1; i < len(dirPartitions)-1; i++ {
            curr = curr.subDirs[dirPartitions[i]]
        }
        if _, ok := curr.files[dirPartitions[len(dirPartitions)-1]]; ok {
            files = append(files, dirPartitions[len(dirPartitions)-1])
            return files
        } else {
            curr = curr.subDirs[dirPartitions[len(dirPartitions)-1]]
        }
    }
    
    for dirName := range curr.subDirs {
        files = append(files, dirName)
    }
    for fileName := range curr.files {
        files = append(files, fileName)
    }
    
    sort.Strings(files)
    
    return files
}


func (this *FileSystem) Mkdir(path string)  {
    curr := this.root
    dirPartitions := strings.Split(path, "/")
    for i := 1; i < len(dirPartitions); i++ {
        if _, ok := curr.subDirs[dirPartitions[i]]; !ok {
            curr.subDirs[dirPartitions[i]] = &Dir{make(map[string]*Dir), make(map[string]string)}
        }
        curr = curr.subDirs[dirPartitions[i]]
    }
}


func (this *FileSystem) AddContentToFile(filePath string, content string)  {
    curr := this.root
    dirPartitions := strings.Split(filePath, "/")
    for i := 1; i < len(dirPartitions)-1; i++ {
        curr = curr.subDirs[dirPartitions[i]]
    }
    oldContent := curr.files[dirPartitions[len(dirPartitions)-1]]
    curr.files[dirPartitions[len(dirPartitions)-1]] = oldContent + content
}


func (this *FileSystem) ReadContentFromFile(filePath string) string {
    curr := this.root
    dirPartitions := strings.Split(filePath, "/")
    for i := 1; i < len(dirPartitions)-1; i++ {
        curr = curr.subDirs[dirPartitions[i]]
    }
    return curr.files[dirPartitions[len(dirPartitions)-1]]
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
