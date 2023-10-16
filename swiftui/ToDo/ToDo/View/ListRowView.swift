//
//  ListRowView.swift
//  ToDo
//
//  Created by Manuel Antony on 07/06/23.
//

import SwiftUI

struct ListRowView: View {
    
    let title: String
    let isComplemted: Bool
    
    var body: some View {
        HStack {
            Image(systemName: isComplemted ? "checkmark.circle" : "circle")
                .foregroundColor(isComplemted ? .green : .red)
            Text(title)
            Spacer()
        }
        .font(.title2)
        .padding(.vertical, 5)
    }
}

struct ListRowView_Previews: PreviewProvider {
    
    static var item1 = ItemModel(title: "Item 1", isCompleted: true)
    static var item2 = ItemModel(title: "Item 2", isCompleted: false)
    
    static var previews: some View {
        Group{
            ListRowView(title: item1.title, isComplemted: item1.isCompleted)
            ListRowView(title: item2.title, isComplemted: item2.isCompleted)
        }
        .previewLayout(.sizeThatFits)
    }
}
