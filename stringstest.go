package main

import (
	"fmt"
	"strconv"
	"strings"
)

//strings的学习和使用
func main() {
	//判断字符串是否以某个字符串开头
	var str = "This is a example of string"
	fmt.Printf("T/F, does the string \"%s has prefix %s", str, "This")
	fmt.Printf("\t %t\n ", strings.HasPrefix(str, "This"))
	//判断字符串是否已某个字符串结尾
	fmt.Printf("T/F,does the string \"%s has suffix %s", str, "strings")
	fmt.Printf("\t%t\n", strings.HasSuffix(str, "strings"))

	//判断一个字符串是否包含另一个字符串
	var str2 = "second"
	var str3 = "example"
	var second = strings.Contains(str, str2)
	var third = strings.Contains(str, str3)
	fmt.Printf("%t....%T\n", second, second)
	fmt.Printf("%t....%T\n", third, third)

	//index 使用,判断一个字符串在另一个字符串中的位置
	var str4 = "is"
	fmt.Printf("the %s index is %d :\n", str4, strings.Index(str, str4))
	fmt.Printf("the last %s index is %d :\n", str4, strings.LastIndex(str, str4))

	//replace的使用
	var str5 = strings.Replace(str, "is", "IS", -1)
	fmt.Printf("the origin  is :%s\n", str)
	fmt.Printf("the replace result is :%s\n", str5)

	//计算某个字符串出现的次数
	fmt.Printf("the count is %d\n", strings.Count(str, str4))
	var str6 = "gggggggggg"
	fmt.Printf("the count is %d\n", strings.Count(str6, "gg"))

	//repleat的使用,拷贝字符串的多次
	var str7 = "origin!"
	fmt.Printf("the result of repleat is :%s\n", strings.Repeat(str7, 5))

	//转换字符串的大小写
	var str8 = "Hello , this Is a capital and lower casE"
	var lowercase string
	var uppercase string
	fmt.Printf("the origin string is : %s\n", str8)

	lowercase = strings.ToLower(str8)
	fmt.Printf("the lower string is : %s\n", lowercase)

	uppercase = strings.ToUpper(str8)
	fmt.Printf("the upper string is : %s\n", uppercase)

	//截取字符串开头或者结尾的空格
	var str9 = " hello this is a space case     "
	fmt.Printf("the space result is : %s\n", strings.TrimSpace(str9))

	//截取开头或者结尾的指定字符串
	var str10 = "reduce some string both start and end end"
	fmt.Printf("the speical string cut result is : %s\n", strings.Trim(str10, "end"))

	//截取开头或者结尾的指定字符串
	fmt.Printf("the reduce start result is : %s\n", strings.TrimLeft(str10, "red"))
	fmt.Printf("the reduce end result is : %s\n", strings.TrimRight(str10, "nd"))

	//split的使用
	//根据空格进行分割
	var str11 []string = strings.Fields(str10)
	for i := 0; i < len(str11); i++ {
		fmt.Printf("%s \t", str11[i])
	}

	fmt.Println()

	//按照某个指定的字符串进行分割
	var str12 string = "this is , a split by, comma, "
	var str13 []string = strings.Split(str12, ",")
	for i := 0; i < len(str13); i++ {
		fmt.Printf("%s \t", str13[i])
	}

	fmt.Println()
	//splitafter
	var str14 string = "this is , a split by, comma, "
	var str15 []string = strings.SplitAfter(str14, ",")
	for i := 0; i < len(str15); i++ {
		fmt.Printf("%s \t", str15[i])
	}

	fmt.Println()

	//join
	var str16 = "hi quick jump from the wall"
	str17 := strings.Fields(str16)
	fmt.Printf("field reuslt :%v", str17)
	fmt.Println()
	for _, value := range str17 {
		fmt.Printf("iterator reuslt :%s\t", value)
	}

	var str18 = "hello|siple test|china he"
	str19 := strings.Split(str18, "|")
	fmt.Printf("field reuslt :%v", str19)
	fmt.Println()
	for _, value := range str19 {
		fmt.Printf("iterator reuslt :%s\t", value)
	}
	fmt.Println()

	//join
	str20 := strings.Join(str19, "!!")
	fmt.Printf("final join is :%s\n", str20)

	//reader
	var str21 = "this is a reader test case"
	reader := strings.NewReader(str21)
	var bytes []byte = make([]byte, 1024)
	count, err := reader.Read(bytes)
	if err != nil {
		fmt.Printf("total read count : %d\n", count)
		fmt.Println(bytes)
	} else {
		fmt.Println("error:", err)
	}

	fmt.Println(strings.NewReader(str21).ReadRune())

	//将其他类型传化为字符串
	//输出现在平台的int类型

	fmt.Println("current int size : ", strconv.IntSize)
	fmt.Println("int to string :", strconv.Itoa(89))
	fmt.Println("float to string :", strconv.FormatFloat(8792.2323, 'f', 9, 64))

	//从字符串转化为数字
	value, err := strconv.Atoi("987")
	fmt.Println("from string to int:", value)

	value1, _ := strconv.ParseFloat("8877.2323232", 32)
	fmt.Println("from string to float:", value1)

	SDSSD
}
