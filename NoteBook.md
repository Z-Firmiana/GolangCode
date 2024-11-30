# Golang笔记

## 其他类型转换为字符串

- `fmt.Sprintf()`：int为%d；float为%f；bool为%t；byte为%c
- `strconv`：可以使用此包实现字符串转化为数值类型

## 运算符

- `i++`和`i--`是单独语句，只能单独使用，不能放在赋值运算符中

## 流程控制

- GO语言中没有`while`语句
- `fallthrough`穿透，可以执行满足条件的case的下一个case
- 在多次循环中，可以使用标号`label`标出想`break`的循环
- `continue`语句可以结束当前循环，开始下一次的循环迭代过程，仅限在`for`循环内使用。在`continue`语句后添加标签时，表示开始标签对应的循环
- `goto`语句通过标签进行代码间的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助
  
  ```go
        n := 27
        if n > 24 {
            fmt.Println("Adult")
            goto label2
        }

        fmt.Println("A")
        fmt.Println("B")
    label2:
        fmt.Println("C")
        fmt.Println("D")
  ```

## 数组

- 自动推导数组长度：`var arr = [...]int{12,34,56,79}`
- 循环（range）：

  ```go
  for k, v := range arr {
    fmt.Printf(v)
  }
  ```

- 数组是**值类型**，切片（Slice）为**引用类型**

## 切片（Slice）

- 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。非常灵活，支持**自动扩容**
- 声明：`var slice1 := []int{1, 2, 3, 4}`。即声明时不指定个数，若指定则为数组
- Go语言中声明切片以后切片的默认值就是**nil**
- 基于数组定义切片：

  ```go
  a := [5]int{1, 2, 3, 4, 5} //定义数组a
  b := a[:] //获取数组所有的值
  c := a[2:4]
  d := a[2:]
  e := a[:3]
  ```

- 基于切片再切片
- 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。使用`cap()`来获得
- `make([]T, size, cap)`：创建切片

  ```go
  var sliceA = make([]int, 4, 8)
  sliceA[0] = 10
  ```

- 切片使用`append()`给其扩容，同时也可以进行合并：
  
  ```go
  var sliceA []int
  var sliceB = []int{2, 4, 5}
  sliceA = append(sliceA, 12)
  sliceA = append(sliceA, sliceB...)
  ```

- `copy()`：复制切片
  
  ```go
  sliceA := []int{1, 2, 3, 4}
  sliceB := make([]int, 4, 4)
  copy(sliceB, sliceA) //将sliceA的值赋值给sliceB
  ```

- Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素，使用切片再切片然后使用`append()`合并
- sort排序，默认为从小到大
  
  ```go
  //升序
  sort.Ints(intSlice)
  sort.Float64s(float8Slice)
  sort.Strings(stringSlice)
  //降序
  sort.Sort(sort.Reverse (sort.IntSlice(intSlice)))
  sort.Sort(sort.Reverse(sort.Float64Slice(float8Slice)))
  sort.Sort(sort. Reverse (sort.StringSlice(stringSlice)))
  ```
