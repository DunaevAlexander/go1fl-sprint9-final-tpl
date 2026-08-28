package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "positive size",
			size: 10,
			want: 10,
		},
		{
			name: "zero size",
			size: 0,
			want: 0,
		},
		{
			name: "negative size",
			size: -5,
			want: 0,
		},
		{
			name: "large size",
			size: 1000000,
			want: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)
			
			if len(result) != tt.want {
				t.Errorf("generateRandomElements() length = %v, want %v", len(result), tt.want)
			}
			
			if tt.size > 0 {
				for i, val := range result {
					if val <= 0 {
						t.Errorf("generateRandomElements() element at index %d is not positive: %v", i, val)
					}
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "normal case",
			data: []int{1, 5, 3, 9, 2, 7},
			want: 9,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "all equal",
			data: []int{5, 5, 5, 5},
			want: 5,
		},
		{
			name: "sorted ascending",
			data: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "sorted descending",
			data: []int{10, 9, 8, 7, 6},
			want: 10,
		},
		{
			name: "negative numbers",
			data: []int{-5, -2, -8, -1},
			want: -1,
		},
		{
			name: "mixed numbers",
			data: []int{-5, 0, 3, -1, 10, -3},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			if result != tt.want {
				t.Errorf("maximum() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "normal case",
			data: []int{1, 5, 3, 9, 2, 7, 4, 8, 6},
			want: 9,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "all equal",
			data: []int{5, 5, 5, 5, 5, 5, 5, 5},
			want: 5,
		},
		{
			name: "sorted ascending",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
		{
			name: "sorted descending",
			data: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: 10,
		},
		{
			name: "negative numbers",
			data: []int{-5, -2, -8, -1, -3, -7, -4, -6},
			want: -1,
		},
		{
			name: "less than chunks",
			data: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "exactly chunks count",
			data: []int{10, 20, 30, 40, 50, 60, 70, 80},
			want: 80,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.data)
			if result != tt.want {
				t.Errorf("maxChunks() = %v, want %v", result, tt.want)
			}
		})
	}
}

func BenchmarkMaximum(b *testing.B) {
	data := generateRandomElements(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximum(data)
	}
}

func BenchmarkMaxChunks(b *testing.B) {
	data := generateRandomElements(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxChunks(data)
	}
}