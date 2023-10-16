package main

func canConstruct(ransomNote string, magazine string) bool {
    
    mMap := make(map[rune]rune)

    for _, v := range magazine {
     
		if _, ok := mMap[v]; !ok {
            mMap[v] = 1
        } else {
        	 mMap[v] += 1	
		}
    }

    for _, v := range ransomNote {
        if _, ok := mMap[v]; !ok {
            return false
        } else {
            if mMap[v] <= 0 {
                return false
            }
            mMap[v] -= 1
        }
        
    }
    return true
}