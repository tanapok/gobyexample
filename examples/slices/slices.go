// _Slice_ 是 Go 中一个重要的数据类型，它提供了比数组更强大的序列交互方式。

package main

import "fmt"

func main() {

	// 与数组不同，切片的类型仅由它们包含的元素（而不是元素的数量）决定。
	// 未初始化的切片等于 nil，长度为 0。
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// 要创建一个非零长度的空切片，使用内置的 `make`。
	// 在这里，我们创建了一个长度为 3 的字符串切片（初始值为零值）。
	// 默认情况下，新切片的容量等于它的长度；
	// 如果我们提前知道切片将增长，则可以明确地将容量作为附加参数传递给 `make`。
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// 我们可以像数组一样赋值或取值。
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` 返回切片的长度。
	fmt.Println("len:", len(s))

	// 除了这些基本操作外，切片还支持比数组更丰富的操作。
	// 如，切片支持内建函数 `append`，该函数会返回一个包含了一个或者多个新值的切片。
	// 注意，由于 `append` 可能返回一个新的切片，因此我们需要接收其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片也可以使用 `copy` 复制。
	// 这里我们创建一个与 `s` 长度相同的空切片 `c`，然后将 `s` 复制给 `c`。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片支持语法为 `slice[low:high]` 的“切片”运算符。
	// 例如，右边的操作可以得到一个包含元素 `s[2]`、`s[3]` 和 `s[4]` 的切片。
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个切片包含 `s[5]` 之前的元素（不包含 `s[5]`）。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个切片包含 `s[2]` 及其之后的元素（包含 `s[2]`）。
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中声明并初始化一个切片变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 切片可以组成多维数据结构。
	// 与多维数组不同，内部切片的长度可以变化。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
