### 数据序列化

客户端和服务端之间的通信需要数据交换。这些数据可能是高度结构化的，但是为了传输必须要序列化。

客户端和服务端需要通过消息来交换信息。TCP 和 UDP 提供了消息的传输机制。客户端和服务端间的两个进程之间还需要一种协议来使得消息交换变得有意义。

消息以字节序列的方式在网络中传输，本身是没有结构的，仅仅是线性的字节流。

程序往往会创建复杂的数据结构来保存当前程序的状态。在跟远程的客户端或服务端通信时，程序会尝试通过网络传输这些数据结构，但是 IP、TCP、UDP 等网络包并不知道这些数据结构的意义。因此程序必须把所有的数据结构进行序列化为字节流以便写入，另外对字节流进行反序列化成合适的数据结构以便读取。

以下是一个序列化的小示例：

| | |
| :--: | :--: |
| fred | programmer |
| liping | analyst |
| surreerat | manager |

要传输上表的数据，有序列化的方式。例如，可以按以下方式：

```
3               // 3    rows, 2 columns assumed
4   fred        // 4    char string, col1
10  programmer  // 10   char string, col2
6   liping      // 6    char string, col1
7   analyst     // 7    char string, col2
8   sureerat    // 8    char string, col1
7   manager     // 7    char string, col2
```
