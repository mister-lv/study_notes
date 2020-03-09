# go语言编程

## 第2章 
### 顺序编程
1. 变量 
    - 1.1
        - 变量声明1
            ```
                var v1 int
                var v2 string
                var v3 [10]int
                var v4 []int
                var v5 struct {
                    f int
                }
                var v6 *int
                var v7 map[string]int
                var v8 func(a int) int
            ```
    - 1.2
        - 变量声明2
            ```
            var (
                v1 int
                v2 string    
            )
            ```
    - 1.3. 变量初始化
        ```
            var v1 int = 10
            var v2 = 10 // 编译器可以自动推导出v2的类型
            v3 := 10   // 编译器可以自动推导出v3类型
            注意： 除了第三种声明方式不能用于声明全局变量
        ```
    - 1.4. 变量赋值
        - 变量初始化和变量赋值是两个不同的概念
        - 多重赋值
            ```
                i, j = j, i
            ```
    - 1.5. 匿名变量
        ```
            _, _, nick := GetName()
        ```
2. 常量
        `在Go语言中，常量是指编译期间就已知且不可改变的值。常量可以是数值类型（包括整数、浮点数和复数类型）、布尔类型、字符串类型等。
        `
        `字面常量是无类型常量`
    - 2.1 字面常量
        `
        所谓字面常量（literal),是指程序中硬编码的常量
        `
        ```
            -12                 // 无类型的 它可以赋值给int、uint、 int32、 int64、 float32、 float64、 complex64、complex128等类型的变量
            3.1415926           // 浮点型的常量
            3.2 + 12i           // 复数类型的常量
            true                // 布尔类型的常量
            "foo"               // 字符串常量
        ```
    - 2.2 常量定义
        `通过const关键字，你可以给字面常量指定一个友好的名字`  
        ```
            const Pi float64 = 3.1415926
            const zero = 0.0                // 无类型浮点常量
            const (
                size int64 = 1024
                eof = -1                    // 无类型整形常量
            )
            const u, v float32 = 0, 3       // u = 0.0, v = 3.0, 常量的多重赋值
            const a, b, c = 3, 4, "foo"     // a = 3, b = 4, c = "foo", 无类型整数和字符串常量 
        ```
        `常量定义的右值也可以是一个在编译期运算的常量表达式，比如`
        ```
            const mask = 1 << 3
            // 由于常量的赋值是一个在编译期行为，所以右值不能出现在任何需要运行期才能得出结果的表达式。
            const Home = os.GetEnv("HOME")
        ```    
    - 2.3 与定义常量
        ```
            // Go 语言预定义了这些常量： true、false、iota
            // iota比较特殊，可以被认为是一个可被编译器修改的常量，在每一个const关键字出现时被重置为0，然后在下一个const出现之前，每出现一次itoa，其所代表的数字会自增1。
            const (                 // iota 被重设为 0
                c0 = iota           // c0 == 0
                c1 = iota           // c1 == 1
                c2 = iota           // c2 == 2
            )
      
            const (
                a = 1 << iota       // a == 1 (iota 在每一个const开头被重设为0)
                b = 1 << iota       // b == 2
                c = 1 << iota       // v == 4
            )
      
            const (
                u           = iota * 42     // u == 0
                v float64   = iota * 42     // v == 42.0
                w           = iota * 42     // w == 84
            )
        
            const x = iota          // x == 0 (因为iota又被重设为0了)
            const y - iota          // y == 0 (同上)
      
            const (
                c0 = iota           // iota被重设为0
                c1                  // c1 == 1
                c2                  // c2 == 2
            )
      
            const (
                a = 1 << iota       // a ==1 (iota 在每个const开头被重设为0)
                b                   // b == 2
                c                   // c == 4
            )
        ```
    - 2.4 枚举
        ```
            枚举指一系列相关的常量，Go语言并不支持众多其他语言支持的enum关键字
            const (
                SUNDAY = iota
                Monday
                Tuesday
                Wednesday
                Thursday
                Friday
                Saturday
                numberOfDays        // 这个常量没有导出
            )
            注意：同Go语言的其他符号(symbol)一样，以大写字母开头的常量在包外可见
        ```  
       
            
            