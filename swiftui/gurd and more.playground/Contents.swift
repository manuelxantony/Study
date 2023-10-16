////let name = "man"
////let age = 20
////
////func checkName() {
////    if name == "" { return }
////    print(name, 1)
////
////    // or
////
////    if name != "" { // checking against what we wanted
////        print(name, 2)
////    }
////}
////
////
////checkName()
////
////func checkNameGurad() {
////    guard name.count > 0 else { return }
////
////    print(name)
////}
////
////checkNameGurad()
////
//
//// can be used for conditonal unwrapping
//var car: String?
//car = "Audi"
//
//if let carName = car  {
//    print(carName)
//}
//
//car = "Proche"
//func getCarName() {
//    guard let carName: String = car else { return }
//    print(carName)
//}
//
//getCarName()
//
//// let
//let aVar: String
//aVar = "okaee"
//print(aVar)
//
//


//struct MyText {
//    var text: String
//}
//
//extension MyText {
//    func makeLower() -> String {
//        return text.lowercased()
//    }
//}
//
//var t = MyText(text: "SOME TEXT")
//print(t.makeLower())


let num = [String](repeating: "rer", count: 10)

print(num)

let num2d = [[String]](repeating: [String](repeating: "2d arr", count: 2), count: 2)

print(num2d)
