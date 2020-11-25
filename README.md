# plugin-poc

```bash
cd pd-apollo-basic-storage
go build -buildmode=plugin
cp pd-apollo-basic-storage.so ../storage.so
cd pd-apollo-basic-receiver
go build -buildmode=plugin
cp pd-apollo-basic-receiver.so ../receiver.so
cd ..
go run main.go
```
