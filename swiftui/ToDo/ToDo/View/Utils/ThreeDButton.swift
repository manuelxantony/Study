//
//  ThreeDButton.swift
//  ToDo
//
//  Created by Manuel Antony on 07/06/23.
//

import SwiftUI

struct ThreeDButton: ButtonStyle {
   
    let offset: CGFloat = 10
    
    func makeBody(configuration: Configuration) -> some View {
        
        ZStack{
            IsmometricRectangle(adjustment: 10)
                .fill(.black)
                .offset(y: offset)
            
            IsmometricRectangle(adjustment: 10)
                .fill(.yellow)
                .offset(y: configuration.isPressed ? offset : 0)
            
            configuration.label
                .offset(y: configuration.isPressed ? offset : 0)
                .font(.system(size: 19, design: .rounded))
        }
    }
}

struct ThreeDButton_Previews: PreviewProvider {
    static var previews: some View {
        Button {
            print("Add button pressed")
        } label: {
            Text("save")
                .foregroundColor(.black)
                .font(.system(size: 19, weight: .semibold, design: .rounded))

            
        }
        .frame( width: 150, height: 45)
        .buttonStyle(ThreeDButton())
    }
}
