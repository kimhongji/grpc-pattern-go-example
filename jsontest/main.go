package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"

	"grpc-pattern-go-example/proto/ecommerce"
)

type A struct  {
	Status string
}

type B struct  {
	Status string
	Ping string
}

func main() {
	// simulation 1: A 구조의 데이터를 B 구조로 언마샬링
	bstruct := B{}
	testAstruct := &A{
		Status: "connected",
	}
	marshaledA, _ := json.Marshal(testAstruct)
	if err := json.Unmarshal(marshaledA, &bstruct); err != nil {
		fmt.Println("err")
	}

	fmt.Println("A to B unmarshal")
	fmt.Println("status: ", bstruct.Status)
	fmt.Println("ping: ", bstruct.Ping)

	// simulation 2: B 구조의 데이터를 A 구조로 언마샬링
	astruct := A{}
	testBstruct := &B{
		Status: "connected",
		Ping:   "2022-02-16 16:3",
	}
	marshaledB, _ := json.Marshal(testBstruct)
	if err := json.Unmarshal(marshaledB, &astruct); err != nil {
		fmt.Println("err")
	}

	fmt.Println("B to A unmarshal")
	fmt.Println("status: ", astruct.Status)

	// simulation 3: B 구조의 데이터를 B 구조로 언마샬링
	newBstruct := B{}
	if err := json.Unmarshal(marshaledB, &newBstruct); err != nil {
		fmt.Println("err")
	}

	fmt.Println("B to B unmarshal")
	fmt.Println("status: ", newBstruct.Status)
	fmt.Println("ping: ", newBstruct.Ping)

	fmt.Println("==== protojson and json test ====")
	/* protojson unmarsal example */
	exampleJson := &ecommerce.Outer{
		Inner: []*ecommerce.Outer_Inner{{Content: "nayana"}, {Content: "seccconnd"}},
	}
	marshaledjson, _ := json.Marshal(exampleJson)
	outputProto := &ecommerce.Outer{}
	protojson.Unmarshal(marshaledjson, outputProto)

	fmt.Println("marsharedjson", marshaledjson)
	fmt.Println("outputProto", outputProto)
	fmt.Println(outputProto.GetInner()[0].Content)
	fmt.Println(outputProto.GetInner()[1].Content)


	/* protojson enum unmarsal example */
	exampleEnum := &ecommerce.EnumTest{
		State: 1,
	}
	marshaledEnum, _ := json.Marshal(exampleEnum)
	outputEnum := &ecommerce.EnumTest{}
	protojson.Unmarshal(marshaledEnum, outputEnum)
	fmt.Println("marshaledEnum", marshaledEnum)
	fmt.Println("outputEnum", outputEnum)

	/* json enum unmarsal example */
	jsonexampleEnum := "{\"state\" : 1}"
	jsonoutputEnum := &ecommerce.EnumTest{}
	json.Unmarshal([]byte(jsonexampleEnum), jsonoutputEnum)
	fmt.Println("jsonexampleEnum", jsonexampleEnum)
	fmt.Println("jsonoutputEnum", jsonoutputEnum)

	/* protojson oneof unmarsal example */
	exampleOneof := "{\"id\" : 1}"
	outputOneof := &ecommerce.OneofTest{}
	protojson.Unmarshal([]byte(exampleOneof), outputOneof)
	fmt.Println("marshaledOneof", exampleOneof)
	fmt.Println("outputOneof", outputOneof)

	/* protojson oneof unmarsal example */
	jsonexampleOneof := "{\"id\" : 1}"
	jsonoutputOneof := &ecommerce.OneofTest{}
	json.Unmarshal([]byte(jsonexampleOneof), jsonoutputOneof)
	fmt.Println("jsonmarshaledOneof", jsonexampleOneof)
	fmt.Println("jsonoutputOneof", jsonoutputOneof)

	/* protojson prototest unmarsal example */
	exampleProtoTest := "{\"is\" : false, \"num\": 0, \"id\": 1}"
	outputProtoTest := &ecommerce.ProtoTest{}
	protojson.Unmarshal(nil, outputProtoTest)
	fmt.Println("exampleProtoTest", exampleProtoTest)
	fmt.Println("outputProtoTest", outputProtoTest)
	fmt.Println("outputProtoTest", outputProtoTest.Is,outputProtoTest.Type, outputProtoTest.Num)

	/* protojson Repeatest unmarsal example */
	exampleRepeatTest := "{\"a\", \"b\", \"c\"}"
	outputRepeatTest := &ecommerce.RepeatTest{}
	protojson.Unmarshal([]byte(exampleRepeatTest), outputRepeatTest)
	fmt.Println("exampleRepeatTest", exampleRepeatTest)
	fmt.Println("outputRepeatTest", outputRepeatTest)

	var a []string
	a = nil

	b , _:= json.Marshal(a)

	fmt.Println("json marshal null string: ", len(a), len(b))

}
