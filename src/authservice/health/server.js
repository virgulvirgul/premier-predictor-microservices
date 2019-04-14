const express = require('express');
const app = express();
const port = process.env.HEALTH_PORT;

const logger = require('../util/util').logger;

module.exports.startHealthServer = () => {
    app.get('/_health', (req, res) => res.send('Service is healthy!'));
    app.listen(port, () => logger.info(`Starting health server on port ${port}...`));
};