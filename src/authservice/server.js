const path = require('path');
const grpc = require('grpc');
const fs = require('fs');

const MAIN_PROTO_PATH = path.join(__dirname, './protodefs/auth.proto');

const loadProto = require('./util/util').loadProto;

const logger = require('./util/util').logger;

const PORT = process.env.PORT;

const validate = require('./auth/validate').validate;

const startHealthServer = require('./health/server').startHealthServer;

function main() {
    logger.info(`Starting gRPC server on port ${PORT}...`);
    const opts = {
        "grpc.keepalive_time_ms": 60000,
        "grpc.keepalive_permit_without_calls" : 1
    };
    const server = new grpc.Server(opts);

    const authProto = loadProto(MAIN_PROTO_PATH).model;

    server.addService(authProto.AuthService.service, {
        validate,
    });

    // const sslCreds = grpc.ServerCredentials.createSsl(null, [{
    //     private_key: fs.readFileSync('./certs/tls.key'),
    //     cert_chain: fs.readFileSync('./certs/tls.crt')
    // }], true,);
    //
    // server.bind(`0.0.0.0:${PORT}`, sslCreds);
    server.bind(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure());
    server.start();

    startHealthServer();
}

main();

