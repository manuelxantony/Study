//
//  ListItemViewModel.swift
//  ToDo
//
//  Created by Manuel Antony on 09/06/23.
//

import Foundation

class ListItemViewModel: ObservableObject {
 
    @Published var listItems: [ItemModel] = [] {
        didSet{
            saveItems()
        }
    }
    
    let ITEM_SAVE_KEY = "item_save_key"
    
    init() {
        getItems()
    }
    
    func getItems(){
        guard
            let data = UserDefaults.standard.data(forKey: ITEM_SAVE_KEY),
            let savedItems = try? JSONDecoder().decode([ItemModel].self, from: data)
        else { return }
        
        self.listItems = savedItems
    }
    
    func deleteItem(indexSet: IndexSet){
        listItems.remove(atOffsets: indexSet)
    }
    
    func moveItem(from: IndexSet, to: Int) {
        listItems.move(fromOffsets: from, toOffset: to)
    }
    
    func addItem(item: String){
        guard item.count > 0 else { return }
        
        let newListItem: ItemModel = ItemModel(title: item, isCompleted: false)
        listItems.append(newListItem)
    }
    
    func updateItem(itemToUpdate: ItemModel) {
        guard let index = listItems.firstIndex(where: { $0.id == itemToUpdate.id }) else { return }
        listItems[index] = itemToUpdate.updateItemIsCompleted()
    }
    
    func saveItems() {
        guard let encodedData = try? JSONEncoder().encode(listItems) else { return }
        UserDefaults.standard.set(encodedData, forKey: ITEM_SAVE_KEY)
    }
}
