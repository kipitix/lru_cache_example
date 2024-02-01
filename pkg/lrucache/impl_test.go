package lrucache

import (
	"errors"
	"testing"
)

var (
	_errorOnCreate     = errors.New("Cache not created")
	_errorOnAdd        = errors.New("Can`t add element")
	_errorLength       = errors.New("Wrong length")
	_errorCapacity     = errors.New("Wrong capacity")
	_errorMustNotExist = errors.New("Item mustn`t exist")
	_errorMustExist    = errors.New("Item must exist")
)

func TestCreate(t *testing.T) {
	c := NewLRUCache[string, string](10)

	if c == nil {
		t.Error(_errorOnCreate)
	}

	if c.Length() != 0 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}
}

func TestAddSimple(t *testing.T) {
	c := NewLRUCache[string, string](10)

	if c == nil {
		t.Error(_errorOnCreate)
	}

	if !c.Add("x", "1") {
		t.Error(_errorOnAdd)
	}

	if !c.Add("y", "2") {
		t.Error(_errorOnAdd)
	}

	if !c.Add("z", "3") {
		t.Error(_errorOnAdd)
	}

	if c.Length() != 3 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}
}

func TestAddWithDisplacement(t *testing.T) {
	// Create cache with size 3
	c := NewLRUCache[string, int](3)

	if c == nil {
		t.Error(_errorOnCreate)
	}

	// Add three items

	if !c.Add("x", 1) {
		t.Error(_errorOnAdd)
	}

	if !c.Add("y", 2) {
		t.Error(_errorOnAdd)
	}

	if !c.Add("z", 3) {
		t.Error(_errorOnAdd)
	}

	if c.Length() != 3 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 3 {
		t.Error(_errorCapacity)
	}

	// Add one more item

	// Here element "x" will be deleted
	if !c.Add("a", 10) {
		t.Error(_errorOnAdd)
	}

	if c.Length() != 3 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 3 {
		t.Error(_errorCapacity)
	}

	// Check "x"
	val, found := c.Get("x")
	if found || val != 0 {
		t.Error(_errorMustNotExist)
	}

	// Check "y"
	val, found = c.Get("y")
	if !found || val != 2 {
		t.Error(_errorMustExist)
	}

	// Check "z"
	val, found = c.Get("z")
	if !found || val != 3 {
		t.Error(_errorMustExist)
	}

	// Check "a"
	val, found = c.Get("a")
	if !found || val != 10 {
		t.Error(_errorMustExist)
	}
}

func TestGetSimple(t *testing.T) {
	// Create cache with size 3
	c := NewLRUCache[string, int](3)

	if c == nil {
		t.Error(_errorOnCreate)
	}

	if !c.Add("x", 1) {
		t.Error(_errorOnAdd)
	}

	if c.Length() != 1 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 3 {
		t.Error(_errorCapacity)
	}

	// Check "x"
	val, found := c.Get("x")
	if !found || val != 1 {
		t.Error(_errorMustExist)
	}

	// Check "y"
	val, found = c.Get("y")
	if found || val != 0 {
		t.Error(_errorMustNotExist)
	}

	// Check "z"
	val, found = c.Get("z")
	if found || val != 0 {
		t.Error(_errorMustNotExist)
	}
}

func TestDeleteSimple(t *testing.T) {
	// Create cache with size 10
	c := NewLRUCache[string, int](10)

	if c == nil {
		t.Error(_errorOnCreate)
	}

	// Add some elements

	if !c.Add("x", 1) {
		t.Error(_errorOnAdd)
	}

	if !c.Add("y", 2) {
		t.Error(_errorOnAdd)
	}

	if !c.Add("z", 3) {
		t.Error(_errorOnAdd)
	}

	if c.Length() != 3 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}

	// Failure delete

	if c.Remove("a") {
		t.Error(_errorMustNotExist)
	}

	if c.Remove("b") {
		t.Error(_errorMustNotExist)
	}

	if c.Remove("c") {
		t.Error(_errorMustNotExist)
	}

	if c.Length() != 3 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}

	// Success delete

	if !c.Remove("x") {
		t.Error(_errorMustExist)
	}

	if c.Length() != 2 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}

	if !c.Remove("y") {
		t.Error(_errorMustExist)
	}

	if c.Length() != 1 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}

	if !c.Remove("z") {
		t.Error(_errorMustExist)
	}

	if c.Length() != 0 {
		t.Error(_errorLength)
	}

	if c.Capacity() != 10 {
		t.Error(_errorCapacity)
	}

	// Check "x"
	val, found := c.Get("x")
	if found || val != 0 {
		t.Error(_errorMustNotExist)
	}

	// Check "y"
	val, found = c.Get("y")
	if found || val != 0 {
		t.Error(_errorMustNotExist)
	}

	// Check "z"
	val, found = c.Get("z")
	if found || val != 0 {
		t.Error(_errorMustNotExist)
	}
}
