# jsocial

### 服务预览
| 服务           | 预览地址                          |
|:-------------|:------------------------------|
| consul       | http://localhost:8500/        |
| grafana      | http://localhost:3000/        |
| jaeger       | http://localhost:16686/search |
| Prometheus   | http://localhost:9090/graph   |
| AlertManager | http://localhost:9093         |


### GO配置环境变量
```shell
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/home/jidan/go/bin
```
### WSL中通过Docker启动Mysql报错
#### 错误
```text
 jidan@jidan: docker-compose logs deployments-db-1
 deployments-db-1   | mysqld: Cannot change permissions of the file 'private_key.pem' (OS errno 1 - Operation not permitted)
 deployments-db-1   | 2020-09-10T05:04:53.449233Z 0 [ERROR] [MY-010295] [Server] Could not set file permission for private_key.pem
 deployments-db-1   | 2020-09-10T05:04:53.449629Z 0 [ERROR] [MY-010119] [Server] Aborting
```
#### 解决办法

* 创建wsl配置文件
```shell
 jidan@jidan: sudo touch /etc/wsl.conf
```
* 编辑wsl配置文件
```shell
 jidan@jidan: sudo vim /etc/wsl.conf
```
* 添加内容
```shell
 [automount]
 options = "metadata"
```
* 退出wsl
```shell
 jidan@jidan:  exit
```
* 关闭wsl
```shell
 wsl --shutdown
```
* 重新进入wsl
```shell
 wsl
```

