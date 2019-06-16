import * as pino from "pino";

export const logger = pino({
    name: 'fixtureservice-server',
    messageKey: 'message',
    changeLevelName: 'severity',
    useLevelLabels: true
});