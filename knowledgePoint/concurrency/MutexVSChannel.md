### Mutext 和 Channel 对比

可以使用 mutex 和 channel 解决竞争状态问题。如何进行选择？ 这取决于要解决的问题。如果要解决的问题更适合 mutex 那就
使用 mutex。如果需要，不要对使用 mutext 犹豫不决。如果要解决的问题更适合 channel，那就使用它。

所有并发问题都使用 channel 解决是错误的，毕竟语言给我们提供了 mutex 或 channel 的选择。

当 goroutine 需要通信时，通常使用 channel；而当只有一个 goroutine 应该访问临界资源时就使用 mutex。

应该为问题选择工具，而不是尝试用工具适应问题。