//
//  ListView.swift
//  ToDo
//
//  Created by Manuel Antony on 07/06/23.
//

import SwiftUI

struct ListView: View {
    
    //@ObservedObject var listItemViewModel: ListItemViewModel = ListItemViewModel()
    //@StateObject var listItemViewModel: ListItemViewModel = ListItemViewModel()
    
    @EnvironmentObject var listItemViewModel: ListItemViewModel
    
    var body: some View {
        ZStack{
            if listItemViewModel.listItems.isEmpty {
                NoItemsView()
            } else {
                List {
                    ForEach(listItemViewModel.listItems) { item in
                        ListRowView(title: item.title, isComplemted: item.isCompleted)
                            .onTapGesture {
                                withAnimation(.easeIn) {
                                    listItemViewModel.updateItem(itemToUpdate: item)
                                }
                            }
                    }
                    .onDelete(perform: listItemViewModel.deleteItem)
                    .onMove(perform: listItemViewModel.moveItem)
                }
            }
        }
        .navigationTitle("Todo List 📝")
        .listStyle(PlainListStyle())
        .navigationBarItems(
            leading: EditButton(),
            trailing:
                NavigationLink("Add", destination: AddItemView())
        )
    }
}

struct ListView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationView{
            ListView()
        }
        .environmentObject(ListItemViewModel())
    }
}


