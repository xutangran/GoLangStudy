/***************************************************************包******************************************************************/
package main	//指定了某文件属于的包。它应该放在每一个源文件的第一行。
//所有可执行的 Go 程序都必须包含一个 main 函数。这个函数是程序运行的入口。main 函数应该放置于 main 包中



import "fmt"	//导入已经存在的包
//创建自定义的包	----属于某一个包的源文件都应该放置于一个单独命名的文件夹里。按照 Go 的惯例，应该用包名命名该文件夹。
//导入自定义包		----导入自定义包的语法为 import 包路径。我们必须指定自定义包相对于工作区内 src 文件夹的相对路径
//导出名字			----在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。
//init 函数			----所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。
func init() {  
	//init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性。
	/*包的初始化顺序如下：
		1.首先初始化包级别（Package Level）的变量
		2.紧接着调用 init 函数。包可以有多个 init 函数（在一个文件或分布于多个文件中），它们按照编译器解析它们的顺序进行调用
		3.如果一个包导入了另一个包，会先初始化被导入的包。
		4.尽管一个包可能会被导入多次，但是它只会被初始化一次。
	*/
}
/*使用空白标识符	
	导入了包，却不在代码中使用它，这在 Go 中是非法的
	然而，在程序开发的活跃阶段，又常常会先导入包，而暂不使用它。遇到这种情况就可以使用空白标识符 _。
	导入包，并使用包变量给空白标识符赋值 ，例如：var _ = rectangle.Area // 通过空白标识符使用包变量
	有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。
	例如，我们或许需要确保调用了 rectangle 包的 init 函数，而不需要在代码中使用它。这种情况也可以使用空白标识符，如下所示。
	有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。例如，我们或许需要确保调用了 rectangle 包的 init 函数，而不需要在代码中使用它。这种情况也可以使用空白标识符，如下所示。
	import (
    _ "geometry/rectangle" 
	)
	*/




/**************************************************************类型****************************************************************/
/*
基本类型：
	布尔：bool	----值为 true 或者 false
	字符串：string
	字节：byte
	数字：
		整数：int8, int16, int32, int64, int(根据不同的底层平台(Underlying Platform)，表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型)
		无符号整数：uint8, uint16, uint32, uint64, uint(根据不同的底层平台，表示 32 或 64 位无符号整型)
		浮点：float32, float64
		复数：complex64(实部和虚部都是 float32 类型的的复数), complex128(实部和虚部都是 float64 类型的的复数)
		rune:int32 的别名
*/

func main() {  
	/********************************************************变量声明*********************************************************/
	var width, height int = 10,20
	var (
        name   = "naveen"
        age    = 29
        length int
	)
	a, b := 20, 30 // 简短声明，要求 := 操作符左边的所有变量都有初始值。
	fmt.Println("width:",width,"height:",height,"name:",name,"age:",age,"length:",length,"a:",a,"b",b)

	/********************************************************常量声明*********************************************************/
	//类型推断：无类型的常量有一个与它们相关联的默认类型，并且当且仅当一行代码需要时才提供它，类型推断使用这个推断变量的类型
	/*字符串常量*/
	const hello = "Hello World"				//无类型常量字符串：可以分配给任何字符串变量
	const typedhello string = "Hello World"	//有类型常量字符串：只能给指定类型变量赋值
	/*布尔常量：布尔常量和字符串常量没有什么不同。他们是两个无类型的常量 true 和 false。字符串常量的规则适用于布尔常量*/
	/*数字常量：给变量赋值时类型由数字常量的语法决定，数字常量可以在表达式中自由混合和匹配，只有当它们被分配给变量或者在需要类型的代码中的任何地方使用时，才需要类型*/
	
	/********************************************************switch 语句*********************************************************/
	//多表达式判断
	letter := "i"
    switch letter {
    case "a", "e", "i", "o", "u": // 一个选项多个表达式
        fmt.Println("vowel")
    default:
        fmt.Println("not a vowel")
	}
	//无表达式的 switch
	num := 75
    switch { // 表达式被省略了
    case num >= 0 && num <= 50:
        fmt.Println("num is greater than 0 and less than 50")
    case num >= 51 && num <= 100:
        fmt.Println("num is greater than 51 and less than 100")
    case num >= 101:
        fmt.Println("num is greater than 100")
	}
	//Fallthrough 语句	----在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
	switch num := 25; { // num is not a constant
    case num < 50:
        fmt.Printf("%d is lesser than 50\n", num)
        fallthrough		//语句是 case 子句的最后一个语句。如果它出现在了 case 语句的中间，编译器将会报错
    case num < 100:
        fmt.Printf("%d is lesser than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is lesser than 200", num)
	}

	/********************************************************数组*********************************************************/
	//数组声明：组的表示形式为 [n]T。n 表示数组中元素的数量，T 代表每个元素的类型。元素的数量 n 也是该类型的一部分
	var arr[3]int	//数组中的所有元素都被自动赋值为数组类型的零值
	arr2 := [3]int{12,24} 
	arr3 := [...]int{12, 48, 96} 
	fmt.Println("arr",arr,"arr2",arr2,"arr3",arr3)
	//数组是值类型
	arr4 := [...]string{"USA", "China", "India", "Germany", "France"}
    arr5 := arr4 // 
    arr5[0] = "Singapore"
    fmt.Println("arr4 is ", arr4)
	fmt.Println("arr5 is ", arr5) 
	//数组的长度：通过将数组作为参数传递给 len 函数，可以得到数组的长度。
	fmt.Println("arr5 len :", len(arr5)) 
	//使用range 迭代数组
	for i, v := range arr5 {//range returns both the index and value
		fmt.Printf("%d the element of a is %s\n", i, v)
	}
	//多维数组
	arr6 := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, // 这个逗号是必要的。如果省略了这个逗号，编译器就会报错
    }
	for _, v1 := range arr6 {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
	}
	
	/********************************************************切片*********************************************************/
	//切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。
	//创建切片	----带有 T 类型元素的切片由 []T 表示
	var slice[]string = arr5[1:5] //创建一个从[1]到[3]的切片
	slice2 := []int{6, 7, 8} // 创建和数组并返回切片引用
	slice3 := make([]int, 5, 5)	//使用 make 创建一个切片，func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。
	fmt.Println("slice",slice,"slice2",slice2,"slice3",slice3)

	//切片的修改	----切片自己不拥有任何数据。它只是底层数组的一种引用。对切片所做的任何修改都会反映在底层数组中。
	//切片的长度(len)	----切片中的元素数
	//切片的容量(cap)	----切片的容量是从创建切片索引开始的底层数组中元素数。(数组长度-切片开始的索引)
	//切片可以改变长度，最大长度小于容量。例：slice = slice[:cap(slice)]
	//追加切片元素(append)	----追加切片元素后，如果切片长度大于切片容量，会创建一个新数组，现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用。新切片的容量是旧切片的两倍。
	//切片的零值为nil,一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片。
	//内存优化----切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收

}

/**************************************************************函数****************************************************************/
/*函数声明
	func 函数名(参数名 参数类型,...) 返回类型 {  
    // 函数体（具体实现的功能）
	}
	函数中的参数列表和返回值并非是必须的
*/
//如果有连续若干个参数，它们的类型一致，那么我们无须一一罗列，只需在最后一个参数后添加该类型
func rectProps(length, width float64)(float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter	//Go 语言支持一个函数可以有多个返回值
}
//命名返回值
func rectProps2(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
//空白符：_ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。area, _ := rectProps(10.8, 5.6) // 第二个返回值被丢弃
