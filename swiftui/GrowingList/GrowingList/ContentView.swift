//
//  ContentView.swift
//  GrowingList
//
//  Created by Manuel Antony on 01/06/23.
//

import SwiftUI

struct ContentView: View {
    
    let cars: [String] = DataService().getData()
    
    @State var carList: [Car] = [Car]()
    
    var body: some View {
        VStack(alignment: .center){
            List(carList) { car in
                Text(car.name)
            }.listStyle(.plain)
            Spacer()
            Button {
                addCars()
            } label: {
                Text("Add Cars")
            }
        }
    }
    
    func addCars() {
        let rand = Int.random(in: 0..<cars.count)
        let car: Car = Car(name: cars[rand])
        carList.append(car)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
