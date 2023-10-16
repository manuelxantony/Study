//
//  ToDoApp.swift
//  ToDo
//
//  Created by Manuel Antony on 07/06/23.
//

import SwiftUI

@main
struct ToDoApp: App {
    
    //@ObservedObject var listItemViewModel: ListItemViewModel = ListItemViewModel()
    @StateObject var listItemViewModel: ListItemViewModel = ListItemViewModel()
    
    var body: some Scene {
        WindowGroup {
            NavigationView {
                ListView()
            }
          .environmentObject(listItemViewModel)
        }
    }
}
