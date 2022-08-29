package main

import(
    "fmt"
)

func main(){
    fmt.Println("master second")
    fmt.Println("hot-fix 33333")
    fmt.Println("master") 
    fmt.Println("hot-fix 2222")
    fmt.Println("通过远程进行代码修改")
    //乘法表
    for i:=1;i<9;i++{
        for j:=1;j<=i;j++{
            fmt.Printf("%d  ",i*j)
        }
        fmt.Println()
    }
}
