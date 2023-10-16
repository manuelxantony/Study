package main

import (
	"math"
	"sort"
)

type NodeG struct {
	Point    []int
	Distance float64
}

type HeapNode []NodeG

func kClosest(points [][]int, k int) [][]int {

	nodeSlice := make(HeapNode, len(points))

	for i := 0; i < len(points); i++ {
		point := points[i]
		//distance := point[0]*point[0] + point[1]*point[1]
		distance := math.Pow(float64(point[0]), 2) + math.Pow(float64(point[1]), 2)
		node := NodeG{
			Point:    point,
			Distance: distance,
		}
		nodeSlice[i] = node
	}

	sort.Slice(nodeSlice, func(i, j int) bool {
		return nodeSlice[i].Distance < nodeSlice[j].Distance
	})

	out := make([][]int, k)
	for i := 0; i < k; i++ {
		out[i] = nodeSlice[i].Point
	}
	return out
}
