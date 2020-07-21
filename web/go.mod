module blog/web

go 1.14

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.4
	github.com/micro/examples/blog/posts v0.0.0-20200719203910-dd1bbe07bcbc
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/crypto v0.0.0 // indirect
	golang.org/x/lint v0.0.0 // indirect
	golang.org/x/net v0.0.0 // indirect
	golang.org/x/sync v0.0.0 // indirect
	golang.org/x/sys v0.0.0 // indirect
	golang.org/x/time v0.0.0 // indirect
	golang.org/x/tools v0.0.0 // indirect
	golang.org/x/xerrors v0.0.0 // indirect
	google.golang.org/genproto v0.0.0 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.0
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8
	golang.org/x/build => github.com/golang/build v0.0.0-20190311235527-86650285478d
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190306152737-a1d7652674e8
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net => github.com/golang/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190306144031-151b6387e3f2
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190312061237-fead79001313
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190311215038-5c2858a9cfe5
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api => github.com/googleapis/googleapis v0.0.0-20190312042308-abd1c9a99c5c
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc => github.com/grpc/grpc-go v1.26.0
)
