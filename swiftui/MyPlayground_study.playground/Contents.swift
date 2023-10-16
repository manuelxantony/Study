
struct something{
    var s : String {
        return "some"
    }
}

class AClass{
    var name: String

    // should have an initilaizer for properties
    init(Name: String){
        name = Name
    }
}

var c = AClass(Name: "Test")
print(c.name)

struct AStruct {
    var name: String
}

var s = AStruct(name: "Test Struct")
print(s.name)

