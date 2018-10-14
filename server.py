from concurrent import futures
import ask_pb2, ask_pb2_grpc
import time
import grpc

class OracleServicer(ask_pb2_grpc.OracleServicer):
    def Reveal(self, request, context):
        return ask_pb2.Answer(answer="you said {q}!".format(q=request.question), link="")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ask_pb2_grpc.add_OracleServicer_to_server(
        OracleServicer(), server
    )
    server.add_insecure_port("localhost:10000")
    server.start()
    while True:
        print("Mock server is running!")
        time.sleep(60)

if __name__ == "__main__":
    serve()