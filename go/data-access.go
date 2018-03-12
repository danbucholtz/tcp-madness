package main

var commandDependenciesMap = make(map[string][]string)

// IsIndexed returns true if the package name is in the map, and false if it's not
func IsIndexed(packageName string) bool {
	return commandDependenciesMap[packageName] != nil
}

// IsDependedOn checks whether a package name is current dependend on and returns a corresponding boolean
func IsDependedOn(packageName string) bool {
	for _, dependencies := range commandDependenciesMap {
		for i := range dependencies {
			if dependencies[i] == packageName {
				return true
			}
		}
	}
	return false
}

// IndexPackage indexes a package and it's dependencies in the data store
func IndexPackage(packageName string, dependencies []string) {
	commandDependenciesMap[packageName] = dependencies
}

// RemovePackage removes a package and it's dependencies from the data store
func RemovePackage(packageName string) {
	delete(commandDependenciesMap, packageName)
}

// WipeDataStore is a method exposed purely for testing
func WipeDataStore() {
	for packageName := range commandDependenciesMap {
		delete(commandDependenciesMap, packageName)
	}
}

// GetDataStore is a method exposed purely for testing
func GetDataStore() map[string][]string {
	return commandDependenciesMap
}
