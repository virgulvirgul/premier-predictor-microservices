package com.cshep4.premierpredictor

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.autoconfigure.domain.EntityScan
import org.springframework.boot.runApplication
import org.springframework.scheduling.annotation.EnableScheduling


@SpringBootApplication
@EnableScheduling
@EntityScan(basePackages = ["com.cshep4.premierpredictor.entity"])
class PremierPredictorApplication

fun main(args: Array<String>) {
    runApplication<PremierPredictorApplication>(*args)
}