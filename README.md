# statistics
statistics

### Install
```shell
go install github.com/juxuny/statistics/cmd/statistics-cli
```

### Run
每天13:00 - 16:00收集数据，保存到MySQL,些前请先把 `res`目录里两个SQL文件导入(`res/index_list.sql`, `res/stock_code.sql`)
```shell
/opt/gopath/bin/collect -u root -p 123456 -d=true -log=/root/2.log -start=13:00 -end=16:00
```