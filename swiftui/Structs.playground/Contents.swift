struct Car {
    var name: String = "" // if initialized we can ignore variales when intializing the instnace of it
    var cc: Int = 0
    var color: String = ""
    
    func Details() -> String {
        return("Name: \(name), CC: \(cc), Color: \(color)")
    }
}

var c = Car(name: "Porche",cc: 1999, color: "Viper Green")
print(c.name)

var c2 = Car()
c2.name = "bmw"
c2.cc = 1999
print(c2.name)

print(c.Details())
print(c2.Details())


// init methods in Struct
// init is used to intialize the objects
// they are like constructors
struct NCar {
    var name: String
    var cc: Int
    var color: String
    
    init(Name: String, CC: Int, Color: String){
        self.name = Name
        self.cc = CC
        self.color = Color
    }
}

var c3 = NCar(Name: "Mercedez", CC: 1999, Color: "Black")
print(c3.name)

// Struct vs class
// Structs are pass by value type
struct Team{
    var name: String
}

var t1 = Team(name: "Frrari")
var t2 = t1
t2.name = "Red Bull"

print(t1.name, "and", t2.name)

// classes
// call by reference types
class Person {
    var name: String // called as properties
    
    init(Name: String){
        self.name = Name
    }
}

var p1 = Person(Name: "bmw")
var p2 = p1
p2.name = "skoda"

print(p1.name) // will be skoda, because of pass by reference

// also can be used without init
// but not sure it is recommented
class Person2 {
    var name: String = ""
}
var p3 = Person2()
p3.name = "person"
print(p3.name)

var p4 = p3
p4.name = "person2"
print(p3.name)


// lazy properties
// only executed when it called i.e. not when class is initialized
class OS {
    var name: String
    // lazy property
    lazy var nameInfo = {
        print("oiiii")
        return "hey it is \(self.name)" // should have a return
    }() // we have to put () it as it is a function
    init(Name: String){
        self.name = Name
    }
}

var os = OS(Name: "IOS")
print(os.name)
print(os.nameInfo)


// Computed properties
class Rectangle{
    var len: Int
    var wid: Int
    var doubleArea: Int?
    var Area: Int { // Area is acomputed property
        get{
            return len * wid
        }
        set(area){
            doubleArea = 2 * area
        }
    }
    init(Len: Int, Wid: Int){
        self.len = Len
        self.wid = Wid
    }
}

var rec = Rectangle(Len:2, Wid: 3)
print(rec.Area)
rec.Area = 10
print(rec.doubleArea!)

// read only
class Multiplication{
    var num1: Int
    var num2: Int
    var multiply: Int { // read only as there is no setter in computed property
        get{
            return num1 * num2
        }
    }
    
    init(Num1: Int, Num2: Int){
        self.num1 = Num1
        self.num2 = Num2
    }
}

var mul = Multiplication(Num1: 3, Num2: 4)
print(mul.multiply) // we cannot set multiply as there is no setter in computed property


// Observer property
class Number{
    var num = 5{
        willSet(newNum){
            print("Number set is \(newNum)")
        }
        didSet{
            if num > oldValue{
                print("Setted a valid value")
            } else{
                print("Setted a wrong value")
            }
        }
    }
}

var num = Number()
num.num = 10
num.num = 1

// mutating not needed in class
struct Doubler {
    var n: Int
    init(N: Int){
        self.n = N
    }
    
    func double() -> Int {
        return n + n // this will work fine as we are not modifing the class property i.e. n
    }
    // the following code will not work
    /**
    func doubleMut() -> Int{
        n += n
        return n
    }
     */
    mutating func doubleMut() -> Int{ // mutating
        n += n
        return n
    }
}

var d = Doubler(N: 4)
print(d.double())
print(d.doubleMut())
