
import { debug } from './logging';

const commandDependenciesMap = new Map<string, string[]>();

export async function isIndexed(packageName: string): Promise<boolean> {
  debug(`Starting isIndexed for package ${packageName}`)
  const result = commandDependenciesMap.has(packageName);
  debug(`Done with isIndexed for package ${packageName}`)
  return result;
}

export async function isDependedOn(packageName: string): Promise<boolean> {
  debug(`Starting isDependedOn for package ${packageName}`)
  let found = false;
  commandDependenciesMap.forEach((dependencies: string[], key: string) => {
    for (const dependency of dependencies) {
      if (dependency === packageName) {
        debug(`${packageName} isDependedOn by ${key}}`);
        found = true;
      }
    }
  });
  debug(`Done with isDependedOn for package ${packageName}`)
  return found;
}

export async function index(packageName: string, dependencies: string[]): Promise<void> {
  debug(`Starting index for ${packageName}`)
  commandDependenciesMap.set(packageName, dependencies);
  debug(`Done with index for ${packageName}`)
}

export async function remove(packageName: string): Promise<boolean> {
  debug(`Starting remove for ${packageName}`)
  const result = commandDependenciesMap.delete(packageName);
  debug(`Done with remove for ${packageName}`)
  return result;
}

// this method is used purely for testing
export function wipeDataStore() {
  const keysToDelete: string[] = [];
  commandDependenciesMap.forEach((value: string[], key: string) => {
    keysToDelete.push(key);
  });
  keysToDelete.forEach(key => commandDependenciesMap.delete(key));
}