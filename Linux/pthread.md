* pthread_attr_setstacksize

**函数原型**

```c
#include <pthread.h>
int pthread_attr_setstacksize(pthread_attr_t *attr, size_t stacksize);
```

attr 是线程属性变量；stacksize 则是设置的堆栈大小。 返回值0，-1分别表示成功与失败。

**功能**: 重新设置堆栈大小

pthread_create 创建线程时，若不指定分配堆栈大小，系统会分配默认值，查看默认值方法如下：

```shell
# ulimit -s
8192
#
上述表示为8M；单位为KB。
也可以通过
# ulimit -a 
其中 stack size 项也表示堆栈大小。ulimit -s  value 用来重新设置stack 大小。


```

嵌入式中内存不是很大，若采用默认值的话，会导致出现问题，若内存不足，则 pthread_create 会返回 12，定义如下：

```c
#define EAGAIN 11
#define ENOMEM 12 /* Out of memory*/
```

