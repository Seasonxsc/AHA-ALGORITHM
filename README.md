# go啊哈！算法（Golang）

1. 一大波数正在靠近——排序
2. 栈、队列、链表
3. 枚举！很暴力
4. 万能的搜索

## 第1章 一大波数正在靠近——排序

* 桶排序

* 冒泡排序

  冒泡排序的精髓在于双重嵌套循环。如何确定边界条件成了这个算法的痛点。

  假设存在N个数，每一轮排序都是如下操作：

  索引<span style="color:red">**0~N-1**</span>的数 比较+交换

  索引<span style="color:red">**0~N-2**</span>的数 比较+交换

  索引<span style="color:red">**0~N-3**</span>的数 比较+交换

  …

  索引<span style="color:red">**0~1**</span>的数     比较+交换

  必须承认，最后一次排序一定是索引为<span style="color:red">**0~1**</span>的数进行比较和交换操作，一共做了<span style="color:red">**N-1**</span>次。故设定循环如下：

  ```go
  length := len(arr)
  for e := length - 1; e >= 1; e--{
      for i := 0; i < e; i++{
          /* 因内循环的操作均针对arr[i]与arr[i+1]，所以设定其边界为i<e */
          …
      }
  }
  ```

* 冒泡排序（结构体版）

* 快速排序

## 第2章 栈、队列、链表

+ 队列

+ 栈

  * 回文数

    需要对字符串进行操作。Go语言提供了byte与rune两种类型。

    ```go
    type rune = int32	/* 表示单个统一码代码点 */
    ```

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	verse := "云想衣裳花想容，春风拂槛露华浓"
    
    	for i, v := range []rune(verse) {
    		fmt.Printf("(%d, %c)", i, v)
    	}
        
        // Output:
        // (0, 云)(1, 想)(2, 衣)(3, 裳)(4, 花)(5, 想)(6, 容)(7, ，)(8, 春)(9, 风)(10, 拂)(11, 槛)(12, 露)	   (13, 华)(14, 浓)
    }
    ```

  * 有效的括号

    ```
    准备一个栈，用于存放扫描遇到的左括号
    从左向右扫描每一个字符{
    	如果遇到的是左括号，则入栈
    	如果遇到的是右括号{
    		把栈顶字符和当前字符比较
    		若匹配，则弹出栈顶字符，继续向前扫面
    		若不匹配，程序返回不匹配标志
    	}
    }
    当所有字符都扫描完毕，栈应当为空
    ```

+ 纸牌游戏——小猫钓鱼

  本节主要用到的数据结构为Go语言提供的slice和map

  + slice

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	var s1 []int
    
    	for i := 0; i < 9; i += 2 {
    		s1 = append(s1, i)
    	}
    	fmt.Println("s1: ", s1)
    
    	s2 := make([]int, 16)
    	s3 := make([]int, 10, 32)
    	fmt.Println("s2: ", s2)
    	fmt.Println("s3: ", s3)
    
    	fmt.Println("===复制切片:===")
    	copy(s2, s1)
    	fmt.Printf("s2 = %v\n", s2)
    
    	fmt.Println("===删除切片中指定索引的元素:===")
    	// 删除索引i位置的元素, s = append(4[:i],s[i+1:]...)
    	// 以删除索引3位置的元素为例
    	fmt.Printf("被删除的元素s1[3] = %d\n", s1[3])
    	s1 = append(s1[:3], s1[4:]...)
    	fmt.Printf("s1 = %v\n", s1)
    
    	fmt.Println("===删除切片首元素:===")
    	front, s1 := s1[0], s1[1:]
    	fmt.Printf("front = %d\n", front)
    	fmt.Printf("s1 = %v\n", s1)
    
    	fmt.Println("===删除切片尾元素:===")
    	tail, s1 := s1[len(s1)-1], s1[:len(s1)-1]
    	fmt.Printf("tail = %d\n", tail)
    	fmt.Printf("s1 = %v\n", s1)
        
        // Output:
        // s1:  [0 2 4 6 8]
    	// s2:  [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    	// s3:  [0 0 0 0 0 0 0 0 0 0]
    	// ===复制切片:===
    	// s2 = [0 2 4 6 8 0 0 0 0 0 0 0 0 0 0 0]
    	// ===删除切片中指定索引的元素:===
    	// 被删除的元素s1[3] = 6
    	// s1 = [0 2 4 8]
    	// ===删除切片首元素:===
    	// front = 0
    	// s1 = [2 4 8]
    	// ===删除切片尾元素:===
    	// tail = 8
    	// s1 = [2 4]
    }
    ```

  + map

+ 链表

  ```go
  package main
  
  import (
  	"container/list"
  	"fmt"
  )
  
  func main() {
  	l := list.New()
  	e4 := l.PushBack(4)
  	e1 := l.PushFront(1)
  	l.InsertBefore(3, e4)
  	l.InsertAfter(2, e1)
      
  	for e := l.Front(); e != nil; e = e.Next() {
  		fmt.Println(e.Value)
  	}
  	
      // Output:
      // 1
      // 2
      // 3
      // 4
  }
  ```

  

+ 模拟链表

