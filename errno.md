CurveAdm 错误码
===

简介
---

CurveAdm 错误码是 Curve 团队为了提高用户部署成功率、指导用户解决部署中可能出现的各类问题而设计的全局唯一的跟踪码。

```yaml
---
Error-Code: 540000
Error-Description: port is already in use
Error-Clue: host=centos7-002, port=8202
How to Solve:
  * Website: https://github.com/opencurve/curveadm/wiki/errno#540000
  * Log: /root/.curveadm/logs/curveadm-2022-08-11_10-52-23.log
  * WeChat: opencurve_bot
```

错误码分类
---

错误码是一个 6 位固定长度的数字，每一个错误码由以下 3 部分组成：

 * 主类：主要错误类型
 * 子类：次要错误类型
 * 错误码：具体的错误码序列号

<img src="image/errno.jpg" width="60%" height="22%">

#### 查看错误码列表

```shell
$ curveadm -d
```

```yaml
000001 get user home dir failed
000002 create curveadm subdirectory failed
000003 init logger failed
100000 init SQLite database failed
110000 execute SQL failed which get hosts
110001 execute SQL failed which update hosts
111000 execute SQL failed which insert cluster
111001 execute SQL failed which get current cluster
111002 execute SQL failed which get cluster by name
111003 execute SQL failed which get all clusters
...
630007 remove container failed (docker rm CONTAINER)
630008 list containers failed (docker ps)
630009 run a command in container failed (docker exec CONTAINER COMMAND)
630010 copy file from container failed (docker cp CONTAINER:SRC_PATH DEST_PATH)
630011 copy file into container failed (docker cp SRC_PATH CONTAINER:DEST_PATH)
630012 get container low-level information failed (docker inspect ID)
630013 get container logs failed (docker logs ID)
690000 start crontab in container failed
900000 cancel operation
999999 unknown error
---
244 error codes
```


错误码列表
---

> :hourglass_flowing_sand: 错误码解释及解决方案正在编写中，将在近期发布