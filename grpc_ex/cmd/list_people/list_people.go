package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/proto"

	pb "grpc_ex/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
)

// 与えられたio.Writer(出力先)にpb.Person構造体の情報を書き込むためのヘルパー関数
func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID:", p.Id)
	fmt.Fprintln(w, "Name:", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "E-mail address:", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_MOBILE:
			fmt.Fprint(w, "Mobile phone #: ")
		case pb.Person_HOME:
			fmt.Fprint(w, "Home phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, "Mobile phone #: ")
		}

		fmt.Fprintln(w, pn.Number)

	}
}

// 　与えられたio.Writerにpb.AddressBook構造体内の全てのpb.Person情報をリストするための関数
func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

// プログラムのエントリーポイント
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	// [END unmarshal_proto]

	listPeople(os.Stdout, book)
}
