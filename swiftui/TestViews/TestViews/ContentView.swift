//
//  ContentView.swift
//  TestViews
//
//  Created by Manuel Antony on 15/06/23.
//

import SwiftUI

struct ContentView: View {
    let data = (1..<10).map {"Itembbbbnnnkjkjnmnmnmn \($0)"}
    
    let col = Array(repeating: GridItem (.adaptive(minimum: 100 ), spacing: 20), count: 2)
    
    var body: some View {
        ScrollView {
            LazyVGrid(columns: col, spacing: 30) {
                ForEach(data, id: \.self) { name in
                    ZStack(alignment: Alignment(horizontal: .trailing, vertical: .top)){
                        VStack  (alignment: .leading, spacing: 20){
                            Image(systemName: "figure.run")
                            
                            Text(name)
                        }
                        .padding()
                        .background(.green)
                        .cornerRadius(10)
                    }
                                        }
            }
            //.padding()
            .padding()
        }
        
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
        
    }
}
