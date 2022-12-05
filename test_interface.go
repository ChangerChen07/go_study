package main

import "fmt"

type USB interface {
	read()
	write()
}
type Computer struct {
	name string
}
type Mobile struct {
	model string
}

func (pc Computer) read() {
	fmt.Printf("pc.name: %v\n", pc.name)
	fmt.Println("pc read.......")
}
func (pc Computer) write() {
	fmt.Printf("pc.name: %v\n", pc.name)
	fmt.Println("pc write.......")
}

func (f Mobile) read() {
	fmt.Printf("f.model: %v\n", f.model)
	fmt.Println("mobile read......")
}
func (f Mobile) write() {
	fmt.Printf("f.model: %v\n", f.model)
	fmt.Println("mobile write......")
}

func main() {
	var pc_usb USB
	pc_usb = Computer{name: "联想"}
	pc_usb.read()
	pc_usb.write()
	var m_usb USB = Mobile{model: "5G mobile"}
	m_usb.read()
	m_usb.write()
}
