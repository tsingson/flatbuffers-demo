package main

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/tsingson/flatbuffers-demo/Serializer"
)

func main() {
	v := &Serializer.EventT{
		Action: Serializer.ActionFOCUS,
		Data: &Serializer.DataTypeT{
			Type: Serializer.DataTypeStringList,
			Value: &Serializer.StringListT{
				Values: []string{"test", "tsingson", "java"},
			},
		},
		ShouldCensor:   true,
		WindowId:       "test",
		Sequence:       1,
		SequenceIndex:  1,
		ParentWindowId: "who",
	}
	b := flatbuffers.NewBuilder(0)

	b.Finish(v.Pack(b))
	buf := b.FinishedBytes()

	err := WriteToFile(buf, "/Users/qinshen/go/src/github.com/tsingson/flatbuffers-demo/test.bin")

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("flatbuffers data write success")
	}
}
