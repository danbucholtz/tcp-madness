import * as net from 'net';

import { processConnection } from './process-socket-connection';

const server = net.createServer();

server.on('close', (event: any) => {
    console.log('close: ', event);
});

server.on('connection', (socket: net.Socket) => {
    console.log('Connection achieved!');

    socket.on('data', (data: Buffer) => {
        const stringRepresentation = data.toString();
        processConnection(stringRepresentation).then((toRespond) => {
            socket.write(toRespond);
        }).catch((ex: any) => {
          console.log('Unknown ERROR: ', ex);
        });
    });
});

server.listen(8080, '0.0.0.0', 0, () => {
    console.log('Listening for connections');
});

