package template

import "fmt"

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// CookeMenu
// @Description: 类似于一个抽象类
type CookeMenu struct {
}

// fire
// @Description:
// @receiver CookeMenu
func (CookeMenu) fire() {
	fmt.Println("fire...")
}

// cooke
// @Description: 做菜，交给具体的子类实现
// @receiver CookeMenu
func (CookeMenu) cooke() {

}

// outfire
// @Description:
// @receiver CookeMenu
func (CookeMenu) outfire() {
	fmt.Println("out of fire...")
}

// doCook
// @Description: 封装具体步骤
// @param cook
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

// XiHongShi
// @Description: 相当于继承了 CookeMenu 而 CookMenu 实现了 Cooker 接口
// 也就是 XiHongShi 实现了 Cooker 接口
type XiHongShi struct {
	CookeMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("cook the tomato...")
}

type ChaoJiDan struct {
	CookeMenu
}

func (ChaoJiDan) cooke() {
	fmt.Println("cooke for ChaoJiDan...")
}
