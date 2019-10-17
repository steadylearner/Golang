# How to follow the offical gRPC instruction 

You can use this example while you use go mod API. 

You can use other modules in the example inside import but you should use this command for protoc-gen-go.

```console
$go get -u github.com/golang/protobuf/protoc-gen-go // if it fails, for example, on Ubuntu use this command instead
$sudo apt  install golang-goprotobuf-dev
```

You can verify **protoc-gen-go** installed or not with **$protoc-gen-go** in your console.

## Notes

1. You should use your go module name(steadylearner/gRPC here) to find other files in your project.

2. You should install [protobuf](https://github.com/golang/protobuf) correctly to use **gRPC.bash**.

3. The code below doesn't seem to be relevant to current protobuf CLI anymore.

```golang
// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
```

Use Server made from protobuf without Unimplemented from the example and it will work.
