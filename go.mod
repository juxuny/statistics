module github.com/juxuny/bridge

require (
	cloud.google.com/go v0.34.0 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20180710155616-bc664df96737 // indirect
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20181014144952-4e0d7dc8888f // indirect
	github.com/elazarl/goproxy v0.0.0-20181111060418-2ce16c963a8a // indirect
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gofrs/uuid v3.1.0+incompatible // indirect
	github.com/gomarkdown/markdown v0.0.0-20190222000725-ee6a7931a1e4 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/iGoogle-ink/gopay v1.3.6 // indirect
	github.com/importcjj/sensitive v0.0.0-20190611120559-289e87ec4108 // indirect
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/silenceper/wechat v1.0.0 // indirect
	github.com/steakknife/hamming v0.0.0-20180906055917-c99c65617cd3 // indirect
	github.com/tidwall/gjson v1.2.1 // indirect
	github.com/tidwall/match v1.0.1 // indirect
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	google.golang.org/appengine v1.4.0 // indirect
	gopkg.in/appleboy/gofight.v2 v2.0.0-00010101000000-000000000000 // indirect
	qiniupkg.com/x v7.0.8+incompatible // indirect
)


replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/net => github.com/golang/net v0.0.0-20180724234803-3673e40ba225
	golang.org/x/text => github.com/golang/text v0.3.0
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	gopkg.in/appleboy/gofight.v2 => github.com/appleboy/gofight/v2 v2.1.2
)

go 1.13
