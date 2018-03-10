import * as net from 'net';

import { processConnection } from './process-socket-connection';

const server = net.createServer();

server.on('close', (event: any) => {
});

server.on('connection', (socket: net.Socket) => {

    socket.on('data', (data: Buffer) => {
        const stringRepresentation = data.toString();
        processConnection(stringRepresentation).then((toRespond) => {
            socket.write(toRespond);
        }).catch((ex: any) => {
          socket.write('ERROR\n');
        });
    });

    socket.on('error', (err: Error) => {
    });
});

server.on('error', (err: Error) => {
    process.exit(1);
});

server.listen(8080, '0.0.0.0', 0, () => {
});

