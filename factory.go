package main

import "fmt"

type AbstractData interface {
	toString() string
}

type AbstractDatabase interface {
	getData() AbstractData
}

type PostgreData struct {
}

func (cpa PostgreData) toString() string {
	return "It is PostgreData"
}

type PostgreDatabase struct {
}

func (cfa PostgreDatabase) getData() AbstractData {
	return PostgreData{}
}

type MongoData struct {
}

func (cpb MongoData) toString() string {
	return "It is MongoData"
}

type MongoDatabase struct {
}

func (cfb MongoDatabase) getData() AbstractData {
	return MongoData{}
}

func main() {
	p := PostgreDatabase{}
	m := MongoDatabase{}

	pd := p.getData()
	md := m.getData()

	fmt.Println(pd.toString())
	fmt.Println(md.toString())
}
