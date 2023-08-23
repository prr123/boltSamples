// addBucket
// adds a bucket to the db
//
// Author: prr, azulsoftware
// Date: 23. August 2023
// copyright 2023 prr, azul software


package main

import (
	"fmt"
	"log"
	"os"
//	"strings"

	util "github.com/prr123/utility/utilLib"
	boltLib "db/bbolt/boltLib"
)


func main () {

    numarg := len(os.Args)
    dbg := false
    flags:=[]string{"dbg","parent","child"}

    useStr := "./addBucket [/parent=bucket] /child=child bucket [/dbg]"
    helpStr := "program to add a bucket to a parent bucket or root to the bolt db"

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
    bval, ok := flagMap["parent"]
    if ok {
        if bval.(string) == "none" {log.Fatalf("error: no bucket name provided!")}
        buckStr = bval.(string)
    }

	childStr := ""
    chval, ok := flagMap["child"]
    if !ok {
        log.Fatalf("no child flag set!\n")
    } else {
        if chval.(string) == "none" {log.Fatalf("error: no bucket name provided!")}
        childStr = chval.(string)
    }

    log.Printf("debug: %t\n", dbg)

	if len(buckStr) == 0 {
		log.Printf("no parent bucket -> root bucket\n")
	} else {
		log.Printf("Bucket: %s\n", buckStr)
	}
    log.Printf("Child:  %s\n", childStr)

	dbobj, err := boltLib.Initdb("boltTest.db")
	if err !=nil {
		log.Fatalf("error -- cannot init db: %v\n", err)
	}

	dbobj.Dbg = dbg
	defer dbobj.Db.Close()
	log.Println("success opening boltdb!")

	err = dbobj.AddBucket(buckStr, childStr)
	if err !=nil {
		log.Fatalf("error -- cannot add bucket: %v\n", err)
	}
	log.Println("success adding bucket!")

}

