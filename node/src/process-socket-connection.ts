import {
  ERROR_RESPONSE,
  FAIL_RESPONSE,
  OK_RESPONSE
} from './constants'
import { processCommand } from './process-command';
import { validateCommand } from './validate-command';

export async function rawCommandRequestToResponse(input: string): Promise<string> {
  try {
    const command = validateCommand(input);
    const result = await processCommand(command);
    if (result) {
      return OK_RESPONSE;
    }
    return FAIL_RESPONSE;
  } catch (ex) {
    return ERROR_RESPONSE;
  }
}