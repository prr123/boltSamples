# bolt db examples
This repository will contains simple examples of using the embedded db bbolt.
bbolt is a embedded key value db. See https://github.com/etcd-io/bbolt for more information.  

## addEntry  
program that add a bucket and an entry into the bucket.  

objects require a colon(:) char to separate the key from the value. Values should be enclosed in quotes unless the value is a simple string. The following example is an acceptable cli:  
./addEntry /bucket=buck1 /obj=key1:"{x1:val1,x2:val2}" /dbg  

In contrast this cli will result in an error:  
./addEntry /bucket=buck1 /obj=key1:{x1:val1,x2:val2} /dbg  

## listEntries
program that lists all buckets or all entries of a bucket.  

## updEntry
tbd  

## delEntry
tbd  
