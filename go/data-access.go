package main

import (
	"sync"
)

var mutex = &sync.Mutex{}
var commandDependenciesMap = make(map[string][]string)

// IsIndexed returns true if the package name is in the map, and false if it's not
func IsIndexed(packageName string) bool {
	Debugf("Starting IsIndexed for %s: ", packageName)
	mutex.Lock()
	result := commandDependenciesMap[packageName] != nil
	mutex.Unlock()
	Debugf("Done with IsIndexed for %s. The value is set to %t", packageName, result)
	return result
}

// IsDependedOn checks whether a package name is current dependend on and returns a corresponding boolean
func IsDependedOn(packageName string) bool {
	Debugf("Starting IsDependedOn for %s: ", packageName)
	mutex.Lock()
	for key, dependencies := range commandDependenciesMap {
		for i := range dependencies {
			if dependencies[i] == packageName {
				mutex.Unlock()
				Debugf("Done with IsDependedOn for %s. It is depended on by %s", packageName, key)
				return true
			}
		}
	}
	mutex.Unlock()
	Debugf("Done with IsDependedOn for %s. It's not depended on", packageName)
	return false
}

// IndexPackage indexes a package and it's dependencies in the data store
func IndexPackage(packageName string, dependencies []string) {
	Debugf("Starting IndexPackage for %s: ", packageName)
	mutex.Lock()
	commandDependenciesMap[packageName] = dependencies
	mutex.Unlock()
	Debugf("Done with IndexPackage for %s: ", packageName)
}

// RemovePackage removes a package and it's dependencies from the data store
func RemovePackage(packageName string) {
	Debugf("Starting RemovePackage for %s: ", packageName)
	mutex.Lock()
	delete(commandDependenciesMap, packageName)
	mutex.Unlock()
	Debugf("Done with RemovePackage for %s: ", packageName)
}

// WipeDataStore is a method exposed purely for testing
func WipeDataStore() {
	mutex.Lock()
	for packageName := range commandDependenciesMap {
		delete(commandDependenciesMap, packageName)
	}
	mutex.Unlock()
}

// GetDataStore is a method exposed purely for testing
func GetDataStore() map[string][]string {
	return commandDependenciesMap
}
