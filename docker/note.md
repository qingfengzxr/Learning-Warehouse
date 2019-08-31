#### 重要概念：

* **Docker镜像**

类似于虚拟机镜像，可以将它理解为一个面向Docker引擎的只读模板，包含了文件系统。镜像是创建Docker容器的基础。通过版本管理和增量的文件系统，Docker提供了一套十分简单的机制来创建和更新现有的镜像。

镜像文件一般由若干层组成，层其实是AUFS(Advanced File System,一种联合文件系统)中的重要概念，是实现增量保存与更新的基础。

```shell
$  sudo docker pull ubuntu:14.04 #下载14.04标签的镜像
$  sudo docker pull dl.dockerpool.com:5000/ubuntu #选择从DockerPool社区的镜像源下载
$  sudo docker run -t -i ubuntu /bin/bash #利用镜像创建一个容器并在其中运行bash应用
$  sudo docker images #列出本地主机上已有的所有镜像，镜像ID的信息十分重要，它唯一标识镜像
#可以使用docker tag命令为本地镜像添加新的标签，例如添加一个新的ubuntu:latest镜像标签
$  sudo docker tag dl.dockerpool.com:5000/ubuntu:latest ubuntu:latest
$  sudo docker inspect 5506de2b643b #可以获取镜像的详细信息，其返回一个JSON格式的消息
$  sudo docker search mysql #搜索远端仓库中共享的镜像，默认搜官方仓库，输出结果按星际评价排序
#可以删除镜像，其中IMAGE可以为标签或ID.(*当IMAGE是标签时，只有在镜像只剩一个标签时才会删除该镜像文件的所有AUFS层)
$  sudo docker rmi [IMAGE]
$  sudo docker rm [ID] #删除容器
```

**创建镜像**

```shell
$  sudo docker commit [OPTIONS] 
#OPTIONS = -a,--author=""作者信息
#OPTIONS = -m,--message=""提交信息
#OPTIONS = -p,--pause=true提交时暂停容器运行
#启动一个容器并运行一些操作后，该容器和原镜像相比，已经发生了改变，可以使用docker commit命令来提交为一个新的镜像。提交时可以使用ID或名称来指定容器。
$ sudo docker commit -m "Add a new file" -a "Docker Newbee" a925cb403f0 test
#运行顺利的话，命令返回新创建的镜像的ID信息
```

**基于本地模板导入**

```shell
#也可以直接从一个操作系统模板文件导入一个镜像。
$  sudo cat ubuntu-14.04-x86_64-minimal.tar.gz |docker import - ubuntu:14.04
```

**存出和载入镜像**

```shell
#存出镜像到本地文件
$  sudo docker save -o ubuntu_14.04.tar ubuntu:14.04
#从存出的本地文件再导入到本地镜像库
$  sudo docker load --input ubuntu_14.04.tar
$  sudo docker load < ubuntu_14.04.tar
```

**上传镜像**

```shell
#使用docker push上传镜像到仓库，默认上传到DockerHub官方仓库（需要登录）
$  sudo docker tag test:latest user/test:latest
$  sudo docker push user/test:latest
```



* **Docker容器**

类似于一个轻量级的沙箱，Docker利用容器来运行和隔离应用。容器是从镜像创建的应用运行实例，可以将其启动、开始、停止、删除，而这些容器是**相互隔离、不可见的**。

Docker带有**额外的可写文件层**。如果认为虚拟机是模拟运行的一整套操作系统（提供了运行态环境和其他系统环境）和跑在上面的应用。那么Docker容器就是独立运行的一个或一组应用，以及他们的必须环境。

**新建与启动容器**

```shell
#使用docker create可以新建一个容器.使用该命令创建的容器处于停止状态，可以使用docker start命令来启动它
$  sudo docker create -it ubuntu:latest
#使用docker run等价于先执行docker create 再执行 docker start
#下面的命令输出一个hello world ，之后容器将自动终止
$  sudo docker run ubuntu /bin/echo 'hello world'
#如果添加 --rm 标记，则容器在终止后会立刻删除
```

**使用docker run创建容器，后台运行的标准操作**包括：

1. 检查本地是否存在指定的镜像，不存在则从公有仓库下载
2. 利用镜像创建并启动一个容器
3. 分配一个文件系统，并在只读的镜像层外面挂载一层可读写层
4. 从宿主主机配置的网桥接口中桥接一个虚拟接口到容器中去
5. 从地址池配置一个IP地址给容器
6. 执行用户指定的应用程序
7. 执行完毕后容器被终止

```shell
#下面命令将启动一个bash终端，并允许用户交互
$  sudo docker run -t -i ubuntu:14.04 /bin/bash 
#使用-t选项可以让Docker分配一个伪终端并绑定到容器的标准输入上
#使用-i选型可以让容器的标准输入保持打开
```

_Docker容器认为，当运行的应用退出后，容器也没有了继续运行的必要_

**守护态运行**

```shell
$  sudo docker run -d ubuntu /bin/sh -c "while true;do echo hello world;sleep 1;done"
# -d选项能够让容器在后台以守护态形式运行
# --rm和-d参数不能同时使用
```

**终止容器**

```shell
#可以使用docker stop来终止一个运行中的容器
$  sudo docker stop ce5
```

**进入容器**

```shell
$  sudo docker attach nostalgic_hypatia
#使用attach命令，当多个窗口同时attach到同一个容器时，所有的窗口都会同步显示，一个窗口阻塞全都阻塞
$  sudo docker exec -ti 243c32535da7 /bin/bash
#exec能避免attach的问题
$  sudo nsenter --target 10981 --mount --uts --ipc --net --pid
#为了使用nsenter连接到容器，还需要找到容器的进程PID
```

**删除容器**

```shell
$  sudo docker rm [OPTIONS] CONTAINER [CONTAINER...]
# -f, --force=flase强行终止并删除一个运行中的容器
# -l, --link=flase删除容器的连接，但保留容器
# -v, --volumes=false 删除容器挂载的数据卷
```

**导入和导出容器**

```shell
$  sudo docker export CONTAINER
$  sudo docker export ce5 >test_for_run.tar  #导出容器ce5到test_for_run.tar文件
$  cat test_for_run.tar | sudo docker import - test/ubuntu:v1.0  
#导出的文件可以用docker import 命令导入，成为镜像
```

实际上，既可以使用docker load命令来导入镜像存储文件到本地的镜像库，又可以使用docker import命令来导入一个容器快照到本地镜像库。二者的区别在于容器快照文件将丢弃所有历史记录和元数据信息（即仅保存容器当时的快照状态），而镜像存储文件将保存完整记录，体积也比较大。



* **Docker仓库**

类似于代码仓库，是Docker集中存放镜像文件的场所。

**创建和使用私有仓库**

```shell
$  sudo docker run -d -p 5000:5000 registry
#这将自动下载并启动一个registry容器，创建本地的私有仓库服务
#默认情况下会将仓库创建在容器的/tmp/registry目录下，可以用过-v参数指定存放路径
#此时本地将启动一个私有仓库服务，监听端口为5000
```



* **数据卷**

数据卷是一个可供容器使用的特殊目录，它绕过文件系统，可以提供很多有用的特性：

1. 数据卷可以在容器之间共享和重用
2. 对数据卷的修改会立马生效
3. 对数据卷的更新不会影响镜像
4. 卷会一直存在，直到没有容器使用

数据卷的使用，类似于Linux下对目录或文件进行mount操作。



**创建数据卷**

```shell
#使用docker run命令加 -v 标记可以在容器内创建一个数据卷，多次使用-v标记可以创建多个数据卷
#下面使用training/webapp镜像创建一个web容器，并创建一个数据卷挂载到容器的/webapp目录：
$  sudo docker run -d -P --name web -v /webapp training/webapp python app.py
#*使用-v标记也可以指定挂载一个本地的已有目录到容器中去作为数据卷。本地目录的路径必须是绝对路径，如果路径不##存在，Docker会自动创建
#下面的命令加载主机的/src/webapp目录到容器的/opt/webapp目录
$  sudo docker run -d -P --name web -v /src/webapp:/opt/webapp training/webapp python app.py
#Docker挂载数据卷的默认权限是读写（rw），用户也可以通过，ro指定为只读
$  sudo docker run -d -P --name web -v /src/webapp:/opt/webapp:ro training/webapp python app.py 
#-v标记也可以从主机挂载单个文件到容器作为数据卷
$  sudo docker run --rm --it -v ~/.bash_history:/.bash_history ubuntu /bin/bash
#这样就可以记录在容器里输入过的命令历史了
```

**数据卷容器**

需要在容器之间共享一些持续更新的数据，最简答的方式是使用数据卷容器。数据卷容器其实就是一个普通的容器，只是专门用它来提供数据卷供其他容器挂载使用。

```shell
#首先：创建一个数据卷容器dbdata,并在其中创建一个数据卷挂载到/dbdata
$  sudo docker run -it -v /dbdata --name dbdate ubuntu
#然后：可以在其他容器中使用 --volumes-from来挂载dbdata容器中的数据卷
$  sudo docker run -it --volumes-from dbdata --name db1 ubuntu
$  sudo docker run -it --volumes-from dbdata --name db2 ubuntu
#上面的命令将创建db1和db2两个容器，并从dbdata容器挂载数据卷，此时两个容器都挂载同一个数据卷到相同的#/dbdata目录，三个容器任何一方在该目录下写入，其他容器都可以看到。
```

删除了挂载的容器，数据卷并不会被自动删除。如果要删除一个数据卷，必须在删除最后一个还挂载着它的容器时显式使用docker rm -v命令来指定同时删除关联的容器。

可以利用数据卷容器对其中的数据卷进行备份、恢复，以实现数据的迁移。

**备份**

```shell
$  sudo docker run --volumes-from dbdata -v $(pwd):/backup --name worker ubuntu tar cvf /backup/backup.tar /dbdata
#上面的命令首先利用ubuntu镜像创建了一个容器worker。使用--volumes-from dbdata参数让worker容器挂载#dbdata容器的数据卷(即dbdata数据卷)；使用-v$(pwd):/backup参数来挂载本地的当前目录到worker容器的#/backup/目录。
#work容器启动后，使用了tar cvf /backup/backup.tar /dbdata命令来将/dbdata下内容备份为容器内的#/backup/backup.tar,即宿主主机当前目录下的backup.tar
```

**恢复**

```shell
#首先要创建一个带有数据卷的容器dbdata2
$  sudo docker run -v /dbdata --name dbdata2 ubuntu /bin/bash
#然后创建另一个新的容器，挂载dbdata2的容器，并使用untar解压备份文件到所挂载的容器卷中即可
$  sudo docker run --volumes -form dbdata2 -v $(pwd):/backup busybox tar xvf /backup/backup.tar
```



* **网络配置**

Docker目前提供了映射容器端口到宿主主机和容器互联机制来为容器提供网络服务。除了使用端口映射机制来将容器内应用服务提供给外部网络，还可以通过容器互联系统让多个容器之间进行快捷的网络通信。

**端口映射实现访问容器**

在启动容器的时候，如果不指定对应参数，在容器外部是无法通过网络来访问容器内的网络应用和服务的。

```shell
#通过 -P 或者 -p 参数可以指定端口映射。
#当使用-P(大写的)标记时，Docker会随机映射一个49000~49900的端口至容器内部开放的网络端口
$  sudo docker run -d -P training/webapp python app.py
$  sudo docker ps -l
#上面的现象：本地主机的49155被映射到了容器的5000端口
#访问宿主主机的49115端口即可访问容器内Web应用提供的界面
# -p(小写的)则可以指定要映射的端口，注意，在一个指定端口上只可以绑定一个容器
#支持的格式：ip:hostPort:containerPort | ip::containerPort | hostPort:containerPort
#
#映射所有接口地址,此时默认会绑定本地所有接口上的所有地址
$  sudo docker run -d -p 5000:5000 training/webapp python app.py
#映射到指定地址的指定端口 ip:hostPort:containerPort
$  sudo docker run -d -p 127.0.0.1:5000:5000 traning/webapp python app.py
#映射到指定地址的任意端口 ip::containerPort
$  sudo docker run -d -p 127.0.0.1::5000 training/webapp python app.py
#使用udp标记来指定udp端口
$  sudo docker run -d -p 127.0.0.1:5000:5000/udp training/webapp python app.py
#查看映射端口配置
$  sudo docker port 
```

**容器的连接系统**

它会在源和接收容器之间创建一个隧道，接收容器可以看到源容器指定的信息。

连接系统依据容器的名称来执行，因此，首先需要自定义一个好记的容器名。自定义命名容器有两个好处：

1. 自定义的命名，比较好记
2. 当要连接其他容器时候，可以作为一个有用的参考点，比如连接web容器到db容器

```shell
#使用--name标记可以为容器自定义命名：
$  sudo docker run -d -P --name web training/webapp python app.py
$  sudo docker inspect -f "{{.name}}" 2adab9c82959 #可以查看容器的名字
```

```shell
#使用--link参数可以让容器之间安全的进行交互
$  sudo docker run -d -P --name web --link db:db training/webapp python app.py
#上述命令将使db容器和web容器建立互联我关系
#--link参数的格式为--link name:alias,其中name是要连接的容器的名称，alias是这个连接的别名。
```













































































