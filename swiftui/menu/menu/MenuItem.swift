//
//  MenuItem.swift
//  menu
//
//  Created by Manuel Antony on 31/05/23.
//

import Foundation

struct MenuItem: Identifiable{
    var id: UUID = UUID() // coing from foundation
    var name: String
    var price: String
    var imageName: String
}
