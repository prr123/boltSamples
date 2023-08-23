// dispBuckets
//  displays all buckets of db
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
    flags:=[]string{"dbg","db"}

    useStr := "./listBuckets /db=dbPath [/dbg]"
    helpStr := "program that lists all buckets of  bolt db instance at dbPath"

    if numarg > 3 {
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

	dbPath := ""
    bval, ok := flagMap["db"]
    if !ok {
        log.Fatalf("no db flag set!\n")
    } else {
        if bval.(string) == "none" {log.Fatalf("error: no db Path provided!")}
        dbPath = bval.(string)
    }

    log.Printf("debug: %t\n", dbg)
    log.Printf("db:    %s\n", dbPath)

	dbobj, err := boltLib.Initdb(dbPath)
	if err !=nil {
		log.Fatalf("error -- cannot init db: %v\n", err)
	}

	if dbg {
		fmt.Printf("dbobj: %v\n", dbobj)
	}
//	defer dbobj.db.Close()
	log.Println("success opening boltdb!")
	defer dbobj.DbClose()

	log.Printf("dispBuckets\n")

	dbobj.DispBuckets(nil, 2)

//	PrintBucketList(namList)
	log.Println("success listing buckets!")

}


func PrintBucketList(namList []string) {

	fmt.Println("********* Buckets ********")
	for i:=0; i< len(namList); i++ {
		fmt.Printf("  %d: %s\n", i+1, namList[i])
	}
	fmt.Println("******* end Buckets ******")
}
