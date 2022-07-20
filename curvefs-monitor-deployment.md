部署 CurveFS 监控
===

- [部署 CurveFS 监控](#部署-curvefs-监控)
  - [第 1 步：环境准备](#第-1-步环境准备)
  - [第 2 步：准备配置文件](#第-2-步准备配置文件)
  - [第 3 步：部署监控系统](#第-3-步部署监控系统)
  - [第 4 步：导入 graphana 面板 (可选)](#第-4-步导入-graphana-面板-可选)

第 1 步：环境准备
---

1.部署监控系统的机器需要安装如下组件：

docker、docker-compose、jq、

- docker安装

  ```bash
  curl -fsSL get.docker.com -o get-docker.sh
  sudo sh get-docker.sh --mirror Aliyun
  ```

  或者直接安装

  ```bash
  apt-get install docker-ce
  apt-get install docker-ce-cli
  ```

- docker-compose

  ``` bash
  curl -L https://github.com/docker/compose/releases/download/1.18.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
  chmod +x /usr/local/bin/docker-compose
  ```

  或者直接安装

  ``` bash
  apt-get install docker-compose
  ```

- jq

  ```bash
  apt-get install jq
  ```

2. 准备监控

```bash
git clone git@github.com:opencurve/curve.git
mkdir -p /etc/curvefs/monitor
cp -r curve/curvefs/monitor /etc/curvefs/
cd /etc/curvefs/monitor
chmod -R 0777 grafana
chmod -R 0777 prometheus
```

**注意**：
prometheus 会保存数据到 /etc/curvefs/monitor/prometheus/data 目录，请保证空间足够。
如要使用其他目录，请修改 /etc/curvefs/monitor/docker-compose.yml 中的 prometheus/volume 下的 `- ./prometheus/data:/prometheus:rw`。
并将 prometheus 目录中的内容复制到修改后的内容。

第 2 步：准备配置文件
---

prometheus target.json 文件的生成依赖于工具 curve_ops_tool, 工具的运行需要配置文件 /etc/curve/tools.conf 。
该配置文件可以从 curveadm 部署的 docker 容器中获取。
在运行 curvefs（mds、metaserver）的机器上使用一下命令获取容器id（如：ae677e3c4b74）：
  
  ```bash
  docker ps

  CONTAINER ID   IMAGE                          COMMAND                  CREATED       STATUS       PORTS     NAMES
  ae677e3c4b74   opencurvedocker/curvefs:v1.2   "/entrypoint.sh --ro…"   3 hours ago   Up 3 hours             curvebs-chunkserver-8487da64a304
  43ba681f26f3   opencurvedocker/curvefs:v1.2   "/entrypoint.sh --ro…"   3 hours ago   Up 3 hours             curvebs-chunkserver-9bdcf13f6ec6
  67a7d64e5fbf   opencurvedocker/curvefs:v1.2   "/entrypoint.sh --ro…"   3 hours ago   Up 3 hours             curvebs-etcd-684a906fc55a
  93e98310e1a1   opencurvedocker/curvefs:v1.2   "/entrypoint.sh --ro…"   3 hours ago   Up 3 hours             curvebs-chunkserver-addd8d9120f9
  0bc77e13d604   opencurvedocker/curvefs:v1.2   "/entrypoint.sh --ro…"   3 hours ago   Up 3 hours             curvebs-mds-6b7274e16dca
  ```

使用 `docker cp` 从镜像中复制配置文件到本地（注意替换 container id）：

```bash
docker cp ae677e3c4b74:/etc/curvefs/tools.conf /etc/curvefs/tools.conf
```

第 3 步：部署监控系统
---

- 修改相关配置

  1. 修改 update_dashboard.sh，将 URL 和 LOGIN 改为对应的地址和用户名密码
  2. 修改 docker-compose.yml 文件，主要是映射的目录路径

- 启动docker-compose

```bash
cd /etc/curvefs/monitor
sudo bash curve-monitor.sh start
```

- 启动更新 target.json 服务：

```bash
# 请将 opencurvedocker/curvefs:v1.2 更换为自己部署的镜像
docker run -d -v /etc/curvefs/monitor/prometheus:/curvefs/monitor/prometheus -v /etc/curvefs/:/etc/curvefs/ opencurvedocker/curvefs:v1.2 entrypoint.sh --role=monitor
```

该服务更新容器内的 /curvefs/monitor/prometheus/target.json 文件，而 prometheus 需要读取该文件，从而抓去相应服务的监控数据。
因此需要将 /etc/curvefs/monitor/prometheus 映射到 /curvefs/monitor/prometheus。

**注意**：
更新 target.json 镜像（本文中为 opencurvedocker/curvefs:v1.2，请更换为自己的镜像）尽量以所部署的镜像为准，较老的镜像可能不支持相关服务，使用命令 `docker logs $(CONTAINER ID)` (CONTAINER ID 为`docker run -d`的输出) 会出现下列提示，请联系开发者。

```bash
Usage:
entrypoint.sh --role=ROLE
entrypoint.sh --role=ROLE --args=ARGS

Examples:
entrypoint.sh --role=etcd
entrypoint.sh --role=client --args="-o default_permissions"
```

访问 prometheus （ <http://127.0.0.1:9090/targets>）, 可以看到监控的服务。
**注意**：
请将 127.0.0.1:9090 更换为自己的ip和端口号，下同。

第 4 步：导入 graphana 面板 (可选)
--

默认从 monitor/grafana/provisioning/dashboards 导入，如需要更新可以按照以下步骤操作：

- 添加 prometheus 数据源

  访问 graphana 数据源设置页面（ http://127.0.0.1:3000/datasources ），选择 add data source
  填写 prometheus 对应的 url （ http://127.0.0.1:9090 ），保存即可。

- 添加面板

  访问 graphana 面板设置页面（ http://127.0.0.1:3000/dashboard/import ) 。
  所需的 json 文件从 [dashboards](https://github.com/opencurve/curve/tree/master/curvefs/monitor/grafana/provisioning/dashboards) 下载。
