package main


type RandomizedSet struct {
    RSet map[int]int
}


// func Constructor() RandomizedSet {
//     return RandomizedSet{
// 		RSet: map[int]int{},
// 	}
// }


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.RSet[val]; !ok {
		this.RSet[val] = val
		return true
	}
	return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.RSet[val]; ok {
		delete(this.RSet, val)
		return true
	}
	return false
}


func (this *RandomizedSet) GetRandom() int {

	for k := range this.RSet {
		return k
	}
	return -1
}
