package com.cshep4.premierpredictor.notification.config;


import com.cshep4.premierpredictor.auth.AuthServiceGrpc;
import com.cshep4.premierpredictor.auth.AuthServiceGrpc.AuthServiceBlockingStub;
import io.grpc.ManagedChannelBuilder;
import lombok.val;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AuthService {
//    @Value("${AUTH_ADDR}")
//    private String authServiceAddress;
//
//    @Bean
//    public AuthServiceBlockingStub client() {
//        val channel = ManagedChannelBuilder
//                .forTarget(authServiceAddress)
//                .usePlaintext()
//                .build();
//
//        return AuthServiceGrpc.newBlockingStub(channel);
//    }
}
