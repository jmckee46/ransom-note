package main

import "testing"

func TestRansomNote1(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}

	checkMagazine(magazine, note)
}

func TestRansomNote2(t *testing.T) {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}

	checkMagazine(magazine, note)
}

func TestRansomNote3(t *testing.T) {
	magazine := []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}
	note := []string{"ive", "got", "some", "coconuts"}

	checkMagazine(magazine, note)
}
