[TOC]

### 声明

以下所有相关命令行内容均基于ubuntu18.04。具体学习过程源自书籍《MySQL必知必会》。



### 关系数据库的基础知识

* 关系表

理解关系表的最好办法是看一个现实的例子。

```c
/*
假如有一个包含产品目录的数据库表，其中每种类别的物品占一行。对于每种物品要存储的信息包括产品描述和价格，以及生产该产品的供应商信息。
现在，假如有由同一供应商生产的多种物品，那么在何处存储供应商信息（如，供应商名、地址、联系方法等）呢？将这些数据与产品信息分开存储有如下几个重要的理由。

1. 因为同一个供应商生产的每个产品的供应商信息都是相同的，对每个产品重复此信息既浪费时间又浪费空间。
2. 如果供应商信息改变，只需改变一次即可。
3. 如果有重复数据（即每种产品都存储供应商信息），很难保证每次输入该数据的方式都相同。不一致的数据在报表中很	  难利用。

*/
```

1. **相同的数据出现多次决不是一件好事，此因素时关系数据库设计的基础。**

2. **关系表的设计就是要保证把信息分解成多个表，一类数据一个表。各表通过某些常用的值（即关系设计中的关系）互相关联。**

3. 关系数据库的可伸缩性远比非关系数据库要好。

**主键**：唯一标识表中每行的这个列（或这组列）称为主键。

**外键**：外键为某个表中的一列，它包含另一个表的主键值，定义了两个表之间的关系。







### MySQL的登录

```
mysql -h 127.0.0.1 -u root -p 
```

参数说明：

- **-h** : 指定客户端所要登录的 MySQL 主机名, 登录本机(localhost 或 127.0.0.1)该参数可以省略;
- **-u** : 登录的用户名;
- **-p** : 告诉服务器将会使用一个密码来登录, 如果所要登录的用户名密码为空, 可以忽略此选项。



### 建立一个新的数据库

```mysql
CREATE DATABASE learning;
```

这样就创建了一个名为 _learning_ 的数据库（数据源）

### 使用一个指定的数据库

```mysql
USE learning;
```

### 执行SQL脚本文件

```mysql
source XX.sql
```

注意：XX.sql 需要路径支持

### 显示内容

```mysql
SHOW DATABASES; //显示存在的数据库
SHOW TABLES;	//显示当前数据库存在的表
SHOW COLUMNS FROM TABLES;	//显示表列
```



### 检索数据

1. **数据的格式化是一个表示问题，而不是一个检索问题**



* 检索单个列

```mysql
SELECT 待检索列名 FROM 预检索表名;
```

* 检索多个列

```mysql
SELECT 待检索列名1,待检索列名2,待检索列名3,...,最后一个待检索列名 FROM 预检索表名;
```

* 检索所有列

```mysql
SELECT * FROM 预检索表名;		
```

* 检索时只返回不同的值

```mysql
SELECT DISTINCT 待检索列名 FROM 预检索表名;
```

注： DISTINCT 关键字应用于所有列而不是它的前置列。意味着当检索了多个列时，只有每个列的所有行都不同，才被检索出来。

* 限制检索结果数量

```mysql
SELECT 待检索列名 FROM 预检索表名 LIMIT 检索开始位置，要检索的行数;
```

注： 检索开始位置可省略，默认为0；如：

```Mysql
SELECT prod_name FROM products LIMIT 5;
```

返回不多于5行，从第  行0  开始（位置参数为0）。



### 排序检索数据

1. **关系数据库设计理论认为，如果不明确规定排序顺序，则不应该假定检索出的数据的顺序有意义**

* **子句**

SQL语句由子句组成，有些是必需的，有些是可选的。一个子句通常由一个关键字和所提供的数据组成。当前最明显的例子便是 `SELECT `语句的 ```FROM``` 子句。

* 字句的顺序

在给出 `ORDER BY` 字句时，应该保证它位于 `FROM` 字句之后。如果使用 `LIMIT` ,它必须位于 `ORDER BY` 之后。使用子句的次序不对将产生错误消息。

* 单列排序

```mysql
SELECT 待检索列名 FROM 预检索表名 ORDER BY 带检索列名;
```

注： 可以通过非选择检索列进行排序。

* 多列排序

```mysql
SELECT 待检索列名1,待检索列名2,待检索列名3 FROM 预检索表名 ORDER BY 待检索列名2,待检索列名3;
```

在上述排序中，将优先按照  _待检索列名2_  进行排序，仅在多个行具有相同的  _待检索列名2_  时才按 _待检索列名3_

进行再排序。

* 指定排序顺序

```mysql
SELECT 待检索列名1,待检索列名2,待检索列名3 FROM 预检索表名 ORDER BY 待检索列名1 DESC,待检索列名2;
```

数据排序默认以升序排序进行。当要进行 **降序排序** 时，必须指定 ```DESC``` 关键字。 **升序排序** 的关键字为 ```ASC```。

注意： 与 ```DISTINCT``` 关键字不同，```DESC``` 关键字只应用到直接位于其前面的列名。在上述排序命令中，降序只对待检索列名1进行指定，而待检索列名2仍然以默认的升序进行排序。



### 过滤数据

只检索所需数据需要指定 _搜索条件(search criteria)_ ，搜索条件也成为 _过滤条件_。

1. **MySQL在执行匹配时默认不区分大小写**

* 子句顺序

`WHERE`子句在表名`FROM`子句之后给出。

在同时使用了 `ORDER BY` 子句和 `WHERE` 子句时，应该让 `ORDER BY` 位于 `WHERE` 之后，否则会产生错误。

* WHERE子句条件操作符

[][]

| 操作符  | 说明             |
| ------- | ---------------- |
| =       | 等于             |
| <>      | 不等于           |
| !=      | 不等于           |
| <       | 小于             |
| <=      | 小于等于         |
| >       | 大于             |
| >=      | 大于等于         |
| BETWEEN | 在指定两个值之间 |

* 检查单个值

```mysql
SELECT 待检索列名1，待检索列名2 FROM 预检索表名 WHERE 待匹配列名 操作符 范围值;
```

注：待匹配列名 可以是 选择检索列也可以是非选择检索列。

比较字符串时需要用单引号来进行限定。数值比较则不需要引号。如：

```mysql
SELECT prod_name,prod_price FROM products WHERE prod_name = 'fuses';
SELECT vend_id,prod_name FROM products WHERE vend_id <> 1003;
```

* 检测范围值

```mysql
SELECT 待检索列名1，待检索列名2 FROM 预检索表名 WHERE 待匹配列名 BETWEEN 开始值 AND 结束值;
```

注： `BETWEEN` 匹配范围中所有的值，包括指定的开始值和结束值。

* 空值检查

在创建表时，可以指定其中的列是否可以不包含值。在一个列不包含值时，称其为包含空值 **NULL**。

**NULL** 无值，它与字段包含0、空字符串或仅仅包含空格不同。

```MYSQL
SELECT 待检索列名1 FROM 预检索表名 WHERE 待匹配列名 IS NULL;	 
```

NULL与不匹配：

在通过过滤选择出不具有特定值的行时，你可能希望返回具有NULL值的行。但是，不行。

因为未知具有特殊含义，数据库不知道他们是否匹配，所以在匹配过滤或者不匹配过滤时不返回它们。

因此，在过滤数据时，一定要验证返回数据中确实给出了被过滤列具有NULL的行。

* 组合条件检查

1. **MySQL支持使用NOT对IN、BETWEEN和EXISTS子句取反**

| 逻辑操作符 | 含义         |
| ---------- | ------------ |
| AND        | 与           |
| OR         | 或           |
| NOT        | 非           |
| IN         | 指定条件范围 |

计算次序：SQL像大多数语言一样，优先处理`AND`操作符。因此要使用圆括号明确地分组相应的操作符。示例：

```mysql
SELECT prod_name,prod_price FROM products 
WHERE (vend_id = 1002 OR vend_id = 1003) AND prod_price >= 10;

SELECT prod_name,prod_price FROM products WHERE vend_id IN (1002,1003) ORDER BY prod_name;

SELECT prod_name,prod_price FROM products WHERE vend_id NOT IN (1002,1003);
```

* 通配符过滤

**通配符**: 用来匹配值的一部分的特殊字符。

**搜索模式**:由字面值、通配符或者两者组合而成的搜索条件。

为在搜索子句中使用通配符，必须使用`LIKE`操作符。`LIKE`指示MySQL，后跟的搜索模式利用通配符匹配而不是直接使用相等匹配进行比较。

**使用通配符的技巧**:

1. 不要过度使用通配符。在能达到目的的情况下优先使用其他操作符。
2. 除非绝对必要，不要把通配符放在搜索模式的开始处。把通配符置于搜索模式的开始处，搜索起来最慢。

| 通配符 | 作用                            |
| ------ | ------------------------------- |
| %      | 任何字符出现任意次数            |
| _      | 和%作用一样，但是只匹配单个字符 |

注： 通配符不可以匹配 NULL。

```MYSQL
SELECT 待检索列名1，待检索列名2 FROM 预检索表名 WHERE 待匹配列名 LIKE '%sample';
```



### 正则表达式搜索

MySQL使用`WHERE`对正则表达式提供了初步的支持，允许你指定正则表达式，过滤`SELECT`检索出来的数据。

1. **MySQL仅支持多数正则表达式实现的一个很小的子集**。

2. **`LIKE`与`REGEXP`之间存在着一个终于差别，`LIKE`匹配整个列，而`REGEXP`在列值中进行匹配**。

3. **MySQL中的正则表达式匹配不区分大小写。为区分大小写需要使用`BINARY`关键字。如：

    `WHERE prod_name REGEXP BINARY 'JetPack .000'`

* 基本字符匹配

```mysql
SELECT prod_name FROM products WHERE prod_name REGEXP '.000' ORDER BY prod_name;
```

示例输出:

```shell
+--------------+
| prod_name    |
+--------------+
| JetPack 1000 |
| JetPack 2000 |
+--------------+
```

* 匹配几个字符之一

```mysql
SELECT prod_name FROM products WHERE prod_name REGEXP '[123] Ton' ORDER BY prod_name;
```

示例输出：

```shell
+-------------+
| prod_name   |
+-------------+
| 1 ton anvil |
| 2 ton anvil |
+-------------+
```

* 匹配特殊字符

MySQL的转义使用`\\`两个反斜杠。MySQL自己解析一个，正则表达式解析另一个。

| 空白元字符 | 说明     |
| ---------- | -------- |
| \\\f       | 换页     |
| \\\n       | 换行     |
| \\\r       | 回车     |
| \\\t       | 制表符   |
| \\\v       | 纵向制表 |

说明：每个元字符前面**只有两个反斜杠**，此处为了抵消掉markdown的语法，写了三个。

| 重复元字符 | 说明                          |
| ---------- | ----------------------------- |
| *          | 0个或多个匹配                 |
| +          | 1个或多个匹配（ 等于{1,} ）   |
| ?          | 0个或1个匹配 （ 等于{0,,1} ） |
| {n}        | 指定数目的匹配                |
| {n,}       | 不少于指定数目的匹配          |
| {n,m}      | 匹配数目的范围（m不超过255）  |

* 预定义字符集，字符类

| 字符类     | 说明                                                 |
| ---------- | ---------------------------------------------------- |
| [:alnum:]  | 任意字母和数字(同[a-ZA-Z0-9])                        |
| [:alpha:]  | 任意字符([a-zA-Z])                                   |
| [:blank:]  | 空格和制表(同[\\\t])                                 |
| [:cntrl:]  | ASCII控制字符(ASCII0到31和127)                       |
| [:digit:]  | 任意数字(同[0-9])                                    |
| [:graph:]  | 与[:print:]相同，但不包括空格                        |
| [:lower:]  | 任意小写字母(同[a-z])                                |
| [:print:]  | 任意可打印字符                                       |
| [:punct:]  | 既不在[:alnum:]又不在[:cntrl:]中的任意字符           |
| [:space:]  | 包括空格在内的任意空白字符(同[\\\f\\\n\\\r\\\t\\\v]) |
| [:upper:]  | 任意大写字母(同[A-Z])                                |
| [:xdigit:] | 任意十六进制数字(同[a-fA-F0-9])                      |

说明：每个双反斜杠**只有两个反斜杠**，此处为了抵消掉markdown的语法，写了三个。

* 定位元字符

| 定位元字符 | 说明       |
| ---------- | ---------- |
| ^          | 文本的开始 |
| $          | 文本的结尾 |
| [[:<:]]    | 词的开始   |
| [[:>:]]    | 词的结尾   |

* 简单的正则表达式测试

可以在不使用数据库表的情况下使用`SELECT`来测试正则表达式。`REGEXP`检查总是返回0(没有匹配)或者1(匹配)。可以用带文字串的REFEXP来测试表达式，并试验它们。如：

```mysql
SELECT 'hello' REGEXP '[0-9]';
```

示例输出：

```shell
+------------------------+
| 'hello' REGEXP '[0-9]' |
+------------------------+
|                      0 |
+------------------------+
```



### 创建计算字段

* 拼接

拼接是指将值联结到一起构成单个值。在MySQL的`SELECT`语句中，可使用`Concat()`函数来拼接两个列。

```mysql
SELECT Concat(vend_name,' (',vend_country,')') FROM vendors ORDER BY vend_name;
```

示例输出:

```shell
+-----------------------------------------+
| Concat(vend_name,' (',vend_country,')') |
+-----------------------------------------+
| ACME (USA)                              |
| Anvils R Us (USA)                       |
| Furball Inc. (USA)                      |
| Jet Set (England)                       |
| Jouets Et Ours (France)                 |
| LT Supplies (USA)                       |
+-----------------------------------------+
```

* 格式调整

可以使用`Trim()`函数去掉串左右两边的空格。`LTrim()`去掉串左边的空格。`RTtim()`去掉串右边的空格。

```mysql
SELECT Concat(RTrim(vend_name),' (',RTrim(vend_country),')') FROM vendors ORDER BY vend_name;
```

* 使用别名(alias)

别名：一个字段或值的替换名。别名使用`AS`关键字赋予。

别名的常见用途：在实际的表列名包含不符合规定的字符（如空格）时重新命名它，在原来的名字含混或容易误解 							   时扩充它等等。

```mysql
SELECT Concat(RTrim(vend_name),' (',RTrim(vend_country),')') AS vend_title FROM vendors
ORDER BY vend_name;	
```

示例输出：

```shell
+-------------------------+
| vend_title              |
+-------------------------+
| ACME (USA)              |
| Anvils R Us (USA)       |
| Furball Inc. (USA)      |
| Jet Set (England)       |
| Jouets Et Ours (France) |
| LT Supplies (USA)       |
+-------------------------+
```

注: 仔细观察拼接段落的示例，很容易就可以看出两者的区别。

* 执行算术计算

计算字段的另一常见用途是对检索出的数据进行算术计算。

测试计算：`SELECT`可以省略`FROM`子句以便简单地访问和处理表达式。如`SELECT 3*2;` `SELECT Now();`。

```mysql
SELECT prod_id,quantity,item_price FROM orderitems WHERE order_num = 20005;
SELECT prod_id,quantity,item_price,
	   quantity*item_price AS expanded_price
FROM orderitems WHERE order_num = 20005;
```

示例输出:

```shell
+---------+----------+------------+
| prod_id | quantity | item_price |
+---------+----------+------------+
| ANV01   |       10 |       5.99 |
| ANV02   |        3 |       9.99 |
| TNT2    |        5 |      10.00 |
| FB      |        1 |      10.00 |
+---------+----------+------------+
==============================================================
+---------+----------+------------+----------------+
| prod_id | quantity | item_price | expanded_price |
+---------+----------+------------+----------------+
| ANV01   |       10 |       5.99 |          59.90 |
| ANV02   |        3 |       9.99 |          29.97 |
| TNT2    |        5 |      10.00 |          50.00 |
| FB      |        1 |      10.00 |          10.00 |
+---------+----------+------------+----------------+
```



### 使用数据处理函数

SQL支持利用函数来处理数据。函数一般是在数据上执行的，它给数据的转换和处理提供了方便。

**函数没有SQL的可移植性强，如果决定使用函数，应该保证做好代码注释**

**大多数SQL实现支持以下类型的函数：**

1. 用于处理文本串的文本函数
2. 用于在数值数据上进行算术操作的数值函数
3. 用于处理日期和时间值并从这些值中提取特定成分的日期和时间函数
4. 返回DBMS正使用的特殊信息（如返回用户登录信息，检查版本细节）的系统函数。



* 文本处理函数

常用的文本处理函数如下：

| 函数        | 说明              |
| ----------- | ----------------- |
| Left()      | 返回串左边的字符  |
| Length()    | 返回串的长度      |
| Locate()    | 找出串的一个子串  |
| Lower()     | 将串转换为小写    |
| LTrim()     | 去掉串左边的空格  |
| Right()     | 返回串右边的字符  |
| RTrim()     | 去掉串右边的空格  |
| Soundex()   | 返回串的SOUNDEX值 |
| SubString() | 返回子串的字符    |
| Upper()     | 将串转换为大写    |



* 日期和时间处理函数

**用日期进行过滤需要注意一些别的问题和使用特殊的MySQL函数**

1. 无论什么时候指定一个日期，不管是插入或更新表值还是用WHERE字句进行过滤，日期格式必须为yyyy-mm=dd。这是首选的日期格式，因为它排除了多义性。

常用日期和时间处理函数如下：

| 函数          | 说明                             |
| ------------- | -------------------------------- |
| AddDate()     | 增加一个日期（天、周等）         |
| AddTime()     | 增加一个时间（时、分等）         |
| CurDate()     | 返回当前日期                     |
| CurTime()     | 返回当前时间                     |
| Date()        | 返回日期时间的日期部分           |
| DateDiff()    | 计算两个日期之差                 |
| Date_Add()    | 高度灵活的日期运算函数           |
| Date_Format() | 返回一个格式化的日期或时间串     |
| Day()         | 返回一个日期的天数部分           |
| DayOfWeek()   | 对于一个日期，返回对应的是星期几 |
| Hour()        | 返回一个时间的小时部分           |
| Minute()      | 返回一个时间的分钟部分           |
| Month()       | 返回一个时期的月份部分           |
| Now()         | 返回当前日期和时间               |
| Second()      | 返回一个时间的秒部分             |
| Time()        | 返回一个日期时间的时间部分       |
| Year()        | 返回一个日期的年份部分           |



* 数值处理函数

常用数值处理函数如下：

| 函数   | 说明               |
| ------ | ------------------ |
| Abs()  | 返回一个数的绝对值 |
| Cos()  | 返回一个角度的余弦 |
| Exp()  | 返回一个数的指数值 |
| Mod()  | 返回除操作的余数   |
| Pi()   | 返回圆周率         |
| Rand() | 返回一个随机数     |
| Sin()  | 返回一个角度的正弦 |
| Sqrt() | 返回一个数的平方根 |



### 汇总数据

* 聚集 函数

我们经常需要汇总数据而不用把它们实际检索出来，为此MySQL提供了专门的函数。这些函数就是**聚集函数**。

**聚集函数：**运行在行组上，计算和返回单个值的函数。

**标准偏差聚集函数**



SQL聚集函数表：

| 函数    | 说明             |
| ------- | ---------------- |
| AVG()   | 返回某列的平均值 |
| COUNT() | 返回某列的行数   |
| MAX()   | 返回某列的最大值 |
| MIN()   | 返回某列的最小值 |
| SUM()   | 返回某列值之和   |

补充说明：

AVG()：只能用来确定特定数值列的平均值，而且列名必须作为函数参数给出。为了获得多个列的平均值，必须使			 用多个AVG()函数。

​			 AVG()函数忽略列值为`NULL`的行。

COUNT(): 如果指定列名，则指定列值为`NULL`的行被COUNT()函数忽略，但如果COUNT()函数中用的是星号(*),则				 不忽略。

MAX()：忽略列值为`NULL`的行。

MIN()：忽略列值为`NULL`的行。	

SUM()：忽略列值为`NULL`的行。



* 聚集不同值

1. **不允许使用COUNT( DISTINCT )**

以上5个聚集函数都可以如下使用：

1. 对所有的行执行计算，指定`ALL`参数或不给参数(默认为`ALL`)。

2. 只包含不同的值，指定`DISTINCT`参数。

 ```mysql
SELECT AVG(prod_price) AS avg_price FROM products WHERE vend_id = 1003;
SELECT AVG(DISTINCT prod_price) AS avg_price FROM products WHERE vend_id = 1003;
 ```

示例输出：

```shell
+-----------+
| avg_price |
+-----------+
| 13.212857 |
+-----------+
================================================
+-----------+
| avg_price |
+-----------+
| 15.998000 |
+-----------+
```

* 组合聚集函数

`SELECT`语句可根据需要包含多个聚集函数。

```mysql
SELECT COUNT(*) AS num_items,
	   MIN(prod_price) AS price_min,
	   MAX(prod_price) AS price_max,
	   AVG(prod_price) AS price_avg
FROM products;
```

示例输出：

```shell
+-----------+-----------+-----------+-----------+
| num_items | price_min | price_max | price_avg |
+-----------+-----------+-----------+-----------+
|        14 |      2.50 |     55.00 | 16.133571 |
+-----------+-----------+-----------+-----------+
```

  

### 分组数据

* 数据分组

**分组**允许把数据分为多个逻辑组，以便能对每个组进行聚集计算。分组是在`SELECT`语句的`GROUP BY`子句中建立的。

**子句顺序**：`GROUP BY`子句必须出现在`WHERE`子句之后，`ORDER BY`子句之前。

```mysql
SELECT vend_id,COUNT(*) AS num_prods 
FROM products 
GROUP BY vend_id;
```

示例输出：

```shell
+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1001 |         3 |
|    1002 |         2 |
|    1003 |         7 |
|    1005 |         2 |
+---------+-----------+
```

说明：上述的`SELECT`语句指定了两个列，vend_id包含产品供应商的ID,num_prods为计算字段(用COUNT(*)函数建立)。`GROUP BY`子句指示MySQL按照vend_id排序并分组数据。这导致对每个vend_id而不是整个表计算num_prods一次。从输出中可以看到，供应商1001有3个产品，供应商1002有2个产品，供应商1003有7个产品，而供应商1005有2个产品。

`GROUP BY`子句后还可以跟`WITH ROLLUP`关键字，表示在分组统计的基础上再次进行汇总统计（在每个分组下都会有汇总统计）。

[更多内容参考][https://blog.csdn.net/qq_42254088/article/details/81904819]

```mysql
SELECT vend_id,COUNT(*) AS num_prods
FROM products
GROUP BY vend_id WITH ROLLUP;
```

示例输出：

```shell
+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1001 |         3 |
|    1002 |         2 |
|    1003 |         7 |
|    1005 |         2 |
|    NULL |        14 |
+---------+-----------+
```

**`GROUP BY`字句的重要规定**:

1. `GROUP BY`子句可以包含任意数目的列。这使得能对分组进行嵌套，为数据分组提供更细致的控制。
2. 如果在`GROUP BY`子句中嵌套了分组，数据将在最后规定的分组上进行汇总。换句话说，在建立分组时，指定的所有列都一起计算，不能再从个别的列中取回数据。
3. `GROUP BY`子句中列出的每个列都必须是**检索列**或**有效的表达式**（但不能是聚集函数）。如果`SELECT`中使用表达式，则必须在`GROUP BY`子句中指定相同的表达式。**不能使用别名**。

4. 除聚集计算语句外，`SELECT`语句中的每个列都必须在`GROUP BY`子句中给出。
5. 如果分组列中具有`NULL`值，则`NULL`将作为一个分组返回。如果列中有多行`NULL`值，它们将分为一组。
6. `GROUP BY`子句必须出现在`WHERE`子句之后，`ORDER BY`子句之前。



* 过滤分组

除了能用`GROUP BY`分组数据外，MySQL还允许过滤分组。

`WHERE`的过滤指定的是行而不是分组。`WHERE`没有分组的概念。为了实现分组过滤，需要使用`HAVING`子句，`HAVING`非常类似于`WHERE`，它能替代绝大部分的`WHERE`功能。两者唯一的差别是`WHERE`过滤行，而`HAVING`过滤分组。

**`HAVING`支持所有`WHERE`操作符**。

```mysql
SELECT cust_id,COUNT(*) AS orders 
FROM orders
GROUP BY cust_id
HAVING COUNT(*) >= 2;
```

示例输出：

```shell
+---------+--------+
| cust_id | orders |
+---------+--------+
|   10001 |      2 |
+---------+--------+
```

**同时使用`WHERE`和`HAVING`：**

```mysql
SELECT vend_id,COUNT(*) AS num_prods
FROM products
WHERE prod_price >= 10
GROUP BY vend_id
HAVING COUNT(*) >= 2;
```

示例输出：

```shell
+---------+-----------+
| vend_id | num_prods |
+---------+-----------+
|    1003 |         4 |
|    1005 |         2 |
+---------+-----------+
```

对于`WHERE`和`HAVING`的差别，也可以结合上述示例换另一种理解方法：

`WHERE`在数据分组前进行过滤，`HAVING`在数据分组后进行过滤。这是一种重要的区别，`WHERE`排除的行不包括在分组中，这可能会改变计算值，从而影响`HAVING`子句中基于这些值过滤掉的分组。

* 分组和排序

仅管我们经常发现用`GROUP BY`分组的数据确实以分组顺序输出，但情况并不总是这样，它不是SQL规范所要求的。因此一般在使用`GROUP BY`时也应该给出`ORDER BY`,以保证数据正确排序。



### SELECT字句顺序表

| 字句     | 说明               | 是否必须使用             |
| -------- | ------------------ | ------------------------ |
| SELECT   | 要返回的列或表达式 | 是                       |
| FROM     | 从中检索数据的表   | 仅在从表中选择数据时使用 |
| WHERE    | 行级过滤           | 否                       |
| GROUP BY | 分组说明           | 仅在按组计算聚集时使用   |
| HAVING   | 组级过滤           | 否                       |
| ORDER BY | 输出排序顺序       | 否                       |
| LIMIT    | 要检索的行数       | 否                       |



### 子查询

**查询(query)**:任何SQL语句都是查询。但此术语一般值`SELECT`语句。

SQL在版本4.1开始引入子查询的支持。**子查询(subquery)**即嵌套在其他查询中的查询。

1. 对于能嵌套的子查询数目没有限制，不过在实际使用中由于性能的限制，不能嵌套太多的子查询。

```mysql
SELECT cust_id
FROM orders 
WHERE order_num IN(SELECT order_num
                   FROM orderitems
                   WHERE prod_id = 'TNT2');
```

示例输出：

```shell
+---------+
| cust_id |
+---------+
|   10001 |
|   10004 |
+---------+
```

* 作为计算字段使用子查询

```mysql
SELECT cust_name,
	   cust_state,
	   (SELECT COUNT(*)
       	FROM orders
        WHERE orders.cust_id = customers.cust_id) AS orders
FROM customers
ORDER BY cust_name;
```

示例输出：

```shell
+----------------+------------+--------+
| cust_name      | cust_state | orders |
+----------------+------------+--------+
| Coyote Inc.    | MI         |      2 |
| E Fudd         | IL         |      1 |
| Mouse House    | OH         |      0 |
| Wascals        | IN         |      1 |
| Yosemite Place | AZ         |      1 |
+----------------+------------+--------+
```

说明：在这个示例中，使用了**完全限定列名**。这里还涉及到了一个新的概念，**相关子查询(correlated  subquery)**，涉及外部查询的子查询。任何时候只要列名可能有多义性，就必须使用这种语法（表名和列名由一个句点分隔）。

 

### 联结表

SQL最强大的功能之一就是能在数据检索查询的执行中**联结(join)**表。

**联结**：简单的说联结是一种机制，它不是物理实体，它用来在一条`SELECT`语句中关联表。联结在运行时关联表中		    正确的行。

1. **在一条`SELECT`语句中联结几个表时，相应的关系是在运行中构造的。**
2. **在联结两个表时，你实际上做的是将第一个表中的每一行与第二个表中的每一行配对。`WHERE`子句作为过滤条件，它只包含那些匹配给定条件（这里是联结条件）的行。**
3. 由没有联结条件的表关系返回的结果为笛卡尔积。检索出的行的数目将是第一个表中的行数乘以第二个表中的行数。笛卡尔积的联结类型又称为叉联结。



* 创建联结

```mysql
SELECT vend_name,prod_name,prod_price
FROM vendors,products
WHERE vendors.vend_id = products.vend_id
ORDER BY vend_name,prod_name;
```

示例输出：

```shell
+-------------+----------------+------------+
| vend_name   | prod_name      | prod_price |
+-------------+----------------+------------+
| ACME        | Bird seed      |      10.00 |
| ACME        | Carrots        |       2.50 |
| ACME        | Detonator      |      13.00 |
| ACME        | Safe           |      50.00 |
| ACME        | Sling          |       4.49 |
| ACME        | TNT (1 stick)  |       2.50 |
| ACME        | TNT (5 sticks) |      10.00 |
| Anvils R Us | .5 ton anvil   |       5.99 |
| Anvils R Us | 1 ton anvil    |       9.99 |
| Anvils R Us | 2 ton anvil    |      14.99 |
| Jet Set     | JetPack 1000   |      35.00 |
| Jet Set     | JetPack 2000   |      55.00 |
| LT Supplies | Fuses          |       3.42 |
| LT Supplies | Oil can        |       8.99 |
+-------------+----------------+------------+
```

说明：这里使用了`WHERE`子句来正确联结。

* 联结多个表

SQL对于一条`SELECT`语句中可以联结的表的数目没有限制。创建联结的基本规则也相同。

1. 出于性能的考虑，不应该联结太多的表，联结的表越多，性能下降的越厉害。

```mysql
SELECT prod_name,vend_name,prod_price,quantity
FROM orderitems,products,vendors
WHERE products.vend_id = vendors.vend_id
	  AND orderitems.prod_id = products.prod_id
	  AND order_num = 20005;
```

示例输出：

```shell
+----------------+-------------+------------+----------+
| prod_name      | vend_name   | prod_price | quantity |
+----------------+-------------+------------+----------+
| .5 ton anvil   | Anvils R Us |       5.99 |       10 |
| 1 ton anvil    | Anvils R Us |       9.99 |        3 |
| TNT (5 sticks) | ACME        |      10.00 |        5 |
| Bird seed      | ACME        |      10.00 |        1 |
+----------------+-------------+------------+----------+
```

* 表别名

表别名只在查询执行中使用。与列别名不一样，表别名不返回到客户机。

```mysql
SELECT cust_name,cust_contact
FROM customers AS c,orders AS o,orderitems AS oi
WHERE c.cust_id = o.cust_id
  AND oi.order_num = o.order_num
  AND prod_id = 'TNT2';
```

示例输出：

```shell
+----------------+--------------+
| cust_name      | cust_contact |
+----------------+--------------+
| Coyote Inc.    | Y Lee        |
| Yosemite Place | Y Sam        |
+----------------+--------------+
```



* 内部联结

内部联结又称为等值联结(equijion),它基于两个表之间的相等测试。

对于这种联结可以使用稍微不同的语法来明确指定联结的类型。

1. 尽管使用`WHERE`字句定义的联结的确比较简单，但是使用明确的联结语法能够确保不会忘记联结条件，有时候这样做也能影响性能。

```mysql
SELECT vend_name,prod_name,prod_price
FROM vendors INNER JOIN products
ON vendors.vend_id = products.vend_id;
```

示例输出：

```shell
+-------------+----------------+------------+
| vend_name   | prod_name      | prod_price |
+-------------+----------------+------------+
| Anvils R Us | .5 ton anvil   |       5.99 |
| Anvils R Us | 1 ton anvil    |       9.99 |
| Anvils R Us | 2 ton anvil    |      14.99 |
| LT Supplies | Fuses          |       3.42 |
| LT Supplies | Oil can        |       8.99 |
| ACME        | Detonator      |      13.00 |
| ACME        | Bird seed      |      10.00 |
| ACME        | Carrots        |       2.50 |
| ACME        | Safe           |      50.00 |
| ACME        | Sling          |       4.49 |
| ACME        | TNT (1 stick)  |       2.50 |
| ACME        | TNT (5 sticks) |      10.00 |
| Jet Set     | JetPack 1000   |      35.00 |
| Jet Set     | JetPack 2000   |      55.00 |
+-------------+----------------+------------+
```



* 自联结

表自己联结自己的联结类型称为自联结。

**自联结通常作为外部语句用来替代从相同表中检索数据时使用的子查询语句。虽然最终结果是一样的，但是有时候处理联结远比处理子查询快得多**。

```mysql
SELECT p1.prod_id, p1.prod_name
FROM products AS p1, products AS p2
WHERE p1.vend_id = p2.vend_id
  AND p2.prod_id = 'DTNTR';
```

示例输出：

```shell
+---------+----------------+
| prod_id | prod_name      |
+---------+----------------+
| DTNTR   | Detonator      |
| FB      | Bird seed      |
| FC      | Carrots        |
| SAFE    | Safe           |
| SLING   | Sling          |
| TNT1    | TNT (1 stick)  |
| TNT2    | TNT (5 sticks) |
+---------+----------------+
```



* 自然联结

标准联结（内部联结）返回所有数据，甚至相同的列多次出现。自然联结排除多次出现，使每个列只返回一次。

**自然联结是这样一种联结，其中你只能选择那些唯一的列。这一般是通过对表使用通配符`SELECT *`，对所有其他表的列使用明确的子集来完成的。**

```mysql
SELECT c.*, o.order_num,o.order_date,
	   oi.prod_id,oi.quantity,oi.item_price
FROM customers AS c, orders AS o, orderitems AS oi
WHERE c.cust_id = o.cust_id
  AND oi.order_num = o.order_num
  AND prod_id = 'FB';
```

示例输出：

```shell
+---------+-------------+----------------+-----------+------------+----------+--------------+--------------+-----------------+-----------+---------------------+---------+----------+------------+
| cust_id | cust_name   | cust_address   | cust_city | cust_state | cust_zip | cust_country | cust_contact | cust_email      | order_num | order_date          | prod_id | quantity | item_price |
+---------+-------------+----------------+-----------+------------+----------+--------------+--------------+-----------------+-----------+---------------------+---------+----------+------------+
|   10001 | Coyote Inc. | 200 Maple Lane | Detroit   | MI         | 44444    | USA          | Y Lee        | ylee@coyote.com |     20005 | 2005-09-01 00:00:00 | FB      |        1 |      10.00 |
|   10001 | Coyote Inc. | 200 Maple Lane | Detroit   | MI         | 44444    | USA          | Y Lee        | ylee@coyote.com |     20009 | 2005-10-08 00:00:00 | FB      |        1 |      10.00 |
+---------+-------------+----------------+-----------+------------+----------+--------------+--------------+-----------------+-----------+---------------------+---------+----------+--
```

说明：在这个例子中，通配符只对第一个表使用。所有其他列明确列出，所以没有重复的列被检索出来。



* 外部联结

联结包含了那些在相关表中没有关联的行。这种类型的联结称为**外部联结**。

在使用`OUTER JOIN`语法时，必须使用`RIGHT`或`LEFT`关键字指定包括其所有行的表(`RIGHT`指出的是`OUTER JOIN`右边的表，而`LEFT`指出的是`OUTER JOIN`左边的表)。这也引出了两种外部联结形式：`左外部联结`和`右外部联结`。

```mysql
SELECT customers.cust_id,orders.order_num
FROM customers LEFT OUTER JOIN orders
ON customers.cust_id = orders.cust_id;
```

示例输出：

```shell
+---------+-----------+
| cust_id | order_num |
+---------+-----------+
|   10001 |     20005 |
|   10001 |     20009 |
|   10002 |      NULL |
|   10003 |     20006 |
|   10004 |     20007 |
|   10005 |     20008 |
+---------+-----------+
```



* 使用带聚集函数的联结

```mysql
SELECT customers.cust_name,customers.cust_id,COUNT(orders.order_num) AS num_ord
FROM customers INNER JOIN orders
ON customers.cust_id = orders.cust_id
GROUP BY customers.cust_id;
```

示例输出：

```shell
+----------------+---------+---------+
| cust_name      | cust_id | num_ord |
+----------------+---------+---------+
| Coyote Inc.    |   10001 |       2 |
| Wascals        |   10003 |       1 |
| Yosemite Place |   10004 |       1 |
| E Fudd         |   10005 |       1 |
+----------------+---------+---------+
```



* 使用联结和联结条件

以下总结一下关于联结及其使用的某些要点。

1. 注意所使用的联结类型。一般我们使用内部联结，但使用外部联结也是有效的。
2. 保证使用正确的联结条件，否则将返回不正确的数据。
3. 应该总是提供联结条件，否则将得到笛卡尔积。
4. 在一个联结中可以包含多个表，甚至对于每个联结可以采用不同的联结类型。虽然这也做是合法的，一般也很有用，但应该在一起测试它们前，分别测试每个联结。这将使故障的排除更为简便。



### 组合查询

多数SQL查询都只包含从一个或多个表中返回数据的单条`SELECT`语句。MySQL也允许执行多个查询（多条`SELECT`语句），并将结果作为单个查询结果集返回。这些组合查询通常称为**并**或**复合查询**。

有两种基本情况，其中需要使用组合查询：

1. 在单个查询中从不同的表返回类似结构的数据；
2. 对单个表执行多个查询，按单个查询返回数据。



* 创建组合查询

利用`UNION`，可给出多条`SELECT`语句，将它们的结果组合成单个结果集。

`UNION`从结果集中自动去除了重复的行。如果想返回所有匹配的行，可使用`UNION ALL`而不是`UNION`。这里值得注意的是多个`WHERE`子句一定会自动去除重复的行，因此如果确实需要每个条件的匹配行全部出现（包括重复行），则必须使用`UNION ALL`而不是`WHERE`。



**`UNION`规则**：

1. UNION必须由两条或两条以上的`SELECT`语句组成，语句之间用关键字`UNION`分隔(因此，如果组合4条`SELECT`语句，将要使用`UNION`关键字)。
2. **`UNION`中的每个查询必须包含相同的列、表达式或聚集函数（不过每个列不需要以相同的次序给出）**。
3. 列数据类型必须兼容：类型不必完全相同，但必须时DBMS可以隐含地转换的类型。



```mysql
SELECT vend_id,prod_id,prod_price
FROM products 
WHERE prod_price <= 5
UNION
SELECT vend_id,prod_id,prod_price
FROM products
WHERE vend_id IN (1001,1002);
```

示例输出：

```shell
+---------+---------+------------+
| vend_id | prod_id | prod_price |
+---------+---------+------------+
|    1003 | FC      |       2.50 |
|    1002 | FU1     |       3.42 |
|    1003 | SLING   |       4.49 |
|    1003 | TNT1    |       2.50 |
|    1001 | ANV01   |       5.99 |
|    1001 | ANV02   |       9.99 |
|    1001 | ANV03   |      14.99 |
|    1002 | OL1     |       8.99 |
+---------+---------+------------+
```

* 对组合查询结果进行排序

在使用`UNION`组合查询时，只能使用一条`ORDER BY`子句，它必须出现在最后一条`SELECT`语句之后。对于结果集，不存在用一种方式排序一部分，而使用另一种方式排序另一部分的情况，因此不允许使用多条`ORDER BY`子句。

```mysql
SELECT vend_id,prod_id,prod_price
FROM products
WHERE prod_price <= 5
UNION 
SELECT vend_id,prod_id,prod_price
FROM products
WHERE vend_id IN (1001,1002)
ORDER BY vend_id,prod_price;
```

示例输出：

```shell
+---------+---------+------------+
| vend_id | prod_id | prod_price |
+---------+---------+------------+
|    1001 | ANV01   |       5.99 |
|    1001 | ANV02   |       9.99 |
|    1001 | ANV03   |      14.99 |
|    1002 | FU1     |       3.42 |
|    1002 | OL1     |       8.99 |
|    1003 | FC      |       2.50 |
|    1003 | TNT1    |       2.50 |
|    1003 | SLING   |       4.49 |
+---------+---------+------------+
```



### 全文本搜索

虽然**通配符**及**正则表达式**它们作为搜索机制非常有用，但仍存在几个重要的限制：

1. 性能----通配符和正则表达式匹配通常要求MySQL尝试匹配表中所有行（而且这些搜索极少使用表索引）。因此，由于被搜索行数不断增加，这些搜索可能非常耗时。
2. 明确控制----使用通配符和正则表达式匹配，很难（而且并不总能）明确地控制匹配什么和不匹配什么。例如，指定一个词必须匹配，一个词必须不匹配，而一个词仅在第一个词确实匹配的情况下才可以匹配或者才可以不匹配。
3. 智能化的结果----虽然基于通配符和正则表达式的搜索提供了非常灵活的搜索，但它们都不能提供一种智能化的选择结果的方法。

所有上述这些限制以及更多的限制都可以用**全文本搜索**来解决。在使用全文本搜索时，MySQL不需要分别查看每个行，不需要分别分析和处理每个词。MySQL创建指定列中各词的一个索引，搜索可以针对这些词进行。这样，MySQL可以快速有效地绝对哪些词匹配（哪些行包含它们），哪些词不匹配，它们匹配的频率，等等。

1. **为了进行全文本搜索，必须索引被搜索的列，而且要随着数据的更改不断重新索引**。在对表列进行适当设计后，MySQL会自动进行所有的索引和重新索引。但值得注意的是，在导入数据时应该先不启用`FULLTEXT`索引，应该首先导入所有数据，然后再修改表，定义`FULLTEXT`建立索引。这样有助于更快的导入数据及更快的更新索引（索引总数据的时间小于每行分时分别索引）。

2. 除非使用`BINARY`方式，否则全文本搜索不区分大小写。

3. **全文本搜索会返回以文本匹配的良好程度排序的数据**。



* 启用全文本搜索支持

一般在创建表时启用全文本搜索。`CREATE TABLE`语句接受`FULLTEXT`子句，它给出别索引列的一个逗号分隔的列表。

在定义后，MySQL自动维护该索引，在增加、更新或删除行时，索引随之自动更新。

**`FULLTEXT`子句可以指定多个列**。

**可以在创建表时指定`FULLTEXT`，也可以在稍后指定**。

```mysql
CREATE TABLE productnotes
(
    note_id    int        NOT NULL    AUTO_INCREMENT
    prod_id    char(10)   NOT NULL,
    note_date  datetime   NOT NULL,
    note_text  text       NULL,
    PRIMARY KEY(note_id),
    FULLTEXT(note_text)
)ENGINE=MyISAM;
```



* 进行全文本搜索

在索引之后，使用两个函数`Match()`和`Against()`执行全文本搜索，其中`Match()`指定被搜索的列，`Against()`指定要使用的搜索表达式。

**传递给`Match()`的值必须与`FULLTEXT()`定义中的相同。如果指定多个列，则必须列出它们（而且次序正确）。**



```mysql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('rabbit');
```

示例输出：

```shell
+----------------------------------------------------------------------------------------------------------------------+
| note_text                                                                                                            |
+----------------------------------------------------------------------------------------------------------------------+
| Customer complaint: rabbit has been able to detect trap, food apparently less effective now.                         |
| Quantity varies, sold by the sack load.
All guaranteed to be bright and orange, and suitable for use as rabbit bait. |
+----------------------------------------------------------------------------------------------------------------------+

```



```mysql
SELECT note_text,
	   Match(note_text) Against('rabbit') AS rank
FROM productnotes;
```

说明：下面的例子和其上的例子最明显的区别时，前者将返回所有行，因为没有`WHERE`字句进行过滤，同时将全文本匹配结果作为新行`rank`进行展示。



* 使用查询扩展

查询扩展使用`WITH QUERY EXPANSION`子句完成。

查询扩展用来设法放宽所返回的全文本搜索结果的范围。使用查询扩展时，MySQL对数据和索引进行两遍扫描来完成搜索：

1. 首先，进行一个基本的全文本搜索，找出与搜索条件匹配的所有行；
2. 其次，MySQL检查这些匹配行并选择所有有用的词(根据某种规则)；
3. 再其次，MySQL再次进行全文本搜索，这次不仅使用原来的条件，而且还使用所有有用的词。

```mysql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('anvils' WITH QUERY EXPANSION);
```



* 布尔文本搜索

MySQL支持全文本搜索的另外一种形式，称为**布尔方式（boolean mode）**。其可提供关于如下内容的细节：

1. 要匹配的词；
2. 要排斥的词（如果某行包含这个词，则不返回该行，即使它包含其他指定的词也是如此）；
3. 排列提示（指定某些词比其他词更加重要，更重要的词等级变高）；
4. 表达式分组；
5. 另外一些内容。

**布尔方式不同于上述使用的全文本搜索语法，它即使在没有定义`FULLTEXT`索引的情况下，也可以使用。但这是一种非常缓慢的操作**。

在布尔方式中，不按等级值降序排序返回的行。

```mysql
SELECT note_text FROM productnotes 
WHERE Match(note_text) Against('heavy' IN BOOLEAN MODE);

SELECT note_text FORM productnotes
WHERE Match(note_text) Against('heavy -rope*' IN BOOLEAN MODE);
```

全文本布尔操作符表：

| 布尔操作符 | 说明                                                         |
| ---------- | ------------------------------------------------------------ |
| +          | 包含，词必须存在                                             |
| -          | 排除，词必须不存在                                           |
| >          | 包含，而且增加等级值                                         |
| <          | 包含，而且减少等级值                                         |
| ()         | 把词组成子表达式(允许这些子表达式作为一个组被包含、排除、排列等) |
| ~          | 取消一个词的排序值                                           |
| *          | 词尾的通配符                                                 |
| " "        | 定义一个短语(与单个词的列表不一样，它匹配整个短语以便包含或排除这个短语) |

示例：

```mysql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('+rabbit +bait' IN BOOLEAN MODE);
```

说明：搜索匹配包含词rabbit和bait的行。

```mysql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('rabbit bait' IN BOOLEAN MODE);
```

说明：搜索匹配包含rabbit和bait中的至少一个词的行。

```mysql
SELECT note_text
FROM productnotes
WHERE Match(note_text) Against('"rabbit bait"' IN BOOLEAN MODE);
```

说明：搜索匹配短语rabbit bait而不是匹配两个词rabbit和bait。



* 全文本搜索的重要说明

1. 在索引全文本数据时，短词被忽略且从索引中排除。短词定义为哪些具有3个或3个以下字符的词（可以根据需要更改）。
2. MySQL带有一个内建的非用词(stopword)列表，这些词在索引全文本数据时将被忽略。如果需要，可以覆盖这个列表。
3. 许多词出现的频率很高，搜索它们没有用处(返回太多的结果)。因此，MySQL规定了一条50%规则，如果一个词出现在50%以上的行中，则将它作为一个非用词忽略。**50%规则不用于`IN BOOLEAN MODE`。
4. 如果表中的行数少于3行，则全文本搜索不返回结果(因此每个词或者不出现，或者出现在50%的行中)。
5. 忽略词中的的单引号。例如：don't 索引为 dont。
6. 不具有词分隔符(包括日语和汉语)的语言不能恰当地返回全文本搜索结果。
7. 仅在MyISAM数据库引擎中支持全文本搜索。



 ### 插入数据

使用`INSERT`语句来插入或添加行到数据库表，插入可以用几种方式使用：

1. 插入完整的行
2. 插入行的一部分
3. 插入多行
4. 插入某些查询结果

**重要注意事项**：

1. 如果某个列没有值，应该使用`NULL`值(假定表允许对该列指定空值)。
2. 对于自动增量，可以省略其值。

3. 如果表的定义允许，则可以在`INSERT`操作中省略某些列。被省略的列应该满足以下某个条件。
   - 该列的定义为允许`NULL`值（无值或空值）。
   - 在表定义中给出默认值。这表示如果不给出值，将使用默认值。



* 插入完整的行

```mysql
INSERT INTO customers
VALUES (NULL,
        'Ppe E.LaPew',
        '100 Main Street',
        'Los Angeles',
        'CA',
        '90046',
        'USA',
        NULL,
        NULL,
        NULL);
```

说明：这种语法没有指定列名，各个列必须以它们在表定义中出现的次序填充。这是一种不安全的语法，因为它极  			度依赖于表中列的定义次序。

```mysql
INSERT INTO customers(cust_name,
                      cust_address,
                      cust_city,
                      cust_state,
                      cust_zip,
                      cust_country,
                      cust_contact,
                      cust_email)
        VALUES('Ppe E.LaPew',
        	   '100 Main Street',
        	   'Los Angeles',
        	   'CA',
        	   '90046',
        	   'USA',
        	   NULL,
               NULL);
```

说明：这种语法与上述的语法显著的区别是明确的指出的列名。在插入行时，MySQL将用VALUES列表中对应的值填入列表中的对应项。因为提供了列名，VALUES必须按其指定的次序匹配指定列名，但不一定需要按各个列在实际表中的次序。这种方式的明显优点是，即使表的结构改变，此语句仍然可以正确工作。



* 插入多个行

可以使用多条`INSERT`语句。或者，只要每条`INSERT`语句中的列名和次序相同，也可以使用一条`INSERT`语句，组合提交多个行。

**MySQL用条`INSERT`语句处理多个插入比使用多条`INSERT`语句快**。

```mysql
INSERT INTO customers(cust_name,
                      cust_address,
                      cust_city,
                      cust_state,
                      cust_zip,
                      cust_country)
VALUES(
    	'Ppe E.LaPew',
        '100 Main Street',
        'Los Angeles',
    	'CA',
        '90046',
        'USA'
		),
	 (
        'M .Martian',
        '42 Galaxy Way',
        'New York',
        'NY',
        '11213',
        'USA'
     );
```



* 插入检索出的数据

可以利用`INSERT`将一条`SELECT`语句的结果插入表中。

```mysql
INSERT INTO customers(cust_id,
                     cust_contact,
                     cust_email,
                     cust_name,
                     cust_address,
                     cust_city,
                     cust_state,
                     cust_zip,
                     cust_country)
        SELECT cust_id,
        	   cust_contact,
               cust_email,
               cust_name,
               cust_address,
               cust_city,
               cust_state,
               cust_zip,
               cust_country
         FROM custnew;
```

说明：上述例子使用`INSERT SELECT`从custnew表中将所有数据导入customers。

* 优先级指定

如果数据检索是最重要的，可以通过在`INSERT`和`INTO`之间添加关键字`LOW_PRIORITY`，指示MySQL降低`INSERT`语句的优先级。如：`INSERT LOW_PRIORITY INTO`。



### 更新和删除数据

为了更新(修改)表中的数据，可以使用`UPDATE`语句。可采用两种方式使用`UPDATE`：

1. 更新表中的特定行

2. 更新表中所有行



* 更新数据

`UPDATE`语句非常容易使用，基本的`UPDATE`语句由3部分组成，分别是：

1. 要更新的表；
2. 列名和它们的新值；
3. 确定要更新行的过滤条件。

**为了删除某个列的值，可设置它为`NULL`（假如表定义允许`NULL`值）。

```mysql
UPDATE customers 
SET cust_email = 'elmer@fudd.com'
WHERE cust_id = 10005;

UPDATE customers
SET cust_name = 'The Fudds',
	cust_email = 'elmer@fudd.com'
WHERE cust_id = 10005;
```

说明：`UPDATE`语句总是以要更新的表的名字开始。`SET`命令用来将新值赋给被更新的列。



* 删除数据

为了从一个表中删除(去掉)数据，使用`DELETE`语句。有两种方式使用`DELETE`:

1. 从表中删除特定的行；
2. 从表中删除所有行。

**在使用`DELETE`语句时一定要小心，MySQL没有撤销操作。一定要注意使用`WHERE`子句**。

**`DELETE`不需要列名或者通配符。它删除整行而不是删除列。为了删除指定的列，应该使用`UPDATE`语句**。



* 删除所有行

如果想从表中删除所有行，不要使用`DELETE`。可使用 `TRUNCATE TABLE`语句，它完成相同的工作，但速度更快。因为它时通过删除原来的表并重新创建一个新表来实现的，而`DELETE`是逐行删除表中的数据。



* 更新和删除的指导原则

1. 除非确实打算更新和删除每一行，否则绝对不要使用不带`WHERE`子句的`UPDATE`或`DELETE`语句。
2. 保证每个表都有主键，尽可能像`WHERE`子句那样使用它。
3. 在对`UPDATE`或`DELETE`语句使用`WHERE`子句前，应该先用`SELECT`语句进行测试，以保证它的过滤时正确的。
4. 使用强制实施引用完整性的数据库，这样MySQL将不允许删除具有与其他表相关联的数据的行。



### 创建和操纵表



* 创建表

表的创建使用`CREATE TABLE`语句完成。为成功创建表，必须给出下列信息：

1. 新表的名字，在关键字`CREATE TABLE`之后给出；
2. 表列的名字和定义，用逗号分隔。

注意事项：

1. 在创建新表时，指定的表名必须不存在。为了防止意外覆盖已有的表，SQL要求首先手工删除该表，然后再重建。**如果仅想在一个表不存在时创建它，应该在表名后面给出`IF NOT EXISTS`**。这样做不检查已有表的模式是否与你打算创建的表模式相匹配，它只检查表名是否存在。

   

```mysql
CREATE TABLE customers 
(
	cust_id			int  		NOT NULL  AUTO_INCREMENT,
    cust_name		char(50)	NOT NULL,
    cust_address	char(50)	NULL,
    cust_city		char(50)	NULL,
    cust_state		char(5)		NULL,
    cust_zip		char(10)	NULL,
    cust_country	char(50)	NULL,
    cust_contact	char(50)	NULL,
    cust_email		char(255)	NULL,
    PRIMARY KEY (cust_id)
)ENGINE=InnoDB;
```

**说明：**

1. 实际的表定义(所有列)括在圆括号中，个各列之间用逗号分隔。

2. 这个表由9列组成，每列的定义以列名(它在表中必须是唯一的)开始，后跟列的数据类型。

3. 表的主键可以在创建表时用关键字`PRIMARY KEY`指定。

4. `NULL`值就是没有值或者缺值。允许`NULL`值的列也允许在插入行时不给出该列的值，反之不允许`NULL`值的行在插入或更新时，该列必须有值。

5. `AUTO_INCREMENT`告诉MySQL，本列为**自动增量**，每增加一行时自动增加。每个表只允许一个`AUTO_INCREMENT`列，而且它必须被索引(如，通过使它成为主键)。

   自动增量也可以通过在`INSERT`语句中使用其他值覆盖，这样后续的增量将使用该手工插入的值。

   **获取**最后一个增量值可以使用`last_insert_id()`来获取，如：`SELECT last_insert_id();`。



* 多个列组成主键

为创建多个列组成的主键，应该以逗号分隔的列表给出各列名。

**主键可以在创建表时定义，也可以在创建表后定义**。

**主键只能使用不允许`NULL`值的列**。

```mysql
CREATE TABLE orderitems
(
order_num	int 		NOT NULL,
order_item	int			NOT NULL,
prod_id		char(10)	NOT NULL,
quantity	int			NOT NULL,
item_price  decimal(8,2)NOT NULL,
PRIMARY KEY (order_num,order_item)
)ENGINE=InnoDB;
```



* 指定默认值

如果插入行时没有给出值，MySQL允许指定此时使用的默认值。在创建表时用`DEFAULT`关键字指定。

```mysql
CREATE TABLE orderitems
(
order_num	int 		NOT NULL,
order_item	int			NOT NULL,
prod_id		char(10)	NOT NULL,
quantity	int			NOT NULL	DEFAULT 1,
item_price  decimal(8,2)NOT NULL,
PRIMARY KEY (order_num,order_item)
)ENGINE=InnoDB;
```



* 引擎类型

MySQL与其他DBMS不一样，它具有多种引擎。可以使用`ENGINE=`语句进行指定。

1. **引擎类型可以混用**。每个表都可以有自己的引擎类型。

2. **外键不能跨引擎**。外键用于强制实施引用完整性，因此不能跨引擎。即使用一个引擎的表不能引用具有使用不同引擎的表的外键。
3. 

```C
/*
几个重要的引擎：
1. InnoDB 是一个可靠的事务处理引擎，它不支持全文本搜索。
2. MEMORY 的功能等同于MyISAM，但它的数据存储在内存(不是磁盘)中，因此速度很快，特别适合临时表。
3. MyISAM 时一个性能极高的引擎，它支持全文本搜索，但不支持事务处理。
*/
```



* 更新表定义

为更新表定义，可使用`ALTER TABLE`语句。但在理想状态下，当表中存储数据以后，该表就不应该再被更新，因此表的设计过程需要花费大量的时间来考虑，十分重要。

**在进行改动前，最好做一个完整的备份。因为数据库表的更改不可撤销**。

为使用`ALTER TABLE`更新表结构，必须给出下面信息：

1. 在`ALTER TABLE`之后给出要更改的表名(该表必须存在，否则将出错)；
2. 所做更改的列表。

```mysql
#给表添加一个列
ALTER TABLE 表名
ADD 列定义;
```

```mysql
ALTER TABLE vendors
ADD vend_phone CHAR(20);
```

说明：必须明确给出列的数据类型。

```mysql
#删除刚刚添加的列
ALTER TABLE vendors
DROP COLUMN vend_phone;
```

**`ALTER TABLE`的一个常见用途是定义外键**。

```mysql
ALTER TABLE orderitems
ADD CONSTRAINT fk_orderitems_orders
FOREIGN KEY (order_Num) REFERENCES orders(order_num);
```



* 复杂的表结构更改

复杂的表结构更改一般需要手动删除过程，它涉及以下步骤：

1. 用新的列布局创建一个新表；
2. 使用`INSERT SELECT`语句，从旧表复制数据到新表。如果有必要，可使用转换函数和计算字段；
3. 检验包含所需数据的新表；
4. 重命名旧表(如果确定可以删除它)；
5. 用旧表原来的名字重命名新表；
6. 根据需要，重新创建触发器、存储过程、索引和外键。



* 删除表

删除表没有确认，也不能撤销。

```mysql
DROP TABLE customers2;
```



* 重命名表

```mysql
RENAME TABLE customers TO customers;

RENAME TABLE backup_customers TO customers,
			 backup_vendors TO vendors,
			 backup_products TO products;
```



### 视图

视图是虚拟的表。与包含数据的表不一样，视图只包含使用时动态检索数据的查询。

**作为视图，它不包含表中应该有的数据，它包含的是一个SQL查询。**

**视图仅仅是用来查看存储在别处的数据的一种设施。视图本身不包含数据，因此它们返回的数据是从其他表中检索出来的。在添加或更改这些表中的数据时，视图将返回改变后的数据。**

**视图提供了一种MySQL的`SELECT`语句层次的封装，可用来简化数据处理及重新格式化基础数据或保护基础数据。**

* 为什么使用视图

1. 重用SQL语句。
2. 简化复杂的SQL操作。在编写查询后，可以方便地重用它而不必知道它的基本查询细节。
3. 使用表的组成部分而不是整个表。
4. 保护数据。可以给用户授予表的特定部分的访问权限而不是整个表的访问权限。
5. 更改数据格式和表示。视图可返回与底层表的表示和格式不同的数据。



* 视图的规则和限制

1. 与表一样，视图必须唯一命名。
2. 对于可以创建的视图数目没有限制。
3. 为了创建视图，必须具有足够的访问权限。这些限制通常由数据库管理人员授予。
4. 视图可以嵌套，即可以利用从其他视图中检索数据的查询来构造一个视图。
5. `ORDER BY`可以用在视图中，但如果从该视图检索数据的`SELECT`语句中也含有`ORDER BY`，那么该视图中的`ORDER BY`将被覆盖。
6. 视图不能索引，也不能有关联的触发器或默认值。
7. 视图可以和表一起使用。例如，编写一条联结表和视图的`SELECT`语句。



* 使用视图

1. 视图使用`CREATE VIEW`创建。
2. 使用`SHOW CREATE VIEW viewname;`来查看创建视图的语句。
3. 用`DROP`删除视图，其语法为`DROP VIEW viewname;`。
4. 更新视图时，可以先用`DROP`再用`CREATE`,也可以直接用`CREATE OR REPLACE VIEW`，如果更新的视图不存在，则创建一个视图，如果存在，则替换视图。

**如果从视图中检索数据时使用了一条`WHERE`子句，则两组子句(一组在视图中，另一组是传递给视图的)将自动结合**。

**视图主要用于检索数据，而不用于更新数据**。

```mysql
CREATE VIEW productcustomers AS
SELECT cust_name,cust_contact,prod_id
FROM customers,orders,orderitems
WHERE customers.cust_id = orders.cust_id
  AND orderitems.order_num = orders.order_num;
```

说明：这条语句创建了一个名为`productcustomers`的视图，它联结三个表，以返回已订购了任意产品的所有客户的列表。为检索订购了产品TNT2的客户，可使用如下语句：

```mysql
SELECT cust_name,cust_contact
FROM productcustomers
WHERE prod_id = 'TNT2';
```

示例输出：

```shell
+----------------+--------------+
| cust_name      | cust_contact |
+----------------+--------------+
| Coyote Inc.    | Y Lee        |
| Yosemite Place | Y Sam        |
+----------------+--------------+
```

 

* 更新视图

可以使用`INSERT`、`UPDATE`、`DELETE`对视图进行更新。不过更新一个视图将更新其基表，如果对视图增加或删除行，实际上是对其基表增加或删除行。

但并不是所有视图都可以更新，如果视图定义中有以下操作，则不能进行更新：

1. 分组(使用`GROUP BY`和`HAVING`)；
2. 联结；
3. 子查询；
4. 并；
5. 聚集函数；
6. DISTINCT；
7. 导出 (计算) 列。



### 存储过程

存储过程简单来说，就是为以后的使用而保存的一条或多条MySQL语句的集合。



* 为什么使用存储过程

1. 通过把处理封装在容易使用的单元中，简化复杂的操作(正如前面例子所述)。
2. 由于不要求反复建立一系列处理步骤，这保证了处理的统一性，也就保证了数据的完整性。
3. 简化对变动的管理。如果表名、列名或业务逻辑有变化，只需要更改存储过程的代码即可。

(2 和 3 在一定的程度上来说都可以延伸为安全性。通过存储过程限制对基础数据的访问减少数据讹误的机会)

4. 提高性能。使用存储过程比使用单独的SQL语句要快。
5. 存在一些只能用在单个请求中的MySQL元素和特性，存储过程可以使用它们来编写功能更强更灵活的代码。



* 执行存储过程

MySQL称存储过程的执行为调用，MySQL执行存储过程的语句为`CALL`。`CALL`接受存储过程的名字以及需要传递给它的任意参数。

```mysql
CALL 存储过程名(@param_1,
              @param_2,
              @param_3);
```



* 创建存储过程

```mysql
CREATE PROCEDURE 存储过程名()
BEGIN
	待执行的SQL语句
END;
```

示例:

```mysql
CREATE PROCEDURE productpricing()
BEGIN
	SELECT Avg(prod_price) AS priceaverage
	FROM products;
END;
```



* 删除存储过程

```mysql
DROP PROCEDURE productpricing;
```





* 注意事项--mysql命令行客户机的分隔符

默认的MySQL语句分隔符为`；`。因为存储过程中的语句中也以`;`作为结束，因此在命令行客户机中使用`END；`时，应临时更改命令行实用程序的语句分隔符，稍作修改，如示：

```mysql
DELIMITER	//
CREATE PROCEDURE productpricing()
BEGIN
	SELECT Avg(prod_price) AS priceaverage
	FROM products;
END	//

DELIMITER ;
```

说明：`DELIMITER` 可以告诉命令行使用程序使用其所带符号作为新的语句结束分隔符。



* 使用参数

**变量**：内存中一个特定的位置，用来临时存储数据。

**变量名**：所有MySQL变量都必须以@开始。

```mysql
CREATE PROCEDURE productpricing(
	OUT pl DECIMAL(8,2),
    OUT ph DECIMAL(8,2),
    OUT pa DECIMAL(8,2)
)

BEGIN 
	SELECT Min(prod_price)
	INTO pl
	FROM products;
	SELECT Max(prod_price)
	INTO ph
	FROM products;
	SELECT Avg(prod_price)
	INTO pa
	FROM products;
END;
```

说明：

此存储过程接受三个参数。在创建带参数的存储过程时，每个参数必须具有指定的类型。关键字`OUT`指出相应的参数用来从存储过程传出一个值(返回给调用者)。MySQL支持`IN`（传递给存储过程）、`OUT` (从存储过程传出一个值，返回给调用者)和`INOUT`(对存储过程传入和传出)类型的参数。存储过程位于`BEGIN`和`END`语句内。

```mysql
CALL productpricing(@pricelow,
                    @pricehigh,
                   	@priceaverage);
```

说明：在调用时，这条语句并不显示任何数据。它返回以后可以显示(或在其他处理中使用)的变量。



* 建立智能存储过程

存储过程可以包含业务规则和处理逻辑。这将使得存储过程更加的智能。



* 检查存储过程

可以使用`SHOW CREATE PROCEDURE`来显示一个创建存储过程的语句。

```mysql
SHOW CREATE PROCEDURE productpricing;
```

如果要获得包括何时、由谁创建等详细信息的存储过程列表，可以使用`SHOW PROCEDURE STATUS`。

```mysql
SHOW PROCEDURE STATUS LIKE 'productpricing';
```



### 游标

游标(cursor)是一个存储在MySQL服务器上的数据库查询，它不是一条`SELECT`语句，而是被该语句检索出来的结果集。在存储了游标之后，应用程序可以根据需要滚动或浏览其中的数据。游标主要用户交互式应用。



* 使用游标

使用游标涉及几个明确的步骤。

1. 在能够使用游标前，必须声明(定义)它。这个过程实际上没有检索数据，它只是定义要使用的`SELECT`语句。
2. 一旦声明后，必须打开游标以供使用。这个过程用前面定义的`SELECT`语句把数据实际检索出来。
3. 对于填有数据的游标，根据需要取出(检索各行)。
4. 在结束游标使用时，必须关闭游标。



* 创建游标

游标用`DECLARE`语句创建。`DECLARE`命名游标，并定义相应的`SELECT`语句，根据需要带`WHERE`和其他子句。例如，下面的语句定义了名为`ordernumbers`的游标，使用了可以检索所有订单的`SELECT`语句。

**DECLARE语句的次序**：`DECLARE`语句的发布存在特定的次序。用`DECLARE`语句定义的局部变量必须在定义任意游标或句柄之前定义，而句柄必须在游标之后定义。

```mysql
CREATE PROCEDURE processorders()
BEGIN 	
	DECLARE ordernumbers CURSOR
	FOR
	SELECT order_num FROM orders;
END;
```

说明：这个游标在存储过程处理完成后消失，它局限于存储过程。



* 打开和关闭游标

```mysql
OPEN 游标名;#打开游标
CLOSE 游标名;#关闭游标
```

**如果不明确关闭游标，MySQL将会在到达`END`语句时自动关闭它**。



* 使用游标数据

在一个游标被打开后，可以使用`FETCH`语句分别访问它的每一行。`FETCH`指定检索什么数据(所需的列)，检索出来的数据存储在什么地方。它还向前移动游标中的内部行指针，使下一条`FETCH`语句检索下一行(不重复读取同一行)。

```mysql
CREATE PROCEDURE local processorders()
BEGIN
	--Declare local variables
	DECLARE done BOOLEAN DEFAULT 0;
	DECLARE o INT;
	
	--Declare the cursor
	DECLARE ordernumbers CURSOR
	FOR
	SELECT order_num FROM orders;
	--Declare continue handler
	DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done=1;
	
	--Open the cursor
	OPEN ordernumbers;
	
	--Loop through all rows
	REPEAT
		--Get order number
		FETCH ordernumbers INTO o;
		
	--End of Loop
	UNTIL done END REPEAT;
	
	--Close the cursor
	CLOSE ordernumbers;
END;
```

说明：这里需要着重说明下语句：

```mysql
DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done=1;
```

这条语句定义了一个`CONTINUE HANDLER`,它是在条件出现时被执行的代码。它指出当`SQLSTATE '02000'`出现时，`SET done=1`。`SQLSTATE '02000'`是一个未找到条件，当`REPEAT`由于没有更多的行供循环而不能继续时，出现这个条件。



### 触发器

触发器是MySQL响应以下任意语句而自动执行的一条MySQL语句(或位于`BEGIN`和`END`语句之间的一组语句):

1. `DELETE`
2. `INSERT`
3. `UPDATE`

**只有表才支持触发器，视图不支持(临时表也不支持)**。

**触发器按每个表每个事件每次地定义，每个表每个事件每次只允许一个触发器，因此每个表最多支持6个触发器**。

**MySQL触发器不支持`CALL`语句，这表示不能从触发器内调用存储过程。所需的存储过程代码需要复制到触发器中**。



* 创建触发器

在创建触发器时，需要给出4条信息：

1. 唯一的触发器名;
2. 触发器关联的表；
3. 触发器应该响应的活动(`DELETE`、`INSERT`或`UPDATE`)；
4. 触发器何时执行(处理之前或之后)。`AFTER` or`BEFORE`。

```mysql
CREATE TRIGGER 触发器名 AFTER INSERT ON 关联表名
FOR EACH ROW SELECT '显示文本';
# 这里的显示文本非必须，只是一个示例。类似语句将在每次成功插入数据后，显示"显示文本"。
```



* 删除触发器

```mysql
DROP TRIGGER 触发器名;
```



* INSERT触发器

INSERT触发器在`INSERT`语句之前或之后执行。需要知道以下几点:

1. 在INSERT触发器代码内，可引用一个名为NEW的虚拟表，访问被插入的行；
2. 在BEFORE INSERT触发器中，NEW中的值也可以被更新(允许更改被插入的值)；
3. 对于`AUTO_INCREMENT`列，NEW在`INSERT`执行之前包含0，在`INSERT`执行之后包含新的自动生成值。



* DELETE触发器

DELETE触发器在`DELETE`语句执行之前或之后执行。需要知道以下两点:

1. 在`DELETE`触发器代码内，你可以引用一个名为OLD的虚拟表，访问被删除的行；
2. `OLD`中的值全都是只读的，不能更新。

```mysql
CREATE TRIGGER deleteorder BEFORE DELETE ON orders
FOR EACH ROW
BEGIN 
	INSERT INTO archive_orders(order_num,order_date,cust_id)
	VALUES(OLD.order_num,OLD.order_date,OLD.cust_id)
END;
```

说明：上述例子演示使用OLD虚拟表保存将要被删除的行到一个存档表中。同时，上述例子还是用了多语句触发器的技巧，使得触发器能够容纳多条语句。



* UPDATE触发器

UPDATE触发器在`UPDATE`语句执行之前或之后执行。需要知道以下几点：

1. 在UPDATE触发器代码中，你可以引用一个名为OLD的虚拟表访问以前(UPDATE语句之前)的值，引用一个名为NEW的虚拟表访问新的更新的值；
2. 在BEFORE UPDATE触发器中，NEW中的值可能也被更新(允许更改将用于`UPDATE`语句中的值)；
3. OLD中的值全都是只读的，不能更新。





















