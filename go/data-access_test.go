package main

import "testing"

func TestIndexPackage(t *testing.T) {
	WipeDataStore()
	IndexPackage("ls", []string{"1", "2", "3"})
	dataStore := GetDataStore()
	dependencies := dataStore["ls"]
	if len(dependencies) != 3 {
		t.Error("Expected 3 dependencies")
	}
	if dependencies[0] != "1" || dependencies[1] != "2" || dependencies[2] != "3" {
		t.Error("Dependencies do not match expectations")
	}

	IndexPackage("ls", []string{"1"})
	IndexPackage("pwd", []string{"cat", "ls"})

	dependencies = dataStore["ls"]
	if len(dependencies) != 1 {
		t.Error("Expected 1 dependency in updated dependency list for ls")
	}
	if dependencies[0] != "1" {
		t.Error("Expected ls to have a dependency on 1")
	}

	dependencies = dataStore["pwd"]
	if len(dependencies) != 2 {
		t.Error("Expected 2 deps for pwd")
	}
	if dependencies[0] != "cat" || dependencies[1] != "ls" {
		t.Error("pwd Dependencies do not match expectations")
	}
}

func TestIsIndexed(t *testing.T) {
	IndexPackage("ls", []string{"1", "2", "3"})
	IndexPackage("cat", []string{"1", "2", "3"})
	IndexPackage("dog", []string{"1", "2", "3"})

	result := IsIndexed("taco")
	if result {
		t.Error("Taco should not be indexed")
	}

	resultTwo := IsIndexed("cat")
	resultThree := IsIndexed("ls")
	resultFour := IsIndexed("dog")
	if resultTwo == false || resultThree == false || resultFour == false {
		t.Error("cat, ls, and dog should all be indexed")
	}
}

func TestRemovePackage(t *testing.T) {
	IndexPackage("ls", []string{"1", "2", "3"})
	IndexPackage("cat", []string{"1", "2", "3"})
	IndexPackage("dog", []string{"1", "2", "3"})

	dataStore := GetDataStore()
	lsDependencies := dataStore["ls"]
	catDependencies := dataStore["cat"]

	if lsDependencies == nil || catDependencies == nil {
		t.Error("ls or cat dependencies were null too early in the test")
	}

	RemovePackage("ls")
	RemovePackage("cat")

	lsDependencies = dataStore["ls"]
	catDependencies = dataStore["cat"]

	if lsDependencies != nil {
		t.Error("ls dependencies should be null")
	}

	if catDependencies != nil {
		t.Error("cat dependencies should be null")
	}

}
