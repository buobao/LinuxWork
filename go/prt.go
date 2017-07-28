package main

import "fmt"

/**
 * 整数i=[10]，指针pInt=[0x184000c0]，指针指向*pInt=[10]
 * 整数i=[3]，指针pInt=[0x184000c0]，指针指向*pInt=[3]
 * 整数i=[5]，指针pInt=[0x184000c0]，指针指向*pInt=[5]
 *
 * Wild的数组指针： <nil>
 * Wild的数组指针==nil[true]
 *
 * New分配的数组指针： &[]
 * New分配的数组指针[0x18443010]，长度[0]
 * New分配的数组指针==nil[false]
 * New分配的数组指针Make后： &[0 0 0 0 0 0 0 0 0 0]
 * New分配的数组元素[3]： 23
 *
 * Make分配的数组引用： [0 0 0 0 0 0 0 0 0 0]
 */
func main() {
	testPointer()
	testMemAllocate()
}

func testPointer() {
	var i int = 10
	var pInt *int = &i
	fmt.Printf("整数i=[%d]，指针pInt=[%p]，指针指向*pInt=[%d]\n",
		i, pInt, *pInt)

	*pInt = 3
	fmt.Printf("整数i=[%d]，指针pInt=[%p]，指针指向*pInt=[%d]\n",
		i, pInt, *pInt)

	i = 5
	fmt.Printf("整数i=[%d]，指针pInt=[%p]，指针指向*pInt=[%d]\n",
		i, pInt, *pInt)
}

func testMemAllocate() {
	var pNil *[]int
	fmt.Println("Wild的数组指针：", pNil)
	fmt.Printf("Wild的数组指针==nil[%t]\n", pNil == nil)

	var p *[]int = new([]int)
	fmt.Println("New分配的数组指针：", p)
	fmt.Printf("New分配的数组指针[%p]，长度[%d]\n", p, len(*p))
	fmt.Printf("New分配的数组指针==nil[%t]\n", p == nil)

	//Error occurred
	//(*p)[3] = 23

	*p = make([]int, 10)
	fmt.Println("New分配的数组指针Make后：", p)
	(*p)[3] = 23
	fmt.Println("New分配的数组元素[3]：", (*p)[3])

	var v []int = make([]int, 10)
	fmt.Println("Make分配的数组引用：", v)
}
