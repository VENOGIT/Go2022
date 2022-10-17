package main

import "fmt"

/*
平铺式设计
那么作为interface数据类型，他存在的意义在哪呢？实际上是为了满足一些面向对象的编程思想。我们知道，软件设计的最高目标就是 高内聚，低耦合。
那么其中一个设计原则叫 开闭原则
什么是开闭原则呢？
看个例子
问题：
	代码很简单，就是一个银行业务员，他可能拥有很多的业务，比如save，transfer，pay支付等
那么如果这个业务员摩卡只有这个几个方法还好，但是随着我们的程序越来越复杂，银行业务员可能要增加方法，会导致业务员摩卡越来越臃肿
优化：
	这样的设计会导致，当我们去给Banker添加新的业务的时候，会直接修改原有的Banker代码，那么Banker模块的功能会越来越多，出现问题的几率也就越来越大
假如此时Banker已经有99个业务了，现在我们要添加第100个业务，可能由于一次的不小心，导致之前99个业务也一起崩溃，因为所有的业务你都在一个Banker类里，
他们的耦合度太高，Banker的职业也不够单一，代码的维护成本随着业务的复杂正比成倍增大。

*/

// Banker 我们要写一个类， Banker银行业务员
type Banker struct {
}

// Save 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务 。。。")
}

// Transfer 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}

// Pay 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

func main() {
	banker := &Banker{}
	banker.Save()
	banker.Transfer()
	banker.Pay()
}
