
import {
  validateCommandFormat
} from './commands/commands';

import { processCommand } from './process-command';

export async function processConnection(input: string): Promise<string> {
  try {
    const command = validateCommandFormat(input);
    const result = await processCommand(command);
    if (result) {
      return 'OK\n';
    }
    return 'FAIL\n';
  } catch (ex) {
    return 'ERROR\n';
  }
}