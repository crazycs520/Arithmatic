1. BinarySearch

2. BitCount

3. sort

   <img src="https://raw.githubusercontent.com/crazycs520/images/master/sort.png" style="zoom:100%" />

   * bubble 冒泡排序
   * Select   选择排序
   * insert   插入排序
   * Shell     希尔排序
   * Merge  归并排序
   * quick    快速排序
   * Heap   堆排序
   * bucket 桶排序 — 适用于能把输入数据 通过映射函数映射到不同的桶的数据
   * Arry 数组排序 — 只适用于输入数据量动态范围不大的情况
     * 如果输入量都是唯一不重复，可以考虑用[bit-map方法排序](http://www.cnblogs.com/Tour/p/4057416.html)或者[实现查找](https://wizardforcel.gitbooks.io/the-art-of-programming-by-july/content/a.1.html)，效率极高，还节省内存。
   * **TimSort **  —  推荐~ ,Python 和 JDK7 内置的排序算法。是归并排序和插入排序的混合算法。(留坑……还没搞懂timSort)
     * 参考资料：[什么是Timsort排序方法--知乎回答](https://www.zhihu.com/question/23928138)
     * 1，https://vimeo.com/146478455 概要的讲解timsort的实现以及timsort的bugs，因为是视频，所以相比论文我觉得更快看得懂，没字幕，听不懂怎么办，没事，演讲者有一个文章重新梳理视频内容[Proving that Android’s, Java’s and Python’s sorting algorithm is broken (and showing how to fix it)**](http://www.envisage-project.eu/proving-android-java-and-python-sorting-algorithm-is-broken-and-how-to-fix-it/)
     * 2，Tim peters自己写的论文 [https://svn.python.org/projects/python/trunk/Objects/listsort.txt](https://svn.python.org/projects/python/trunk/Objects/listsort.txt)

4. 二维“有序数组查找”  —binaryArrayFind

5.  最长递增子序列

6.  ​





# 动态规划 DP

1. 有面值为1元、3元和5元的硬币若干枚，如何用最少的硬币凑够11元？--leastCoin，参考：http://www.hawstein.com/posts/dp-novice-to-advanced.html


2. 背包问题:话说有一哥们去森林里玩发现了一堆宝石，他数了数，一共有n个。 但他身上能装宝石的就只有一个背包，背包的容量为C。这哥们把n个宝石排成一排并编上号： 0,1,2,…,n-1。第i个宝石对应的体积和价值分别为V[i]和W[i] 。排好后这哥们开始思考： 背包总共也就只能装下体积为C的东西，那我要装下哪些宝石才能让我获得最大的利益呢？--backpack.参考http://www.hawstein.com/posts/dp-knapsack.html
3. 最长公共子序列
4. 最长公共子串
5. 在一个数组arr中，找出一组不相邻的数字，使得最后的和最大。
6. 给定一个正整数s, 判断一个正整数数组arr中，是否有一组数字加起来等于s

 # 数据结构与算法--清华大学公开课学习(dsa_thu)
* 二叉树--binaryTree

# 剑指Offer(offer)
3. 二维数组查找
4. 替换空格
