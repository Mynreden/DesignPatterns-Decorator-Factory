package main

import "fmt"

type TextDecorator interface {
	format() string
}

type Document struct {
	Text string
}

func NewDocument(text string) Document {
	return Document{text}
}

func (d Document) format() string {
	return d.Text
}

type BoldDecorator struct {
	Decorator TextDecorator
}

func (bd BoldDecorator) format() string {
	return "<b>" + bd.Decorator.format() + "<b>"
}

type ItalicDecorator struct {
	Decorator TextDecorator
}

func (id ItalicDecorator) format() string {
	return "<i>" + id.Decorator.format() + "<i>"
}

type UnderlineDecorator struct {
	Decorator TextDecorator
}

func (ud UnderlineDecorator) format() string {
	return "<u>" + ud.Decorator.format() + "<u>"
}

func mains() {
	doc := NewDocument("Sultan SE-2201")
	bold := BoldDecorator{doc}
	underline := UnderlineDecorator{bold}
	italic := ItalicDecorator{underline}
	fmt.Println(italic.format())

	doc = NewDocument("Example 2")
	italic = ItalicDecorator{doc}
	bold = BoldDecorator{italic}
	fmt.Println(bold.format())

}
