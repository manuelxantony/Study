//
//  ContentView.swift
//  WeSplit
//
//  Created by Manuel Antony on 01/06/23.
//

import SwiftUI

struct ContentView: View {
    
    @State private var checkAmount = 0.0
    @State private var numberOfPeople = 0
    @State private var tipPercentage = 20
    
    let tipPercentages = [10, 15, 25, 50, 75, 0]
    
    var body: some View {
        NavigationView{
            Form{
                Section {
                    TextField("Amount", value: $checkAmount, format:
                            .currency(code: Locale.current.currency?.identifier ?? "USA"))
                    .keyboardType(.decimalPad)
                    
                    Picker("Number of people", selection: $numberOfPeople){
                        ForEach(2..<30){
                            Text("\($0)")
                        }
                    }
                }
                
                Section {
                    Picker("Tip Percentage", selection: $tipPercentage){
                        ForEach(tipPercentages, id: \.self) { tip in                  Text(tip, format: .percent)
                        }
                    }
                    .pickerStyle(.segmented)
                    
                } header: {
                    Text("How much percentage you want to leave?")
                }
                
                
                Section{
                    Text(checkAmount, format:
                            .currency(code: Locale.current.currency?.identifier ?? "USA"))
                }
                
            }
        }
        .navigationTitle("WeSplit")
        
        
        
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
