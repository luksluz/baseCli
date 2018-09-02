package main

type Header struct {
	Name string
}

type Opts1 struct {
	
}

func SetHeader(name string) {
	h := Header{}
	h.Name = name
}
