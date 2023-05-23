package main

import "testing"

func BenchmarkAVLTree2(b *testing.B) {
	avl := NewAVLTree()
	for i := 0; i < b.N; i++ {
		avl.Insert(i)
	}
}

func BenchmarkRBTree2(b *testing.B) {
	rb := RBTree{}
	for i := 0; i < b.N; i++ {
		rb.Insert(i)
	}
}

func BenchmarkAVLTreeSearch(b *testing.B) {
	avl := NewAVLTree()
	for i := 0; i < b.N; i++ {
		avl.Insert(i)
	}
	for i := 0; i < 10*b.N; i++ {
		avl.Search(i)
	}
}

func BenchmarkRBTreeSearch(b *testing.B) {
	rb := RBTree{}
	for i := 0; i < b.N; i++ {
		rb.Insert(i)
	}
	for i := 0; i < 10*b.N; i++ {
		rb.Search(i)
	}
}

// benchmarker with fixed number of inserts and searches
func BenchmarkAVLTreeSearchFixed(b *testing.B) {
	avl := NewAVLTree()
	for i := 0; i < 1000000; i++ {
		avl.Insert(i)
	}
	for i := 0; i < 5000000; i++ {
		avl.Search(i)
	}
}

func BenchmarkRBTreeSearchFixed(b *testing.B) {
	rb := RBTree{}
	for i := 0; i < 1000000; i++ {
		rb.Insert(i)
	}
	for i := 0; i < 5000000; i++ {
		rb.Search(i)
	}
}

func BenchmarkRBTreeOnlyInserts(b *testing.B) {
	rb := RBTree{}
	for i := 0; i < 1000000; i++ {
		rb.Insert(i)
	}
}

func BenchmarkAVLTreeOnlyInserts(b *testing.B) {
	avl := NewAVLTree()
	for i := 0; i < 1000000; i++ {
		avl.Insert(i)
	}
}
