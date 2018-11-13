### 使一个 goroutine 永远阻塞或睡眠（非译文）

一个空的 select 语句会永远地阻塞。

```
select {}  // 永远阻塞
```
