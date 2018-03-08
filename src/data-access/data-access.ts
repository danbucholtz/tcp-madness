

const commandDependenciesMap = new Map<string, string[]>();

export function getAllKeys() {
  const keys: string[] = [];
  commandDependenciesMap.forEach((deps, key) => {
    keys.push(key);
  });
  return keys;
}

export function isIndexed(packageName: string): Promise<boolean> {
  return new Promise((resolve) => {
    resolve(commandDependenciesMap.has(packageName));
  });
}

export function isDependedOn(packageName: string): Promise<boolean> {
  return new Promise((resolve) => {
    commandDependenciesMap.forEach((dependencies: string[]) => {
      for (const dependency of dependencies) {
        if (dependency === packageName) {
          return resolve(true);
        }
      }
    });
    resolve(false);
  });
}

export function index(packageName: string, dependencies: string[]): Promise<void> {
  return new Promise((resolve) => {
    commandDependenciesMap.set(packageName, dependencies);
    resolve();
  });
}

export function remove(packageName: string): Promise<boolean> {
  return new Promise((resolve) => {
    resolve(commandDependenciesMap.delete(packageName));
  });
}

// this method is used purely for testing
export function wipeDataStore() {
  const keysToDelete: string[] = [];
  commandDependenciesMap.forEach((value: string[], key: string) => {
    keysToDelete.push(key);
  });
  keysToDelete.forEach(key => commandDependenciesMap.delete(key));
}