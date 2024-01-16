import grpc # grpcio
import helloworld_pb2
import helloworld_pb2_grpc

def run():
    print("Will try to greet world ...")
    with grpc.insecure_channel("localhost:50051") as channel: # konek ke server
        stub = helloworld_pb2_grpc.GreeterStub(channel)
        # manggil function yang di server
        response = stub.SayHello(helloworld_pb2.HelloRequest(name="Heri"))
    print("Greeter client received: " + response.message)


run()
