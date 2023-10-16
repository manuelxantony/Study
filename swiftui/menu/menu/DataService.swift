//
//  DataService.swift
//  menu
//
//  Created by Manuel Antony on 01/06/23.
//

import Foundation

struct DataService {
    
    func getData() -> [MenuItem] {
        return [
            MenuItem(name: "avocado-maki", price: "$2.40", imageName: "avocado-maki"),
            MenuItem(name: "california-roll", price: "$1.45", imageName: "california-roll"),
            MenuItem(name: "ikura-sushi", price: "$3.50", imageName: "ikura-sushi"),
            MenuItem(name: "shrimp-sushi", price: "$1.50", imageName: "shrimp-sushi"),
            MenuItem(name: "meguro-sushi", price: "$4.40", imageName: "meguro-sushi"),
            MenuItem(name: "hamachi-sushi", price: "$2.43", imageName: "hamachi-sushi"),
            MenuItem(name: "tamago-sushi", price: "$1.34", imageName: "tamago-sushi"),
            MenuItem(name: "tobiko-spicy-maki", price: "$1.40", imageName: "tobiko-spicy-maki"),
            MenuItem(name: "salmon-sushi", price: "$2.03", imageName: "salmon-sushi"),
            MenuItem(name: "kani-sushi", price: "$1.50", imageName: "kani-sushi"),
            MenuItem(name: "tako-sushi", price: "$1.42", imageName: "tako-sushi"),
            MenuItem(name: "onigiri", price: "$3.32", imageName: "onigiri")
        ]
    }
}
