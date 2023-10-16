// variable
var name: String // declaration
name = "blue" // intialization
var newName = "blue314" // declaration and intialization
var anotherName: String = "bluex" // declaration and intialization
print(name)

// constant
let imAConstant: Int = 20 // let is immutable/constant
print(imAConstant)

// closed range -> ClosedRange<>
var range = 1...5 // from 1 to 5
print(range.contains(5))

var anotherRange: ClosedRange<Int> // see diffrence between Range and ClosedRange
anotherRange = 1...10
print(18, anotherRange.contains(9))

// half open range -> Range<>
var yetAnotherRange = 1..<5 // from 1 to less than 5
print(yetAnotherRange.contains(5))

var aRange: Range<Int> = 4..<30
print(aRange.contains(29))

// String interpolation
var codeName = "blue"
var age = 31
var data = "hi I'm \(codeName) and \(age) years here on earth"
print(31, data)

var data2 = "hi I'm " + codeName + " and " + String(age) + " years here on earth"
print(34, data2)

// optionals
var userName:String? // optional
print(38, userName) // out -> nil
// print(38, userName!) // out -> error

userName = "blue"
print(41, userName) // out -> Optional("blue")

// forced unwrapping (use ! at the end to tell compiler that it is optional)
print(44, userName!) // out -> blue



// optinal unwrapping
var carName: String?

if let getCarName = carName{ // optinal unwrapping
    print(53, getCarName)
}
carName = "Proche"
print(56, carName!) // forced

if let getCarNameAgain = carName{
    print(59, getCarNameAgain)
}


// Arrays
var cars:[String] = ["Porche", "Ferrari", "BMW"]
// the array alos can be used as
// var cars:Array<String> = Array<String>("Porche", "Ferrari", "BMW")
print(67, cars)
cars.append("Mercedez")
print(cars)
cars.insert("Audi", at:3)
print(cars)

// sort and reverse make a copy of array and do the operation
// sorted and reversed do operation on the same array
// no other diffrence in functional wise

//Dictionaries
var dict: [String: Int] = [
    "a": 124,
    "b c": 2432,
    "c d": 35343
    ]

print(dict)
print(dict["a"]!) // see how ! is used
var val = dict["a"]
print(val!) // here also !

dict.keys // {{Builtin.BridgeObject}}
print(dict.keys) // print will print the actual keys

// but if we want to use the keys it will give {{Builtin.BridgeObject}} type
// to extract the keys out of the {{Builtin.BridgeObject}} we have type cast
dict.keys // out -> {{Builtin.BridgeObject}}
var allKeys = [String](dict.keys)
print(allKeys)

// updating a value
var oldValue = dict.updateValue(929, forKey: "a")
print(oldValue!) // 124
print(dict["a"]!) // 929

// removing key
dict["c d"] = nil
print(dict)

dict.removeValue(forKey: "a")
print(dict)

// Nested Dictonaries
var nestedDict = [
    "user 1" : [
        "name" : "blue",
        "age": 29
        ],
    "user 2" : [
        "name" : "blueX",
        "age": 31
        ],
    ]

print(nestedDict)
print(123, nestedDict["user 1"]!) // forced unwrapping
print((nestedDict["user 1"]?["name"])!) // here ? means if user 1 is not nil proced to find name

// optinal unwrapping in dict
if let user1 = nestedDict["user 1"] {
    print(128, user1)
}

// optional unwrapping in dict is used to make sure that progrma wont crash if we give a wrong key
// example

// optinal unwrapping in dict
if let user2 = nestedDict["user two"] {
    print(136, user2) // will not print this line as "user two" is a wrong key
}
// also we can use unwrap multiple values in optinal unwrapping
if let userOne = nestedDict["user 1"], let userTwo = nestedDict["user 2"] {
    print(140, userOne, userTwo)
}

if let userOne = nestedDict["user One"], let userTwo = nestedDict["user 2"] {
    print(144, userOne, userTwo) // will not print anything as "user One" is a wrong key
}

// Set
var gameV1: Set = [ // also can be Set<Strings>
    "start game",
    "get stones",
    "throw stone",
    "get stones", // will be ignored - no duplicate items in a Set
    "get the flag"
]

var gameV2: Set = [
    "start game",
    "get guns",
    "fire guns",
    "get the flag"
]

var commonInV1V2 = gameV1.intersection(gameV2) // common in V1 and V2
print(commonInV1V2)

print(gameV1.symmetricDifference(gameV2)) // diffrences

// Tuples
var user: (name: String, age: Int, active:Bool) = ("blue", 31, true)
print(user.name)

// control flow
// uses for in loop
// can range using 1...5 or 1..<6
// normal if else and while loops

// function
func hello(){
    print("hello")
}
hello()

func helloWho() -> String {
    return "blueX314"
}
print(helloWho())

func helloAll() -> (String, String) {
    return ("blue", "blue314")
}
print(helloAll())

var (n1, n2) = helloAll()
print(194, n2)

var all = helloAll()
print(197, all.0)

func helloName(name: String) {
    print("hello, \(name)")
}

helloName(name: "blue")


// Closure (similar to lamda in Python, or annonumous function in Java)
var movie = {
    print("Welcome to movie")
}

movie()

// declaration of closure
var aClosure: () -> () = {}
                // argument //return type
var numClosure: (Int) -> (Int) = { (num: Int) -> Int in
    return num + 2
}
print(numClosure(2))

// here compiler will infer the types
var anotherNumberClosure = { aNum in
    return aNum + 5
}
print(anotherNumberClosure(5))
var anotherNumberClosure2 = { aNum -> Int in
    return aNum + 5
}

print(anotherNumberClosure2(10))

// closure with only return type
var giveCurrentYear = { () -> Int in
    return 2023
}
print(giveCurrentYear())

// a closure with only parameter
var justPrint = { (printext: String) in
    print(240, printext)

}
justPrint("hello, Closure")

// a closure with only parameter without type definition for parameter
var justPrint2 = { printext in
    print(247, printext)

}
justPrint2("hello, Closure2"

// using clouse in practical example
// clousure function
func calculate(num: Int) -> Int { num * 5 }

// a function which accepts closure
func sum(num1: Int, num2: Int, closureParam: (Int) -> (Int)) -> Int{
    var sum = 0
    for num in num1...num2 {
        sum += closureParam(num)
    }
    return sum
}
print(sum(num1: 1, num2: 3, closureParam: calculate)) // here calculate we just give the reference only

// escaping -> will see
// autoclouse -> will see


// Enum
// will see in practical exapmles

