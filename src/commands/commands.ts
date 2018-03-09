import {
  INDEX,
  INDEX_NUMBER,
  QUERY,
  QUERY_NUMBER,
  REMOVE,
  REMOVE_NUMBER,
  UNKNOWN_NUMBER
} from '../utils/constants';

import { isValidUnixCommand } from '../utils/helpers';

export function validateCommandFormat(rawCommand: string) {

  const lastChar = rawCommand.length ? rawCommand.charAt(rawCommand.length - 1) : null;
  if (lastChar !== '\n') {
    throw new Error('Usage: Commands must end in a newline character');
  }
  
  const commandChunks = rawCommand.split('|');
  if (commandChunks.length !== 3) {
    throw new Error('Usage: command|package|dep1,dep2,dep3,...');
  }

  const cleanedChunks = commandChunks.map(chunk => chunk.trim())
                          .filter(chunk => chunk.length > 0);


  const commandType = cleanedChunks[0];
  if (commandType !== INDEX
    && commandType !== REMOVE
    && commandType !== QUERY) {

    throw new Error(`Usage: First argument must be ${INDEX}, ${REMOVE} or ${QUERY}`);
  }

  if (!isValidUnixCommand(cleanedChunks[1])) {
    throw new Error('Usage: A command must consist of only letters, no spaces');
  }

  const commandObj: Command = {
    type: getTypeNumber(cleanedChunks[0]),
    packageName: cleanedChunks[1],
    dependencies: []
  };

  if (cleanedChunks.length > 2) {
    const rawDependencies = cleanedChunks[2].split(',');
    const dependencies = rawDependencies.map(dependency => dependency.trim())
                            .filter(dependency => dependency.length);

    if (rawDependencies.length !== dependencies.length) {
      throw new Error('Usage: One or more invalid dependency entries');
    }
    
    commandObj.dependencies = dependencies;
  }

  return commandObj;
}

function getTypeNumber(type: string) {
  if (type === INDEX) {
    return INDEX_NUMBER;
  } else if (type === REMOVE) {
    return REMOVE_NUMBER;
  } else if (type === QUERY) {
    return QUERY_NUMBER;
  }
  return UNKNOWN_NUMBER;
}

export interface Command {
  type: number;
  packageName: string;
  dependencies: string[];
}
