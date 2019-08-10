package com.cshep4.premierpredictor.auth.repository

import io.vertx.core.Vertx
import io.vertx.core.json.JsonObject
import io.vertx.ext.mongo.MongoClient

class Repository(vertx: Vertx) {
  private var client: MongoClient

  init {
    val mongoUri: String = System.getenv("MONGO_URI") ?: "mongodb://localhost:27017"
    val config = JsonObject(
      mapOf(
        Pair("connection_string", mongoUri),
        Pair("db_name", "test")
      )
    )
    client = MongoClient.createShared(vertx, config)
  }
}
