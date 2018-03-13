import * as net from 'net';

import { ERROR_RESPONSE } from './constants';
import { debug, warn } from './logging';
import { rawCommandRequestToResponse } from './process-socket-connection';

export function startServer() {
  const server = net.createServer();

  server.on('close', (event: any) => {
    warn('Server shutting down');
  });

  server.on('connection', (socket: net.Socket) => {
    debug('A socket has connected');

    socket.on('data', async (data: Buffer) => {
      try {
        debug('Received data from the socket');
        const stringRepresentation = data.toString();
        const response = await rawCommandRequestToResponse(stringRepresentation);
        debug('Writing back to the socket');
        socket.write(response);
      } catch (ex) {
        warn(`Unexpected error occurred: ${ex.message}`);
        socket.write(ERROR_RESPONSE);
      }
    });

    socket.on('error', (err: Error) => {
      debug(`The socket experienced an error: ${err.message}`)
    });
  });

  server.on('error', (err: Error) => {
    warn(`The server experienced a fatal error: ${err.message}`);
    process.exit(1);
  });

  server.listen(8080, '0.0.0.0', 0, () => {
    warn('Server listening on port 8080');
  });
}
