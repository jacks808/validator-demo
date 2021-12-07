
proto_dir=./protos
pb_dir=./proto

clean:
	rm -vf ${pb_dir}/*.pb.go

proto: clean
	protoc --proto_path=${proto_dir} --go_out=paths=source_relative:${pb_dir} ${proto_dir}/*.proto
	protoc-go-inject-tag -input="./${pb_dir}/*.pb.go"