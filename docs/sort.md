<!-- MDTOC maxdepth:6 firsth1:1 numbering:0 flatten:0 bullets:1 updateOnSave:1 -->

- [sort](#sort)   
   - [type Interface](#type-interface)   
   - [func Sort(data Interface)](#func-sortdata-interface)   
   - [func Stable(data Interface)](#func-stabledata-interface)   
   - [排序算法](#排序算法)   
      - [func insertionSort(data Interface, a, b int)](#func-insertionsortdata-interface-a-b-int)   
      - [func quickSort(data Interface, a, b, maxDepth int)](#func-quicksortdata-interface-a-b-maxdepth-int)   
      - [func heapSort(data Interface, a, b int)](#func-heapsortdata-interface-a-b-int)   
   - [func IsSorted(data Interface) bool](#func-issorteddata-interface-bool)   
   - [func Reverse(data Interface) Interface](#func-reversedata-interface-interface)   

<!-- /MDTOC -->

# sort

```
import "sort"
```

* sort包提供了排序切片和用户自定义数据集的函数
* 必须要对data实现Interface接口才能排序

## type Interface

```
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
```

* **只有满足sort.Interface接口的（集合）类型才可以被本包的函数进行排序。**
* 方法要求集合中的元素可以被整数索引。

## func Sort(data Interface)

* Sort排序data，从大到小顺序
* 它调用1次data.Len确定长度，调用O(n*log(n))次data.Less和data.Swap
* 本函数不能保证排序的稳定性（即不保证相等元素的相对次序不变）

## func Stable(data Interface)

* Stable排序data，并保证排序的稳定性，相等元素的相对次序不变。
* 它调用1次data.Len，O(n*log(n))次data.Less和O(n*log(n)*log(n))次data.Swap。


## 排序算法

### func insertionSort(data Interface, a, b int)
### func quickSort(data Interface, a, b, maxDepth int)
### func heapSort(data Interface, a, b int)



## func IsSorted(data Interface) bool

* data是否有序

## func Reverse(data Interface) Interface

Reverse包装一个Interface接口并返回一个新的Interface接口，对该接口排序可生成递减序列。

```
s := []int{5, 2, 6, 3, 1, 4} // unsorted
sort.Sort(sort.Reverse(sort.IntSlice(s)))
fmt.Println(s) // [6 5 4 3 2 1]
```










---
