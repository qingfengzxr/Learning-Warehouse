* 位域

结构体中的冒号表示位域。位域出现的原因是由于某些信息的存储表示只需要几个bit位就可以表示而不需要一个完整的字节，同时也是为了节省存储空间和方便处理。

```c
其表示形式为：
struct 位域结构名
{
    类型说明符  位域名：位域长度
}
例如：
struct  bit_struct
{
	int  bit1:3;
    int  bit2:5;
    int  bit3:7;
}data;
其中bit_struct表示位域结构体，bit1、bit2、bit3表示对应的位域，data表示位域结构体定义的变量。整个位域结构体占用2个字节，bit1占3位，bit2占5位，bit1和bit2共用一个字节，bit3占7位，独占一个字节。
```

说明：

1. 位域必须存储在同一个类型中，不能跨类型，同时也说明位域的长度不会超过所定义类型的长度。如果一个定义类型单元里所剩空间无法存放下一个域，则下一个域应该从下一单元开始存放。例如：所定义的类型是int类型，一共32为，目前用掉了25位还剩下7位，这时要存储一个8位的位域元素，那么这个元素就只能从下一个int类型的单元开始而不会在前面一个int类型中占7为后面的int类型中占1位。

2. 如果位域的位域长度为0表示是个空域，同时下一个域应当从下一个字节单元开始存放。

3. 使用无名的位域来作为填充和调整位置，切记该位域是不能被使用的。

4. 位域的本质上就是一种结构体类型，不同的是其成员是按二进制位来分配的。

[详细示例]：<https://blog.csdn.net/yihongxiaoxiang/article/details/50327587>



* static 关键字

使用**static**关键字定义在头文件中的变量，其修饰的全局变量的作用域为定义的源文件。在多个源文件中引用该头文件，相当于每个变量都会在引用该头文件的源文件中定义一次。



#### 一些实用的写法

```c
#define CHECK_STATE_FLAG(flag, state)\
    do{\
        if(flag != state)\
        {\
            printf("func:%s,line:%d,flag statue fail\n",__FUNCTION__,__LINE__);\
            return FAILURE;\
        }\
    }while(0)
//仔细体会这个函数宏，真的是一种非常巧妙的用法，尽管功能简单，却能够在各处有需要的地方调用，
//并显示所在位置。就像面向对象语言中的错误处理，错误抛出一样实用。
#define UARTMSG_LOCK()   pthread_mutex_lock(&g_tUartMsgMutex);
#define UARTMSG_UNLOCK() pthread_mutex_unlock(&g_tUartMsgMutex);
//这只是一个简单的宏定义用法，但事实上，当代码里充斥着这种编写风格与思想时，代码的可阅读性确实
//有很大的提高。
// 短整型大小端互换
#define BigLittleSwap16(A)  ((((short int)(A) & 0xff00) >> 8) | \
                             (((short int)(A) & 0x00ff) << 8))
// 长整型大小端互换
#define BigLittleSwap32(A)  ((((long int)(A)  & 0xff000000) >> 24) | \
                             (((long int)(A)  & 0x00ff0000) >> 8) | \
                             (((long int)(A)  & 0x0000ff00) << 8) | \
                             (((long int)(A)  & 0x000000ff) << 24))


```

