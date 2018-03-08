import * as net from 'net';

const server = net.createServer();

server.on('close', (event: any) => {
    console.log('close: ', event);
});

server.on('connection', (socket: net.Socket) => {
    console.log('Connection achieved!');

    socket.on('data', (data: Buffer) => {
        const stringRepresentation = data.toString();
        console.log('stringRepresentation: ', stringRepresentation);
    });
});

server.listen(8080, '0.0.0.0', 0, () => {
    console.log('Listening for connections');
});