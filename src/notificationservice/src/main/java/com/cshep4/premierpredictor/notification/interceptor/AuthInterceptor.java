package com.cshep4.premierpredictor.notification.interceptor;

import com.cshep4.premierpredictor.notification.component.AuthServiceClient;
import io.grpc.*;
import lombok.val;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;

@Component
public class AuthInterceptor implements ServerInterceptor {
    private static final Metadata.Key<String> AUTHORISATION = Metadata.Key.of("Authorisation", ASCII_STRING_MARSHALLER);

    @Autowired
    public AuthServiceClient authServiceClient;

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata metadata, ServerCallHandler<ReqT, RespT> next) {
        val token = metadata.get(AUTHORISATION);

        if (!authServiceClient.validate(token)) {
            call.close(Status.UNAUTHENTICATED.withDescription("JWT Token is missing from Metadata"), metadata);
            return new ServerCall.Listener<ReqT>() {};
        }

        return next.startCall(call, metadata);
    }
}
