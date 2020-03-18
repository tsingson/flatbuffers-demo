# flatbuffers-demo

a example for flatbuffers

flatbuffers IDL in ./

go in ./src/cmd/main.go to serialized data and write to ./test.bin

java in ./src/main/java/hello.java read ./test.bin as binary slice and wrap into ByteBuffer , java un-serialized to flatbuffers.

it's work great.