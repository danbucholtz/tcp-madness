import * as net from 'net';

import { processConnection } from './process-socket-connection';

const server = net.createServer();

server.on('close', (event: any) => {
    console.log('close: ', event);
});

server.on('connection', (socket: net.Socket) => {

    socket.on('data', (data: Buffer) => {
        const stringRepresentation = data.toString();
        processConnection(stringRepresentation).then((toRespond) => {
            // console.log(`Sending Back ${toRespond} in response to ${stringRepresentation}`)
            socket.write(toRespond);
        }).catch((ex: any) => {
          //console.log('Unknown ERROR: ', ex);
        });
    });

    socket.on('error', (err: Error) => {
        console.log('Socket error: ', err);
    });
});

server.on('error', (err: Error) => {
    console.log('Server error: ', err);
});

server.listen(8080, '0.0.0.0', 0, () => {
    console.log('Listening for connections');
});

