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



* Docker容器

类似于一个轻量级的沙箱，Docker利用容器来运行和隔离应用。容器是从镜像创建的应用运行实例，可以将其启动、开始、停止、删除，而这些容器是**相互隔离、不可见的**。

* Docker仓库

类似于代码仓库，是Docker集中存放镜像文件的场所。