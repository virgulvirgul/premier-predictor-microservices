const pino = require('pino');

const protoLoader = require('@grpc/proto-loader');
const grpc = require('grpc');

module.exports.error = (message, status) => {
    return {
        code: status,
        message: message
    };
};

module.exports.logger = pino({
    name: 'authservice-server',
    messageKey: 'message',
    changeLevelName: 'severity',
    useLevelLabels: true
});

module.exports.loadProto = (path) => {
    const packageDefinition = protoLoader.loadSync(
        path,
        {
            keepCase: true,
            longs: String,
            enums: String,
            defaults: true,
            oneofs: true
        }
    );
    return grpc.loadPackageDefinition(packageDefinition);
};