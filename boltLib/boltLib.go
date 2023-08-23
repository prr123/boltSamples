// boltLib
// a library that facilitates using boltdb
//
// Author: prr, azulsoftware
// Date: 23. August 2023
// copyright 2023 prr, azul software

package boltLib

import (
    "fmt"
//    "log"
//    "os"
//    "strings"

//    util "github.com/prr123/utility/utilLib"
    bolt "go.etcd.io/bbolt"
)

type DBObj struct {
     Db *bolt.DB
	 Dbg bool
    }

func Initdb(dbPath string)(db *DBObj, err error) {

//    path := "boltTest.db"
    dbobj, err := bolt.Open(dbPath, 0666, nil)
    if err != nil {
        return nil, fmt.Errorf("bolt.Open: %v\n", err)
    }
    dbObj := DBObj {
        Db: dbobj,
    }

    return &dbObj, nil
}

func (dbobj *DBObj) DbClose() (err error){

    db := dbobj.Db
	err = db.Close()
	return err
}


func (dbobj *DBObj) CheckDb() (errList []string, err error) {

    db := dbobj.Db

	err =db.View(func(tx *bolt.Tx) error {
		count := 0
//		for err := range tx.Check(bolt.WithKVStringer(CmdKvStringer())) {
		for err := range tx.Check() {
			fmt.Printf(" %d: %v\n", count, err)
			str := fmt.Sprintf(" %d: %v\n", count, err)
			errList = append(errList, str)
			count++
		}

		// Print summary of errors.
		if count > 0 {
			fmt.Printf("%d errors found\n", count)
			return fmt.Errorf("found %d errors\n", count)
		}

		// Notify user that database is valid.
		fmt.Printf("status: ok\n")
		return nil
	})

	return errList, err
}

func (dbobj *DBObj) AddBucket(par, child string) (err error) {

    db := dbobj.Db

	if dbobj.Dbg {fmt.Printf("parent: %s child: %s\n", par, child)}

    err = db.Update(func(tx *bolt.Tx) error {
//      root, err := tx.CreateBucketIfNotExists([]byte(bucket))
		var err error
		if len(par) > 0 {
			root := tx.Bucket([]byte(par))
			if root == nil {
				return fmt.Errorf("parent bucket: %s does not exist!", par)
			}
			if dbobj.Dbg {fmt.Printf("root: %v\n", root)}
	        chbuck, err := root.CreateBucketIfNotExists([]byte(child))
			if err != nil {
				return fmt.Errorf("could not create bucket: %v", err)
        	}
			if dbobj.Dbg {fmt.Printf("child: %v\n", chbuck)}
		} else {
	        _, err = tx.CreateBucketIfNotExists([]byte(child))
		}
		if err != nil {
			return fmt.Errorf("could not create bucket: %v", err)
        }
        return nil
    })
    return err
}




func (dbobj *DBObj) ListBuckets() (namList []string, err error) {

    db := dbobj.Db
//	namList := []string{}

	err = db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			namList = append(namList, string(name))
//			fmt.Printf("bucket: %s\n",string(name))
			return nil
		})
	})

//	fmt.Printf("namList: %d\n", len(namList))

	return namList, err
}

// from https://github.com/boltdb/bolt/issues/373
func (dbobj *DBObj) DispBuckets(buk *bolt.Bucket, pos int) {

    db := dbobj.Db
	if dbobj.Dbg {fmt.Printf("entering DispBuckets: %d\n", pos)}

    err := db.View(func(tx *bolt.Tx) error {
        var c *bolt.Cursor
        if buk == nil {
            c = tx.Cursor()
            fmt.Println("ROOT")
        } else {
            c = buk.Cursor()
        }
        for k, v := c.First(); k != nil; k, v = c.Next() {
            if k == nil {
                fmt.Println(" ----nil - never") //never will happend
            } else {
                for i := 0; i < pos; i++ {
                    fmt.Print(" ")
                }
                if v == nil {
                    fmt.Printf("%s\n", k) //bucket
                    var buk2 *bolt.Bucket
                    if buk == nil {
                        buk2 = tx.Bucket(k)
                    } else {
                        buk2 = buk.Bucket(k)
                    }
                    dbobj.DispBuckets(buk2, pos+2)
                } else {
                    fmt.Printf("%s=%s\n", k, v) // k = v
                }
            }
        }
        return nil
    })
    TestError("disp2", err)
}

func TestError(msg string, err error) {
    if err != nil {
        fmt.Println(msg)
        panic(err)
    }
}

func (dbobj *DBObj) AddEntry(bucket, key, val string) (err error) {

    db := dbobj.Db
    err = db.Update(func(tx *bolt.Tx) error {
        err := tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(val))
        if err != nil {
            return fmt.Errorf("could not insert entry: %v", err)
        }
        return nil
    })
//    fmt.Println("Added Entry")
    return err
}
