

import {
  INDEX_NUMBER,
  REMOVE_NUMBER,
  QUERY_NUMBER
} from './constants';
import {
  index,
  isIndexed,
  isDependedOn,
  remove
} from './data-access';

import { reduceBools } from './helpers';
import { Command } from './interfaces';
import { debug } from './logging';

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

export async function processIndexCommand(command: Command): Promise<boolean> {
  debug(`Starting to process index command for ${command.packageName}`);
  // okay, first thing we need to check is if all of the dependencies are indexed
  const promises = command.dependencies.map(dependency => isIndexed(dependency));
  const results = await Promise.all(promises);
  const allSucceeded = results.length ? reduceBools(results) : true;
  if (allSucceeded) {
    // sweet, we can go ahead and index it since all of the dependencies are good to go
    await index(command.packageName, command.dependencies);
    debug(`Successfully processed index command for ${command.packageName}`);
    return true;
  }
  debug(`Failed to process index command for ${command.packageName} due a dependency not being indexed`);
  return false;
}

export async function processRemoveCommand(command: Command): Promise<boolean> {
  debug(`Starting to process remove command for ${command.packageName}`);
  // okay, first thing we need to do is check if this command is a dependency of another command.
  // if it is, we cannot remove it.
  const dependedOn = await isDependedOn(command.packageName);
  if (dependedOn) {
    debug(`Failed to process remove command for ${command.packageName} since it's depended on by another package`);
    return false;
  }

  debug(`Sucessfully processed remove command for ${command.packageName}`);
  await remove(command.packageName);
  return true;
}

export function processQueryCommand(command: Command): Promise<boolean> {
  debug(`Starting to process query command for ${command.packageName}`);
  const result = isIndexed(command.packageName);
  debug(`Successfully processed query command for ${command.packageName}. The result is ${result}`);
  return result;
}