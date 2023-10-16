//
//  MenuItemView.swift
//  menu
//
//  Created by Manuel Antony on 01/06/23.
//

import SwiftUI

struct MenuItemView: View {
    var name: String
    var imageName: String
    var price: String
      
    var body: some View {
        HStack{
            Image(imageName)
                .resizable()
                .aspectRatio(contentMode: .fit)
                .frame(height: 50)
                .cornerRadius(10)
            Text(name)
                .bold()
            
            Spacer()
            
            Text(price)
        }
        .listRowSeparator(.hidden)
        .listRowBackground(
            Color(.brown)
                .opacity(0.1)
        )
    }
}

struct MenuItemView_Previews: PreviewProvider {
    static var previews: some View {
        MenuItemView(name: "Preview Name", imageName: "avocado-maki", price: "$2.20")
    }
}
