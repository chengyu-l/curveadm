使用 CurveAdm 部署 CurveFS 集群
===

* [第 1 步：环境准备](#第-1-步环境准备)
* [第 2 步：在中控机上安装 CurveAdm](#第-2-步在中控机上安装-curveadm)
* [第 3 步：部署 Minio（可选）](#第-3-步部署-minio可选)
* [第 4 步：准备集群拓扑文件](#第-4-步准备集群拓扑文件)
* [第 5 步：添加集群并切换集群](#第-5-步添加集群并切换集群)
* [第 6 步：部署集群](#第-6-步部署集群)
* [第 7 步：查看集群运行情况](#第-7-步查看集群运行情况)
* [第 8 步：验证集群健康状态](#第-8-步验证集群健康状态)

第 1 步：环境准备
---

* [软硬件环境需求](install-curveadm#软硬件环境需求)
* [安装依赖](install-curveadm#安装依赖)

第 2 步：在中控机上安装 CurveAdm
---

* [安装 CurveAdm](install-curveadm#安装-curveadm)

第 3 步：部署 Minio（可选）
---

该步骤为可选步骤。

由于目前 CurveFS 只支持 S3 作为后端存储，CurveBS 后端即将支持。
所以你需要部署一个 S3 存储或使用公有云对象存储，如亚马逊 S3、阿里云 OSS、腾讯云 OSS 等。
下面将展示如果利用 Docker 快速部署一个 [Minio][minio] 来作为 S3 后端存储：

```shell
$ mkdir minio-data
$ sudo docker run -d --name minio \
-p 9000:9000 \
-p 9900:9900 \
-v minio-data:/data \
--restart unless-stopped \
minio/minio server /data --console-address ":9900"
```
> 📢 **注意：**
>
> 运行参数中的 minio-data 为本地路径，你需要在运行 minio 容器之前，提前创建这个目录

> :bulb: **提醒：**
> 
> 以下这些信息将用于 CurveFS 拓扑文件中 [S3 相关配置项][important-config]的填写：
> * root 用户默认的 `Access Key` 以及 `Secret Key` 都为 `minioadmin`
> * S3 服务的访问地址为 `http://$IP:9000`， 你需要通过浏览器访问 `http://$IP:9000` 来创建一个桶
> * 关于部署的更多详细信息，你可以参考 [deploy-minio-standalone][deploy-minio-standalone]

第 4 步：准备集群拓扑文件
---

我们根据常见的场景，给用户准备了不同的拓扑文件模板，用户可根据需求自行选择，并进行编辑调整：

* [单机部署][curvefs-stand-alone-topology]

  所有服务都运行在一台主机，一般用于体验或测试


* [多机部署][curvefs-cluster-topology]

  通用的多机部署的基础模板，可用于生产环境或测试

关于拓扑文件中的各项配置项，请参考 [CurveFS 集群拓扑][curvefs-topology]。

```shell
$ vim topology.yaml
```

```yaml
kind: curvefs
global:
  user: curve
  ssh_port: 22
  private_key_file: /home/curve/.ssh/id_rsa
  container_image: opencurvedocker/curvefs:latest
  log_dir: /home/${user}/curvefs/logs/${service_role}
  data_dir: /home/${user}/curvefs/data/${service_role}
  s3.ak: <minioadmin>
  s3.sk: <minioadmin>
  s3.endpoint: <http://127.0.0.1:9000>
  s3.bucket_name: <>
  variable:
    machine1: 10.0.1.1
    machine2: 10.0.1.2
    machine3: 10.0.1.3

etcd_services:
  config:
    listen.ip: ${service_host}
    listen.port: 2380
    listen.client_port: 2379
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}

mds_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6700
    listen.dummy_port: 7700
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}

metaserver_services:
  config:
    listen.ip: ${service_host}
    listen.port: 6800
    listen.external_port: 7800
    metaserver.loglevel: 0
  deploy:
    - host: ${machine1}
    - host: ${machine2}
    - host: ${machine3}
      config:
        metaserver.loglevel: 3
```

第 5 步：添加集群并切换集群
---

#### 1. 添加 'my-cluster' 集群，并指定集群拓扑文件

```shell
$ curveadm cluster add my-cluster -f topology.yaml
```

#### 2. 切换 'my-cluster' 集群为当前管理集群

```shell
$ curveadm cluster checkout my-cluster 
```

第 6 步：部署集群
---

```shell
$ curveadm deploy
```

如果部署成功，将会输出类似 `Cluster 'my-cluster' successfully deployed ^_^.` 的字样。

第 7 步：查看集群运行情况
---

```shell
$ curveadm status
```

CurveAdm 默认会显示服务 ID、服务角色、主机地址、已部署的副本服务数量、容器 ID、运行状态：

```shell
Get Service Status: [OK]

cluster name    : my-cluster
cluster kind    : curvefs
cluster mds addr: 10.0.1.1:6700,10.0.1.2:6700,10.0.1.3:6700

Id            Role           Host      Replica  Container Id  Status
--            ----           ----      -------  ------------  ------
c9570c0d0252  etcd           10.0.1.1  1/1      ced84717bf4b  Up 45 hours
493b7831907c  etcd           10.0.1.2  1/1      907f8b84f527  Up 45 hours
8438cc5ecb52  etcd           10.0.1.3  1/1      44eca4798424  Up 45 hours
505da008b59c  mds            10.0.1.1  1/1      37c05bbb39af  Up 45 hours
e7bfb934182b  mds            10.0.1.2  1/1      044b56281928  Up 45 hours
1b322781339c  mds            10.0.1.3  1/1      b00481b9872d  Up 45 hours
2912bbdbcb48  metaserver     10.0.1.1  1/1      8b7a14b872ff  Up 45 hours
b862ef6720ed  metaserver     10.0.1.2  1/1      8e2a4b9e16b4  Up 45 hours
ed4533e903d9  metaserver     10.0.1.3  1/1      a35c30e3143d  Up 45 hours
```

* 若想查看其余信息，如日志目录、数据目录等，可添加 `-v` 参数
* 对于同一台主机上的[副本][replicas]服务来说，其状态默认是折叠的，可添加 `-s` 参数来显示每一个副本服务

第 8 步：验证集群健康状态
---

集群服务正常运行，并不意味着集群的健康，所以我们在每一个容器内内置了 `curvefs_tool` 工具。
该工具不仅可以查询集群的健康状态，还提供了许多其他特性，如显示各组件详细状态、创建/删除文件系统等。

首先，我们需要进入任意一个服务容器内（服务 ID 可通过 `curveadm status` 查看）：

```shell
$ curveadm enter <Id>
```

在该容器内执行以下命令查看：
```shell
$ curvefs_tool status
```

如果集群健康，在输出的最后会出现 `cluster is healthy` 的字样。

[minio]: https://github.com/minio/minio
[deploy-minio-standalone]: https://docs.min.io/minio/baremetal/installation/deploy-minio-standalone.html
[important-config]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-重要配置项
[curvefs-stand-alone-topology]: https://github.com/opencurve/curveadm/blob/master/configs/fs/stand-alone/topology.yaml
[curvefs-cluster-topology]: https://github.com/opencurve/curveadm/blob/master/configs/fs/cluster/topology.yaml
[curvefs-topology]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-集群拓扑 
[replicas]: https://github.com/opencurve/curveadm/wiki/topology#replica 