### 切片和数组概述

> 原文：[https://yourbasic.org/golang/slices-explained/](https://yourbasic.org/golang/slices-explained/)

#### 基础

切片不保存任何数据，它只描述底层数组的一部分。

*   当改变切片的一个元素时，同时也修改了底层数组中对应的元素，共享同一个底层数组的其它切片也会被改变；
*   一个切片可以在底层数组的边界范围内进行缩放；
