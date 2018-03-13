import {
  ERROR_RESPONSE,
  FAIL_RESPONSE,
  OK_RESPONSE
} from './constants'
import { debug } from './logging';
import { processCommand } from './process-command';
import { validateCommand } from './validate-command';

export async function rawCommandRequestToResponse(input: string): Promise<string> {
  debug(`Beginning to process request`);
  try {
    const command = validateCommand(input);
    const result = await processCommand(command);
    if (result) {
      debug(`Request was processed and was successful`);
      return OK_RESPONSE;
    }
    debug(`Request was processed and not successful`);
    return FAIL_RESPONSE;
  } catch (ex) {
    debug(`Request was processed and threw an exception: ${ex.message}`);
    return ERROR_RESPONSE;
  }
}