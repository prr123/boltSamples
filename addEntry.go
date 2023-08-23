// addEntry
// adds a value to the db
//
// Author: prr, azulsoftware
// Date: 23. August 2023
// copyright 2023 prr, azul software


package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	util "github.com/prr123/utility/utilLib"
	bolt "go.etcd.io/bbolt"
)

func main () {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","bucket","obj"}

    useStr := "./addEntry /bucket=bname /obj=kv [/dbg]"
    helpStr := "program to add an obkject into a bucket in the bolt db"

    if numarg > 4 {
        fmt.Printf("too many arguments in cl!\n")
        fmt.Printf("usage: %s\n", useStr)
        os.Exit(-1)
    }

    if numarg > 1 && os.Args[1] == "help" {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s\n", useStr)
        os.Exit(1)
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    _, ok := flagMap["dbg"]
    if ok {dbg = true}
    if dbg {
        fmt.Printf("dbg -- flag list:\n")
        for k, v :=range flagMap {
            fmt.Printf("  flag: /%s value: %s\n", k, v)
        }
    }

	buckStr := ""
    bval, ok := flagMap["bucket"]
    if !ok {
        log.Fatalf("no bucket flag set!\n")
    } else {
        if bval.(string) == "none" {log.Fatalf("error: no bucket name provided!")}
        buckStr = bval.(string)
    }

	objStr := ""
    oval, ok := flagMap["obj"]
    if !ok {
        log.Fatalf("no obj flag set!\n")
    } else {
        if oval.(string) == "none" {log.Fatalf("error: no object value provided!")}
        objStr = oval.(string)
    }

    if dbg {fmt.Printf("Object: %s\n", objStr)}

	idx := strings.Index(objStr, ":")
	if idx == -1 {
		log.Fatalf("error: objStr has no kv seperator!")
	}
	objKey:=string(objStr[:idx])
	objVal := string(objStr[idx+1:])

    log.Printf("debug: %t\n", dbg)
    log.Printf("Bucket: %s\n", buckStr)
	log.Printf("Object Key: %s Value: %s\n", objKey, objVal)

	db, err := Initdb()
	if err !=nil {
		log.Fatalf("error -- cannot init db: %v\n", err)
	}
	defer db.Close()
	log.Println("success opening boltdb!")
}


func Initdb()(db *bolt.DB, err error) {

	path := "boltTest.db"
	db, err = bolt.Open(path, 0666, nil)
	if err != nil {
  		return nil, fmt.Errorf("bolt.Open: %v\n", err)
	}
	return db, nil

}
