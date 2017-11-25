package main

import (
  "fmt"
  "math"
  )

const dataSize = 4
var generes = []string{"Romantic","Action"}

type Point struct{
  x int
  y int
  label string
  distance float64
}

type KNN struct{
  k int
  data []*Point
}

func newKnn(k int, data []*Point) *KNN{
  knn := new(KNN)
  knn.k = k
  knn.data = data
  return knn
}

func (knn *KNN) classfy(p *Point) string{
  distances := knn.euclideanDistance(p)
  label := knn.majorityClass(distances)
  return label
}

func (knn *KNN) euclideanDistance(p *Point) [dataSize]float64{
  var sqMat [dataSize]float64
  for i:=0;i < dataSize;i++{
    //difference of the two Xs ( of p and given point )
    //then square root it and add it to the
    //difference of the two Ys ( of p and given point )
    //then take the square root of it
    sqMat[i] = math.Sqrt(math.Pow(float64(p.x - knn.data[i].x), 2) + math.Pow(float64(p.y - knn.data[i].y), 2))
    knn.data[i].distance = sqMat[i]
  }
  fmt.Println(sqMat)
  return sqMat
}

func (knn *KNN) majorityClass(dists [dataSize]float64) string{
	for i := 0; i < dataSize; i++ {
		iMin := i
		for j := i + 1; j < dataSize-1; j++ {
			if knn.data[j].distance > knn.data[iMin].distance {
				iMin = j
			}
		}
    if iMin != i{
      temp := knn.data[i]
  		knn.data[i] = knn.data[iMin]
  		knn.data[iMin] = temp
    }
	}
  labels := []int{0,0}
  for j := 0; j < knn.k; j++{
    d := knn.k-j
    if knn.data[d].label == generes[0]{
      labels[0]++
    }else{
      labels[1]++
    }
  }
  if labels[0] > labels[1]{
    return generes[0]
  }else{
    return generes[1]
  }
}

func main(){
  p1 := &Point{
    x : 1,
    y : 2,
    label : generes[1],
  }
  p2 := &Point{
    x : 1,
    y : 4,
    label : generes[1],
  }
  p3 := &Point{
    x : 4,
    y : 1,
    label : generes[0],
  }
  p4 := &Point{
    x : 4,
    y : 2,
    label : generes[0],
  }
  pp := &Point{
    x : 4,
    y : 3,
  }
  data := []*Point{p1, p2, p3, p4}
  knn := newKnn(3, data)
  fmt.Println(knn.classfy(pp))
}
