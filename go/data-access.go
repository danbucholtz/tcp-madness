package main

import (
	"sync"
)

var mutex = &sync.Mutex{}
var commandDependenciesMap = make(map[string][]string)

// IsIndexed returns true if the package name is in the map, and false if it's not
func IsIndexed(packageName string) bool {
	mutex.Lock()
	result := commandDependenciesMap[packageName] != nil
	mutex.Unlock()
	return result
}

// IsDependedOn checks whether a package name is current dependend on and returns a corresponding boolean
func IsDependedOn(packageName string) bool {
	mutex.Lock()
	for _, dependencies := range commandDependenciesMap {
		for i := range dependencies {
			if dependencies[i] == packageName {
				mutex.Unlock()
				return true
			}
		}
	}
	mutex.Unlock()
	return false
}

// IndexPackage indexes a package and it's dependencies in the data store
func IndexPackage(packageName string, dependencies []string) {
	mutex.Lock()
	commandDependenciesMap[packageName] = dependencies
	mutex.Unlock()
}

// RemovePackage removes a package and it's dependencies from the data store
func RemovePackage(packageName string) {
	mutex.Lock()
	delete(commandDependenciesMap, packageName)
	mutex.Unlock()
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
