package main

import "testing"

// Assume you have AVLTree and RBTree structs with Insert method

func BenchmarkAVLTree(b *testing.B) {
	avl := AVLTree{}
	for i := 0; i < b.N; i++ {
		avl.Insert(i)
	}
}

func BenchmarkRBTree(b *testing.B) {
	rb := RBTree{}
	for i := 0; i < b.N; i++ {
		rb.Insert(i)
	}
}
