# flatbuffers-demo

a example for flatbuffers

flatbuffers IDL in ./src/main/java/events.fbs

go in ./src/cmd/main.go to serialized data and write to ./test.bin

java in ./src/main/java/hello.java read ./test.bin as binary slice and wrap into ByteBuffer , java un-serialized to flatbuffers.



compile the IDL

```
cd ./src/main/java
sh ./build.sh 
```



the binary data file name is hard code in file , change it:

```
/Users/qinshen/go/src/github.com/tsingson/flatbuffers-demo/test.bin
```



to serialize the data

```
go run ./src/cmd/main.go
```



to un-serialize the data , run the java in ./src/main/java/hello.java 



it's work great.