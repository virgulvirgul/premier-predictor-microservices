package com.cshep4.premierpredictor.auth.transport.grpc

import com.cshep4.premierpredictor.auth.service.Service
import io.grpc.examples.helloworld.GreeterGrpc
import io.grpc.examples.helloworld.HelloReply
import io.grpc.examples.helloworld.HelloRequest
import io.vertx.core.Future

class GrpcService(val service: Service) : GreeterGrpc.GreeterVertxImplBase() {
  override fun sayHello(req: HelloRequest, fut: Future<HelloReply>) {
    System.out.println("Hello " + req.name)
    fut.complete(
      HelloReply.newBuilder()
        .setMessage("Hi there, " + req.name)
        .build())
  }
}
