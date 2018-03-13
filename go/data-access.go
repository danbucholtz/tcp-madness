package main

import (
	"sync"
)

var rwMutex = &sync.RWMutex{}
var commandDependenciesMap = make(map[string][]string, 1000)

// IsIndexed returns true if the package name is in the map, and false if it's not
func IsIndexed(packageName string) bool {
	Debugf("Starting IsIndexed for %s: ", packageName)
	rwMutex.RLock()
	result := commandDependenciesMap[packageName] != nil
	rwMutex.RUnlock()
	Debugf("Done with IsIndexed for %s. The value is set to %t", packageName, result)
	return result
}

// IsDependedOn checks whether a package name is current dependend on and returns a corresponding boolean
func IsDependedOn(packageName string) bool {
	Debugf("Starting IsDependedOn for %s: ", packageName)
	rwMutex.RLock()
	for key, dependencies := range commandDependenciesMap {
		for i := range dependencies {
			if dependencies[i] == packageName {
				rwMutex.RUnlock()
				Debugf("Done with IsDependedOn for %s. It is depended on by %s", packageName, key)
				return true
			}
		}
	}
	rwMutex.RUnlock()
	Debugf("Done with IsDependedOn for %s. It's not depended on", packageName)
	return false
}

// IndexPackage indexes a package and it's dependencies in the data store
func IndexPackage(packageName string, dependencies []string) {
	Debugf("Starting IndexPackage for %s: ", packageName)
	rwMutex.Lock()
	commandDependenciesMap[packageName] = dependencies
	rwMutex.Unlock()
	Debugf("Done with IndexPackage for %s: ", packageName)
}

// RemovePackage removes a package and it's dependencies from the data store
func RemovePackage(packageName string) {
	Debugf("Starting RemovePackage for %s: ", packageName)
	rwMutex.Lock()
	delete(commandDependenciesMap, packageName)
	rwMutex.Unlock()
	Debugf("Done with RemovePackage for %s: ", packageName)
}

// WipeDataStore is a method exposed purely for testing
func WipeDataStore() {
	rwMutex.Lock()
	for packageName := range commandDependenciesMap {
		delete(commandDependenciesMap, packageName)
	}
	rwMutex.Unlock()
}

// GetDataStore is a method exposed purely for testing
func GetDataStore() map[string][]string {
	return commandDependenciesMap
}
