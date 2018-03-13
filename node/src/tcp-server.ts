import * as net from 'net';

import { ERROR_RESPONSE } from './constants';
import { rawCommandRequestToResponse } from './process-socket-connection';

export function startServer() {
  const server = net.createServer();

  server.on('close', (event: any) => {
  });

  server.on('connection', (socket: net.Socket) => {

    socket.on('data', async (data: Buffer) => {
      try {
        const stringRepresentation = data.toString();
        const response = await rawCommandRequestToResponse(stringRepresentation);
        socket.write(response);
      } catch (ex) {
        socket.write(ERROR_RESPONSE);
      }
    });

    socket.on('error', (err: Error) => {});
  });

  server.on('error', (err: Error) => {
      process.exit(1);
  });

  server.listen(8080, '0.0.0.0', 0, () => {
  });
}
