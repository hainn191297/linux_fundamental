package main

import (
    "flag"
    "fmt"
    "io/fs"
    "log"
    "os"
    "path/filepath"
    "syscall"
)

func main() {
    flag.Parse()
    root := "."
    if flag.NArg() > 0 {
        root = flag.Arg(0)
    }

    err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            fmt.Fprintf(os.Stderr, "error accessing %s: %v\n", path, err)
            return nil
        }
        info, statErr := d.Info()
        if statErr != nil {
            fmt.Fprintf(os.Stderr, "stat error %s: %v\n", path, statErr)
            return nil
        }

        var inode uint64
        var mode fs.FileMode
        var size int64
        mode = info.Mode()
        size = info.Size()

        if sys, ok := info.Sys().(*syscall.Stat_t); ok {
            inode = sys.Ino
        }

        fmt.Printf("%s | inode=%d | mode=%s | size=%d\n", path, inode, mode.Perm(), size)
        return nil
    })
    if err != nil {
        log.Fatal(err)
    }
}

