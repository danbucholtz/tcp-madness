

const commandDependenciesMap = new Map<string, string[]>();

export async function isIndexed(packageName: string): Promise<boolean> {
  return commandDependenciesMap.has(packageName);
}

export async function isDependedOn(packageName: string): Promise<boolean> {
  let found = false;
  commandDependenciesMap.forEach((dependencies: string[], key: string) => {
    for (const dependency of dependencies) {
      if (dependency === packageName) {
        found = true;
      }
    }
  });
  return found;
}

export async function index(packageName: string, dependencies: string[]): Promise<void> {
  commandDependenciesMap.set(packageName, dependencies);
}

export async function remove(packageName: string): Promise<boolean> {
  return commandDependenciesMap.delete(packageName);
}

// this method is used purely for testing
export function wipeDataStore() {
  const keysToDelete: string[] = [];
  commandDependenciesMap.forEach((value: string[], key: string) => {
    keysToDelete.push(key);
  });
  keysToDelete.forEach(key => commandDependenciesMap.delete(key));
}