package main

import (
	"fmt"
	"io/ioutil"
	"log"
	simple "main/src1/simple"
	complex "main/src1/complex"
	enum "main/src1/enum"
	practice "main/src1/practice"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// sm := doSimple()

	// smAsString := toJSON(sm)
	// fmt.Println(smAsString)

	// sm2 := &simple.MyMessage{}
	// fromJSON(smAsString, sm2)
	// fmt.Println("sm2: ", sm2)
	// // readAndWriteDemo(sm)
	// doComplex()
	// doEnum()

	doPractice()
}



func toJSON(pb proto.Message) string {
	out, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return string(out)
}

func fromJSON(in string, pb proto.Message) error{
	err := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal([]byte(in), pb)
	return err
}

func readAndWriteDemo(sm proto.Message) {
	err := writeToFile("simple.bin", sm)
	if err != nil {
		fmt.Println("err3: ", err)
	}

	sm2 := &simple.MyMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func writeToFile(name string, pb proto.Message) error{
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("err1: %s", err)
		return err 
	}

	if err :=ioutil.WriteFile(name, out, 0644); err != nil {
		log.Fatalln("err2: %s", err)
		return err
	}
	fmt.Println("write done")
	return nil
}

func readFromFile(fname string, pb proto.Message) error{
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("err3: %s", err)
		return err
	}
	err = proto.Unmarshal(in ,pb)
	if err != nil {
		log.Fatalln("err4: ", err)
		return err
	}
	return nil
}

func doSimple() *simple.MyMessage{
	mess := simple.MyMessage{
		Id: 12345,
		FirsName: "alo",
		IsValidiated: true,
	}
	fmt.Println("mess: ", mess.GetId())
	return &mess
}

func doEnum() {
	en := enum.EnumMessage{
		Id: 42,
		DayOfTheWeek: enum.DayOfTheWeek_MONDAY,
	}
	fmt.Println("en: ", en)
}

func doPractice() {
	person := practice.Person {
		Name:	"alo",
		Id: 32,
		Email: "1234@gmail.com",
		Phones: []*practice.Person_PhoneNumber{
			&practice.Person_PhoneNumber{
				Number:	"2143124",
				Type: practice.Person_HOME,
			},
		},
	}
	fmt.Println("person: ", person)
}

func doComplex() {
	cm := complex.ComplexMessage{
		OneDummy: &complex.DummyMessage{
			Id:		1,
			Name:	"First message",
		},
		MultipleDummy: []*complex.DummyMessage{
			&complex.DummyMessage{
				Id:	2,
				Name: "Second message",
			},
			&complex.DummyMessage{
				Id: 3,
				Name: "Third message",
			},
		},
	}
	fmt.Println("a: ", cm)
}