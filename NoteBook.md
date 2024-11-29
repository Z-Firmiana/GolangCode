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
