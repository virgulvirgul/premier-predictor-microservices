import {NextFunction, Request, response, Response} from "express";

export class Middleware {
    private client;

    constructor() {
        const PROTO_PATH = __dirname + '/../../protodefs/auth.proto';
        const grpc = require('grpc');
        const protoLoader = require('@grpc/proto-loader');
        const packageDefinition = protoLoader.loadSync(
            PROTO_PATH,
            {
                keepCase: true,
                longs: String,
                enums: String,
                defaults: true,
                oneofs: true
            });
        const authProto = grpc.loadPackageDefinition(packageDefinition).model;

        const authAddr = process.env.AUTH_ADDR;

        this.client = new authProto.AuthService(authAddr, grpc.credentials.createInsecure());
    }

    public validateHttp(req: Request, res: Response, next: NextFunction) {
        const token = req.header("Authorization");

        this.client.validate({token: token}, (err, response) => {
            if (err) {
                res.status(401)
                    .send(err);
                return;
            }

            next();
        });
    }

    public validateGrpc(req: Request, res: Response, next: NextFunction) {
        const token = req.header("Authorization");

        this.client.validate({token: token}, (err, response) => {
            if (err) {
                res.status(401)
                    .send(err);
                return;
            }

            next();
        });
    }
}