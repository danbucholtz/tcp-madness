
import {
  Command
} from './commands/commands';

import {
  getAllKeys,
  index,
  isIndexed,
  isDependedOn,
  remove
} from './data-access/data-access';

import {
  INDEX_NUMBER,
  REMOVE_NUMBER,
  QUERY_NUMBER
} from './utils/constants';

import {
  reduceBools
} from './utils/helpers';

export function processCommand(command: Command): Promise<boolean> {
  if (command.type === INDEX_NUMBER) {
    return processIndexCommand(command);
  } else if (command.type === REMOVE_NUMBER) {
    return processRemoveCommand(command);
  } else if (command.type === QUERY_NUMBER) {
    return processQueryCommand(command);
  }
  throw new Error('Unknown Command Type');
}

export function processIndexCommand(command: Command): Promise<boolean> {
  // okay, first thing we need to check is if all of the dependencies are indexed
  const promises = command.dependencies.map(dependency => isIndexed(dependency));
  return Promise.all(promises).then((results: boolean[]) => {
    const allSucceeded = results.length ? reduceBools(results) : true;
    if (allSucceeded) {
      // sweet, we can go ahead and index it since all of the dependencies are good to go
      return index(command.packageName, command.dependencies).then(() => {
        return true;
      })
    }
    return Promise.resolve(false);
  });
}

export function processRemoveCommand(command: Command): Promise<boolean> {
  // okay, first thing we need to do is check if this command is a dependency of another command.
  // if it is, we cannot remove it.
  return isDependedOn(command.packageName).then((dependedOn) => {
    
    if (dependedOn) {
      return false;
    }

    // it's not depended on, so we can remove it
    
    return remove(command.packageName).then(() => {
      return true;
    });
  });
}

export function processQueryCommand(command: Command): Promise<boolean> {
  return isIndexed(command.packageName);
}