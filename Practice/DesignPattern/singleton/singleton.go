package singleton

import "sync"

type singleton struct {

}

////  ①饿汉方式实现单例
////  缺点在于实例会在包被导入的时候初始化 所有可能会导致程序加载时间较长
//var ins *singleton = &singleton{}
//
//// GetInsOr
//// @Description: 饿汉方式实现单例
//// @return *singleton
//func GetInsOr() *singleton  {
//	return ins
//}

////  ②懒汉方式实现单例
////  缺点在于非并发安全，实际使用中需要加锁
//var ins *singleton
//var mu sync.Mutex
//
//// GetInsOr
//// @Description: 懒汉方式实现单例
//// @return *singleton
//func GetInsOr() *singleton {
//	if ins == nil {
//		mu.Lock()
//		if ins == nil {
//			ins = &singleton{}
//		}
//		mu.Unlock()
//	}
//	return ins
//}

//  ③ 更优雅的方式
var ins *singleton
var once sync.Once

// GetInsOr
// @Description:使用 sync 包优雅实现
// @return *singleton
func GetInsOr() *singleton  {
	//  once.Do 方法当前仅当第一次被调用时才会执行函数，用于必须刚好运行一次的初始化
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

