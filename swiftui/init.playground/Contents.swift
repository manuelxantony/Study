class Concatenate{
    var string1: String
    var string2: String
    var concatenatedString: String
    
    init(String1: String, String2: String){
        string1 = String1
        string2 = String2
        concatenatedString = string1 + ", " + string2
    }
}

var c = Concatenate(String1: "Hello", String2: "World")
print(c.concatenatedString)
    
class Concatenate2{
    var string1 = "hello"
    var string2 = "world"
    
    var concatenatedString: String
    
    init(){
        concatenatedString = string1 + ", " + string2
    }
}

var c2 = Concatenate2()
print(c2.concatenatedString)
    
// failable init
class User{
    var uuid: String
    init?(UUID:String){ // conditional init
        if UUID.count == 5 {
            self.uuid = UUID
        } else {
            return nil
        }
    }
}

var u = User(UUID: "abcd")
print(u)

var u2 = User(UUID: "abcde")
print(u2?.uuid ?? <#default value#> )


// Deinit -> only in classes
// used for cleaning init values
class Hero{
    var name: String
    
    init(_ heroName: String){
        name = heroName
        print("I'm \(name)")
    }
    // ARC
    deinit{
        print("\(name) is gone")
    }
    
}


for _ in 1...5{
   Hero("Batman")
}
