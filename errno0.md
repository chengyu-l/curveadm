## 000001

**类别：** [初始化][init]

**描述：** 获取用户主目录失败

**如何解决：** 用户需确认当前用户是否拥有主目录，并且得确保当前执行用户即为 CurveAdm 的安装用户，切勿添加 `sudo` 或以其他用户执行不属于当前用户的 CurveAdm。此外，用户也可以通过错误线索和查看错误日志来锁定相关问题，并进行排除

示例：

```shell
$ sudo curveadm # 错误，切勿添加 sudo 执行 curveadm
```

## 000002

**类别：** [初始化][init]

**描述：** 创建 `CurveAdm` 项目子目录失败

**如何解决：** 用户需确保当前 CurveAdm 的安装目录的权限依旧属于当前用户，详见[CurveAdm 安装目录][insatall-directory]。此外，用户也可以通过错误线索和查看错误日志来锁定相关问题，并进行排除

## 000003

**类别：** [初始化][init]

**描述：** 初始化日志模块失败

**如何解决：** 我们已知的报告该错误码的情况有以下几个：

* CurveAdm 安装目录所在磁盘的空间已被用尽，详见[CurveAdm 安装目录][insatall-directory]
* CurveAdm 安装目录的权限不属于当前用户，详见[CurveAdm 安装目录][insatall-directory]

用户需要通过错误线索和查看错误日志来锁定具体的问题，并进行排除

[init]: https://github.com/opencurve/curveadm/wiki/errno#主类
[database]: https://github.com/opencurve/curveadm/wiki/errno#主类
[command-options]: https://github.com/opencurve/curveadm/wiki/errno#主类
[configure]: https://github.com/opencurve/curveadm/wiki/errno#主类
[common-logic]: https://github.com/opencurve/curveadm/wiki/errno#主类
[execute-command]:  https://github.com/opencurve/curveadm/wiki/errno#主类
[cluster-options]: https://github.com/opencurve/curveadm/wiki/errno#子类
[client-common-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[curvebs-client-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[curvefs-client-options]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[playground-options]: https://github.com/opencurve/curveadm/wiki/errno#子类
[configure-common]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curveadm-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[hosts-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[topology-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[format-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[client-configure]: https://github.com/opencurve/curveadm/wiki/errno#子类
[hosts-logic]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[service-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curvefs-client-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[curvebs-client-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[polarfs-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[playground-logic]: https://github.com/opencurve/curveadm/wiki/errno#子类
[topology-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[ssh-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[permission-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[kernel-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[network-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[date-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[service-precheck]: https://github.com/opencurve/curveadm/wiki/errno#子类
[command-common]: https://github.com/opencurve/curveadm/wiki/errno#子类
[others-command]: https://github.com/opencurve/curveadm/wiki/errno#子类
[docker-command]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[ssh-connection]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[shell-command]:  https://github.com/opencurve/curveadm/wiki/errno#子类
[database-common-solution]: https://github.com/opencurve/curveadm/wiki/errno#数据库类错误码
[insatall-directory]: https://github.com/opencurve/curveadm/wiki/overview#安装目录
[precheck]: https://github.com/opencurve/curveadm/wiki/precheck
[clean-cluster]: https://github.com/opencurve/curveadm/wiki/maintain-curve#清理集群
[map]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#第-4-步映射-curvebs-卷
[mount]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#第-4-步挂载-curvefs-文件系统
[variable]: https://github.com/opencurve/curveadm/wiki/topology#变量
[curveadm-configure-file]: https://github.com/opencurve/curveadm/wiki/install-curveadm#curveadm-配置文件
[hosts]: https://github.com/opencurve/curveadm/wiki/hosts
[topology]: https://github.com/opencurve/curveadm/wiki/topology
[replicas]: https://github.com/opencurve/curveadm/wiki/topology#replicas
[scale-curve]: https://github.com/opencurve/curveadm/wiki/scale-curve
[migrate-curve]: https://github.com/opencurve/curveadm/wiki/migrate-curve
[format-disk]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment#第-4-步格式化磁盘
[curvebs-client-configure]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#第-3-步准备客户端配置文件
[curvefs-client-configure]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#第-3-步准备客户端配置文件
[add-and-checkout]: https://github.com/opencurve/curveadm/wiki/curvebs-cluster-deployment#第-6-步添加集群并切换集群
[umap]: https://github.com/opencurve/curveadm/wiki/curvebs-client-deployment#其他取消-curvebs-卷映射
[umount]: https://github.com/opencurve/curveadm/wiki/curvefs-client-deployment#其他卸载文件系统
[chunkfilepool]: https://github.com/opencurve/curve/blob/master/docs/cn/chunkserver_design.md#224-chunkfilepool
[config-template]: https://github.com/opencurve/curveadm/tree/master/configs
[execute-command-common-solution]: https://github.com/opencurve/curveadm/wiki/errno#命令执行类错误码
[start-docker-daemon]: https://yeasy.gitbook.io/docker_practice/install/ubuntu#qi-dong-docker
[issue]: https://github.com/opencurve/curveadm/issues
[topology]: https://github.com/opencurve/curveadm/wiki/topology
[install-requirements]: https://github.com/opencurve/curveadm/wiki/install-curveadm#安装依赖
[important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvebs-important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvebs-重要配置项
[curvefs-important-configure]: https://github.com/opencurve/curveadm/wiki/topology#curvefs-重要配置项
