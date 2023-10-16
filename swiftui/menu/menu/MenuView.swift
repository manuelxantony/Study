//
//  ContentView.swift
//  menu
//
//  Created by Manuel Antony on 31/05/23.
//

import SwiftUI

struct MenuView: View {
    
    @State var menuItems: [MenuItem] = [MenuItem]()
    var dataService =  DataService()
    
    /**
     What is this error, for the following commented code?!
     "Cannot use instance member 'dataService' within property initializer; property initializers run before 'self' is available"
     */
    /**
     Reason:
     You cannot initialize properties which depend on each other on the top level of the class. Or use a computed property. But can used inside a method.
     i.e:
     dataService is a property in MenuView and getData() is another property
     for assigining property from another instance directly to a property in this
     Instance use init()
     eg:
     init(){
     print("inside init")
     self.menuItems = dataService.getData()
     }
     but this init is executing and printing -> inside init
     but menuItem is not updated
     Source:
     https://stackoverflow.com/questions/43550813/property-initializers-run-before-self-is-available
     */
    //menuItems = dataService.getData()
    
    var body: some View {

            List(menuItems) { item in
                // custom view
                MenuItemView(
                    name: item.name,
                    imageName: item.imageName,
                    price: item.price
                )
            }
           // .listStyle(.plain)
            .onAppear {
                menuItems = dataService.getData()
            }
    }
}

struct MenuView_Previews: PreviewProvider {
    static var previews: some View {
        MenuView()
    }
}
