mkdir stock order

mkdir api rpc model


go run stock.go -f etc/stock.yaml
go run order.go -f etc/order.yaml

go run main.go -c conf.yml