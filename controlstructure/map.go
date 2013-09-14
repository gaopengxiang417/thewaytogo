package main

import (
	"fmt"
	"sort"
)

func main() {
	//初始化一个map
	var mapfirst map[string]int
	fmt.Printf("the type %T , value is %v\n", mapfirst, mapfirst)
	fmt.Printf("len is %d\n", len(mapfirst))
	fmt.Println(mapfirst)

	fmt.Println(mapfirst["china"])
	//array slice struct不能为key
	var arr = [12]int{}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)

	var slic = new([]int)
	fmt.Println(slic)
	fmt.Printf("%T\n", slic)

	var slic2 = make([]int, 10)
	fmt.Println(slic2)
	fmt.Printf("%T\n", slic2)

	map_create()

	fmt.Println("*******华丽的分割线***********")

	maptest := new(map[string]int)
	fmt.Println(maptest)
	fmt.Printf("the value %T,%v", *maptest, *maptest)
	fmt.Println(*maptest)
	// *maptest["fir"] = 44

	map_value_test()

	map_capcity()

	map_slice_value()

	map_check_isexist()

	map_for_range()

	map_slice_array()

	map_sort()

	map_reverse()
}

//反转一个map，可以作为value，value作为key
func map_reverse() {
	for key, value := range barmap {
		fmt.Printf("%s:%d\t", key, value)
	}
	fmt.Println()

	//进行反转
	var reversemap = make(map[int]string, len(barmap))

	for key, value := range barmap {
		reversemap[value] = key
	}

	for key, value := range reversemap {
		fmt.Printf("%d:%s\t", key, value)
	}
}

func map_sort() {
	fmt.Println("unsorted......")
	for key, value := range barmap {
		fmt.Printf("%s : %d,", key, value)
	}
	keys := make([]string, len(barmap))
	i := 0
	for key, _ := range barmap {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted.......")
	for _, key := range keys {
		fmt.Printf("%s: %d,", key, barmap[key])
	}
	fmt.Println()
}

//构造一个map的slice，注意两种不同的循环方式
//这里为什么出现这种情况，因为对于for indx, value的循环方式，其实value是原始
//内容的拷贝，不是引用，所以会导致这个问题
func map_slice_array() {
	//version one
	var mapslice = make([]map[int]int, 5)

	for i := range mapslice {
		mapslice[i] = make(map[int]int)
		mapslice[i][1] = 34
	}

	fmt.Printf("version 1 the value is %v\n", mapslice)

	//version two
	var mapsliceversion2 = make([]map[int]int, 5)

	for _, value := range mapsliceversion2 {
		value = make(map[int]int)
		value[1] = 434
	}
	fmt.Printf("version 2 the value is %v\n", mapsliceversion2)
}

//map里面存放的是无序的，所以如果需要排序，那么应该拿出来以后在排序
var (
	barmap = map[string]int{"alpla": 23, "bravo": 29, "charli": 44, "delta": 22, "echo": 99,
		"firefox": 1, "golf": 93, "hotel": 78, "indil": 884, "jutl": 2323, "kilo": 12121, "limi": 12}
)

//对map进行循环，但是循环的顺序是不固定的
func map_for_range() {
	map1 := make(map[string]float32)

	map1["first"] = 23.3
	map1["second"] = 54.2232
	map1["third"] = 22.121212
	map1["one"] = 554

	for key, value := range map1 {
		fmt.Printf("the key is : %s, value is %f\n", key, value)
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome",
		"Japan": "Tokyo"}
	for key, value := range capitals {
		fmt.Println(key, ":", value)
	}
}

//判断一个key是否在map中,删除map中的记录
func map_check_isexist() {
	var intmap map[string]int = make(map[string]int)
	intmap["one"] = 0
	intmap["two"] = 2
	intmap["three"] = 3

	//判断是否存在
	value, ok := intmap["one"]
	fmt.Println(ok, value)
	value2, ok2 := intmap["notfound"]
	fmt.Println(ok2, value2)

	var val int
	var ispresent bool
	map1 := make(map[string]int)

	map1["china"] = 55
	map1["USA"] = 30
	map1["jap"] = 99

	val, ispresent = map1["china"]
	if ispresent {
		fmt.Printf("the value \"china\" is %d\n", val)
	} else {
		fmt.Println("map1 does not contain china")
	}

	val, ispresent = map1["uus"]
	fmt.Printf("is exist the key \"uus\" %t\n", ispresent)
	fmt.Printf("the value is %d\n", val)

	//删除一个记录
	delete(map1, "china")
	val, ispresent = map1["china"]
	if ispresent {
		fmt.Println("delete is not success")
	} else {
		fmt.Println("success delete chian")
	}
}

//map中的value可以为slice类型,也可以是array类型
func map_slice_value() {
	var maptt = make(map[string][3]int)
	var array [3]int
	maptt["one"] = array
	fmt.Printf("the type is %T,value is %v\n", maptt, maptt)

	var maptest = make(map[string][]int)
	maptest["one"] = []int{23, 44}
	fmt.Printf("the type is %T,value is %v\n", maptest, maptest)

	var maptesttwo = make(map[string]*[]int)
	var intslice = new([]int)
	maptesttwo["two"] = intslice
	fmt.Printf("the type is %T,the value is %v\n", maptesttwo, maptesttwo)
}

//map的capcity，主要是给map的一个初始容量
func map_capcity() {
	var maptest = make(map[string]float32, 40)
	fmt.Println(maptest)
	fmt.Println(len(maptest))
	//map不支持cap函数
	// fmt.Println(cap(maptest))
}

//map里面的key只能是int，float,string等，不能是map,array,slice,struct
//map里面的value可以是任意类型，包括函数
func map_value_test() {
	valuemap := map[int]func() int{
		1: func() int { return 10 },
		3: func() int { return 20 },
		8: func() int { return 90 }, //注意不能少了这个逗号
	}

	fmt.Println(valuemap)
	fmt.Printf("the type is %T,value is %v", valuemap, valuemap)
}

//演示创建map的功能
func map_create() {
	var mapone map[string]int
	var maptwo map[string]int

	mapone = map[string]int{"one": 4, "two": 5}
	mapcreateed := make(map[string]float32)

	mapcreateed["mapone"] = 432.23
	mapcreateed["maptwo"] = 89.21323

	maptwo = mapone

	maptwo["third"] = 9

	fmt.Println("the length of mapone: ", len(mapone))
	fmt.Println("the length of maptwo: ", len(maptwo))
	fmt.Println("the length of mapthree :", len(mapcreateed))

	fmt.Println("map one :", mapone)
	fmt.Println("mapone's one is ", mapone["one"])

	fmt.Println("map two:", maptwo)
	fmt.Println("maptwo's tow is ", maptwo["two"])

	fmt.Println("map created :", mapcreateed)
}
