package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const arryLen = 500

func randArray(l int) []int {
	a := make([]int, l)
	for i := 0; i < l; i++ {
		a[i] = rand.Intn(l)
	}

	return a
}

/*
func BenchmarkRandArry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randArray(arryLen)
	}
}
*/

func check_sort(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

//冒泡排序
func TestBubbleSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		BubbleSort(a, len(a))
		if !check_sort(a) {
			t.Fatalf("BubbleSort error\n")
		}
	}

	fmt.Println("BubbleSort Pass")
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		BubbleSort(a, len(a))
	}
}

//插入排序
func TestInsertSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		InsertSort(a, len(a))
		if !check_sort(a) {
			t.Fatalf("InsertSort error\n")
		}
	}
	fmt.Println("InsertSort Pass")
}

func BenchmarkInsertSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		InsertSort(a, len(a))
	}
}

//选择排序
func TestSelectSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		SelectSort(a, len(a))
		if !check_sort(a) {
			t.Fatalf("SelectSort error\n")
		}
	}
	fmt.Println("SelectSort Pass")
}

func BenchmarkSelectSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		SelectSort(a, len(a))
	}
}

//希尔排序
func TestShellSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		ShellSort(a, len(a))
		if !check_sort(a) {
			t.Fatalf("ShellSort error\n")
		}
	}
	fmt.Println("ShellSort Pass")
}

func BenchmarkShellSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		ShellSort(a, len(a))
	}
}

//堆排序
func TestHeapSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		HeapSort(a, len(a))
		if !check_sort(a) {
			t.Fatalf("HeapSort error\n")
		}
	}
	fmt.Println("HeapSort Pass")
}

func BenchmarkHeapSort(b *testing.B) {

	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		HeapSort(a, len(a))
	}
}

//快速排序
func TestQuickSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		QuickSort(a)
		if !check_sort(a) {
			t.Fatalf("QuickSort error\n")
		}
	}
	fmt.Println("QuickSort Pass")
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		QuickSort(a)
	}
}

//归并排序
func TestMergeSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randArray(arryLen)
		a = mergeSort(a)
		if !check_sort(a) {
			t.Fatalf("mergeSort error\n")
		}
	}
	fmt.Println("mergeSort Pass")
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := randArray(arryLen)
		a = mergeSort(a)
	}
}

func randBigArray(l int) []int {
	a := make([]int, l)
	for i := 0; i < l; i++ {
		a[i] = rand.Intn(900)
		a[i] += 100
		if a[i] > 900 {
			a[i] = 900
		}

	}

	return a
}

//桶排序
func TestBucketSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randBigArray(arryLen)
		BucketSort(a, 801)
		if !check_sort(a) {
			t.Fatalf("BucketSort error\n")
		}
	}
	fmt.Println("BucketSort Pass")
}

func BenchmarkBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := randBigArray(arryLen)
		BucketSort(a, 801)
	}
}

func TestBucketSort2(t *testing.T) {
	for i := 0; i < 100; i++ {

		a := randBigArray(arryLen)
		BucketSort2(a, 801)
		if !check_sort(a) {
			t.Fatalf("BucketSort2 error\n")
		}
	}
	fmt.Println("BucketSort2 Pass")
}

func BenchmarkBucketSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := randBigArray(arryLen)
		BucketSort2(a, 801)
	}
}

func TestArraySort(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randBigArray(arryLen)
		ArraySort(a, len(a))
		if !check_sort(a) {
			t.Fatalf("ArraySort error\n")
		}
	}
	fmt.Println("ArraySort Pass")
}

/*


BucketSort 20000000 data   3.419898266s
BucketSort2 20000000 data   1.528730197s
mergeSort  20000000 data   5.638596512s
ArraySort 20000000 data   26.447077ms

//数据量大但是数值范围不大的，可以考虑用内存换时间
*/

func TestBigData(t *testing.T) {
	//百万万数据排序
	slen := 1000000

	a := randBigArray(slen)
	start := time.Now()
	BucketSort(a, 801)
	fmt.Printf("BucketSort  %d data   %v\n", slen, time.Since(start))
	if !check_sort(a) {
		t.Fatalf("BucketSort   error\n")
	}

	a = randBigArray(slen)
	start = time.Now()
	BucketSort2(a, 801)
	fmt.Printf("BucketSort2 %d data   %v\n", slen, time.Since(start))
	if !check_sort(a) {
		t.Fatalf("BucketSort2 error\n")
	}

	a = randBigArray(slen)
	start = time.Now()
	a = mergeSort(a)
	fmt.Printf("mergeSort   %d data   %v\n", slen, time.Since(start))
	if !check_sort(a) {
		t.Fatalf("mergeSort error\n")
	}

	a = randBigArray(slen)
	start = time.Now()
	HeapSort(a, len(a))
	fmt.Printf("HeapSort   %d data   %v\n", slen, time.Since(start))
	if !check_sort(a) {
		t.Fatalf("HeapSort error\n")
	}

	a = randBigArray(slen)
	start = time.Now()
	ArraySort(a, len(a))
	fmt.Printf("ArraySort   %d data   %v\n", slen, time.Since(start))
	if !check_sort(a) {
		t.Fatalf("ArraySort error\n")
	}
}
