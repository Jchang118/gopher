切片结构：
    type slice struct {
        array unsafe.Pointer
        len int
        cap int
    }

切片声明：
    var s[]int
    s := make([]int, 5, 10)
