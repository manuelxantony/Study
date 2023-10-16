// Inheritance
class Heros {
    var name: String
    var age: Int
    
    init(Name: String, Age: Int) {
        self.name = Name
        self.age = Age
    }
}

class DC: Heros{}
class Marvel: Heros{}

// can use DC and Marvel in following ways
var dc = DC(Name: "Batman", Age: 38)
var marvel = Marvel.init(Name: "Spiderman", Age: 18)

print(dc.name)
print(marvel.name)

// overriding
class Car {
    func printCar(name: String) {
        print("Car name", name)
    }
}

class SuperCar: Car {
    override func printCar(name: String) {
        //self.printCar(name: name)
        print("Super car name", name)
    }
}

var sc = SuperCar()
sc.printCar(name: "Porche")

class Movie{
    var name: String?
    var year: Int{
        return 0000
    }
    func getGenere() -> String {
        return "movie"
    }
    
}

class IndianMovie: Movie{
    // override var name: String = "2018" // cannot override
    var movieName: String = "2018"
    override var year: Int{
        return 2023
    }
    override func getGenere() -> String {
        return "Thriller, Drama"
    }
}

var im = IndianMovie()
print(im.movieName, im.year, im.getGenere())

