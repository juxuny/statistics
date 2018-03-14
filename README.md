# statistics
statistics


## Install
```shell
mkdir /opt/gopath
export GOPATH=/opt/gopath
go install github.com/juxuny/statistics/cmd/statistics-cli
go install github.com/juxuny/statistics/cmd/collect
```

## 数据收集


#### Run
每天13:00 - 16:00收集数据，保存到MySQL,些前请先把 `res`目录里两个SQL文件导入(`res/index_list.sql`, `res/stock_code.sql`)
```shell
/opt/gopath/bin/collect -u root -p 123456 -d=true -log=/root/2.log -start=13:00 -end=16:00
```


## 数据导出


导出已经标准化的数据到csv文件
```shell
/opt/gopath/bin/statistics-cli -d -log="" -m pre -out=F:\tmp\pre-process -start=2017-11-30 -end=2017-12-07 -host=127.0.0.1 -p=123456 -u=root -code=sh600018
```


导出某一天所有数据
```shell
/opt/gopath/bin/statistics-cli -d -log="" -m export -out=F:\tmp -date=2017-11-30 -host=127.0.0.1 -p=123456 -u=root
```