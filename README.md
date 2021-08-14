# 啊哈！算法（Golang）

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



