package com.cshep4.premierpredictor.auth

import com.cshep4.premierpredictor.auth.repository.Repository
import com.cshep4.premierpredictor.auth.service.Service
import com.cshep4.premierpredictor.auth.transport.grpc.GrpcService
import com.cshep4.premierpredictor.auth.transport.http.Handler
import io.vertx.core.AbstractVerticle
import io.vertx.grpc.VertxServerBuilder
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch

class MainVerticle : AbstractVerticle() {

  override fun start() {
    val repository = Repository(vertx)
    val service = Service(repository)

    GlobalScope.launch {
      val grpcService = GrpcService(service)

      VertxServerBuilder
        .forPort(vertx, 50051)
        .addService(grpcService)
        .build()
        .start { ar ->
          if (ar.succeeded()) {
            System.out.println("gRPC service started")
          } else {
            ar.cause().printStackTrace()
            System.exit(1)
          }
        }
    }

    GlobalScope.launch {
      val handler = Handler(service)

      vertx.createHttpServer()
        .requestHandler(handler.route(vertx))
        .listen(8080) { res ->
          if (res.failed()) {
            res.cause().printStackTrace()
          } else {
            System.out.println("Server listening at: http://localhost:8080/")
          }
        }
    }

  }
}
