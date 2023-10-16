//
//  CoreDataProvider.swift
//  RemindersApp
//
//  Created by Manuel Antony on 24/07/23.
//

import Foundation
import CoreData



class CoreDataProvider {
    
    // singleton
    static let shared = CoreDataProvider()
    
    let persistentContainer: NSPersistentContainer
    
    private init() {
        persistentContainer = NSPersistentContainer(name: "RemindersModel")
        persistentContainer.loadPersistentStores { description, error in
            if let error {
                fatalError("Error handling RemindersModel \(error)")
            }
        }
    }
}
