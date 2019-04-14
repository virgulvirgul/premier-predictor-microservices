package com.cshep4.premierpredictor.notification.component;


import com.cshep4.premierpredictor.auth.AuthServiceGrpc.AuthServiceBlockingStub;
import com.cshep4.premierpredictor.auth.ValidateRequest;
import io.grpc.StatusRuntimeException;
import lombok.val;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class AuthServiceClient {
    @Autowired
    private AuthServiceBlockingStub client;

    public boolean validate(String token) {
        val req = ValidateRequest.newBuilder()
                .setToken(token)
                .build();

        try {
            client.validate(req);

            return true;
        } catch (StatusRuntimeException e) {
            return false;
        }
    }
}
