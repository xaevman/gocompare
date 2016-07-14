package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

// default values
const (
    DEFAULT_BLOCK_SIZE = 2048
    READ_BUFFER_SIZE   = 32 * 1024 * 1024
    WRITE_BUFFER_SIZE  = 32 * 1024 * 1024
)

var (
    fileOne   string
    fileTwo   string
    blockSize int
)

func main() {
    parseFlags()

    f1, err := os.Open(fileOne)
    if err != nil {
        panic(err)
    }
    defer f1.Close()

    f2, err := os.Open(fileTwo)
    if err != nil {
        panic(err)
    }
    defer f2.Close()

    r1 := bufio.NewReaderSize(f1, READ_BUFFER_SIZE)
    r2 := bufio.NewReaderSize(f2, READ_BUFFER_SIZE)
    var b1 byte
    var b2 byte

    for i := 0; err != io.EOF; i++ {
        b1, err = r1.ReadByte()
        if err != nil && err != io.EOF {
            panic(err)
        }

        b2, err = r2.ReadByte()
        if err != nil && err != io.EOF {
            panic(err)
        }

        if b1 != b2 {
            fmt.Printf("[%d] mismatch, %d != %d\n", i, b1, b2)
            break
        }
    }

    fmt.Println("Files are the same")
}

func parseFlags() {
    flag.StringVar(&fileOne, "f1", "", "")
    flag.StringVar(&fileTwo, "f2", "", "")
    flag.IntVar(&blockSize, "blockSize", DEFAULT_BLOCK_SIZE, "Size of blocks to process at a time")
    flag.Parse()
}
