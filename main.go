package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func a() {
	fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))
}

type Opts struct {
	Header string
	Itens  []Opt
}

type Opt struct {
	Name     string
	Callback interface{}
}

func (o *Opts) SetHeader(header string) {
	o.Header = header
}

func (o *Opts) NewItem(name string, callback interface{}) {
	item := Opt{Name: name}
	item.Callback = callback
	o.Itens = append(o.Itens, item)
}

func clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}

func ShowOpts(opts *Opts) {
	clear()
	fmt.Printf("[.]  %s\n", strings.ToUpper(opts.Header))

	itensCount := len(opts.Itens)

	for index, opt := range opts.Itens {
		fmt.Printf("[%d]  %s\n", index+1, opt.Name)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("_")
	result, _ := reader.ReadString('\r')
	result = strings.Replace(result, "\r", "", -1)
	if result == "q" || result == "exit" {
		os.Exit(0)
	}
	checkResult, err := strconv.Atoi(result)
	if err != nil {
		log.Println(err.Error())
		log.Println("Error: no number")
		return
	}
	if checkResult <= 0 && checkResult > itensCount {
		fmt.Println("Error: number out of range")
		return
	}
	fmt.Printf("\r Item escolhido: [%s]", opts.Itens[checkResult-1].Name)
	choosenItem := opts.Itens[checkResult-1].Callback
	if reflect.TypeOf(choosenItem).String() == "*main.Opts" {
		time.Sleep(1200 * time.Millisecond)
		ShowOpts(choosenItem.(*Opts))
	}
}

func main() {
	opt1 := Opts{Header: "Times"}
	opt1.NewItem("Novo", &opt1)
	opt1.NewItem("Editar", &opt1)
	opt1.NewItem("Deletar", &opt1)
	opt1.NewItem("Mostrar", &opt1)
	ShowOpts(&opt1)
}
