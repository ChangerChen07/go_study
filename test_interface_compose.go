/* package main

import "fmt"

type Flyer interface {
	fly()
}

type Swimer interface {
	swim()
}

// 接口嵌套组合
type FlyFisher interface {
	Flyer
	Swimer
}

type FlyFish struct {
	name string
}

func (f *FlyFish) fly() {
	fmt.Println("我是飞鱼，能飞的鱼")
}

func (f *FlyFish) swim() {
	fmt.Println("我是飞鱼，能游泳！！")
}

func main() {
	var ff FlyFisher = new(FlyFish)
	// ff.name = "飞鱼1"
	f := FlyFish{name: "飞鱼1"}
	f.fly()

	ff.fly()
	ff.swim()

}
*/