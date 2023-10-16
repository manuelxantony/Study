//
//  AddItemView.swift
//  ToDo
//
//  Created by Manuel Antony on 07/06/23.
//

import SwiftUI

struct AddItemView: View {
    
    @Environment(\.presentationMode) var presentationMode
    @EnvironmentObject var listItemViewModel: ListItemViewModel
    @State var inputItemText:String = ""
    
    let textFieldBGColor = Color(#colorLiteral(red: 0.921431005, green: 0.9214526415, blue: 0.9214410186, alpha: 0.8256156871))
    let saveButtonColor = Color(UIColor.yellow)
    let saveButtonTextColor = Color(.black)
    
    var body: some View {
        
        ScrollView{
            
            VStack(alignment: .center) {
                
                TextField("Enter your Goal here", text: $inputItemText)
                    .padding(.horizontal)
                    .frame(height: 50)
                    .background(textFieldBGColor)
                Button(
                    action: saveItemPressed,
                    label: {
                    Text("save")
                        .foregroundColor(.black)
                        .font(.system(size: 18, design: .rounded))
                })
                .frame( width: 150, height: 45)
                .buttonStyle(ThreeDButton())
                .padding(.top, 30)
            }
            .padding()
            }
            .navigationTitle("Add Goal")
        }
    
    func saveItemPressed(){
        listItemViewModel.addItem(item: inputItemText)
        presentationMode.wrappedValue.dismiss()
    }
    
    
}

struct AddItemView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationView{
            AddItemView()
        }
        .environmentObject(ListItemViewModel())
    }
}
