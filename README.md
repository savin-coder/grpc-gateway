# grpc-gateway

You can get full details of working of grpc gateway in [OFFICIAL GUIDE](https://github.com/grpc-ecosystem/grpc-gateway) if you are interested. But In this file, I am going to explain how this gateway works.

## Working Of This File

Add your proto files in the folder proto, and write respective HTTP-RULES as defined in  [grpc-transcoding-tutorial](https://cloud.google.com/endpoints/docs/grpc/transcoding) from Google cloud. **Note:** You don't need to host in google cloud, but only for the http-rules defined to work with grpc-gateway.

After adding respective proto inside proto folder, and giving http-endpoints to your grpc services, generate the required go files by running command given below:
```bash
bash generate.sh
```

Required files will be generated in the folder **generated**

Now, on the **main.go** file, add the following code:

```bash
err := gw.Register<YOUR-GRPC-SERVICE-NAME>FromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}
```

The service name of your grpc server is: 
```bash
service <YOUR-GRPC-SERVICE-NAME> {
  rpc 1 {}
  rpc 2 {}
  .
  .
  .
  
}
```
Note: You can change your grpc server port and gateway port inside **main.go**
Now run the gateway by running main function in **main.go**, or simply by command:

```bash
go run main.go  
```

# Thank You !
