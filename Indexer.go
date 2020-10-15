package main

import (
    "fmt"
 //   "io/ioutil"
    "path/filepath"
    "os"
    "log"
    "flag"
)

func indexDir(root_dir string) {
    err := filepath.Walk(root_dir,
    func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        fmt.Println(path, info.Size())
        return nil
    })
    if err != nil {
        log.Println(err)
    }
}

func main() {
//    argsWithProg := os.Args
//    argsWithoutProg := os.Args[1:]
//    arg := os.Args[3]
    // parsing
    var filepath string
    flag.StringVar(&filepath, "filepath", "", "file path to index")
    flag.Parse()

    fmt.Println(filepath)
    indexDir(filepath)

//    fmt.Println(argsWithoutProg)
//    fmt.Println(arg)

}
