package options

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"
)

var ClientKeepAlive = grpc.WithKeepaliveParams(keepalive.ClientParameters{
	Time:                2 * time.Minute,
	Timeout:             20 * time.Second,
	PermitWithoutStream: true,
})

var ServerKeepAlive = grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
	MinTime:             time.Minute,
	PermitWithoutStream: true,
})
