//
//  ContentView.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 28/06/23.
//

import SwiftUI

struct ContentView: View {
    
    // animation value for controlling the liquid shape
    @State var offset: CGSize = .zero
    @State var index: Int = 0
    
    let animationNames = [
        "liquid-blobby-loader",
        "waterdrop",
        "liquid-water-drop",
        
    ]
    
    var body: some View {
        ZStack{
            // lottie animation view
            //LottieView(lottieFile: "waterdrop")
            
            //views[0]
            // all views are aligned in ZStack
            ForEach(0..<animationNames.count, id: \.self) { index in
                LottieView(lottieFile: animationNames[index])
                    .background(.black)
                    .clipShape(LiquidLeftSwip(offset: offset))
                    .ignoresSafeArea()
            }
            
             
                
        }
        .background(.black.opacity(0.9))
        .overlay(
            Image(systemName: "chevron.left")
                .foregroundColor(.white)
                .offset(x: -7 + offset.width, y:583 + offset.height)
                .opacity(offset == .zero ? 1 : 0)
            ,
            alignment: .topTrailing
        )
        .gesture(DragGesture().onChanged({ value in
            //print(value.translation)
            withAnimation(.interactiveSpring(response: 0.5,
                                             dampingFraction: 0.5, blendDuration: 0.4)){
                offset = value.translation
            }
            
        }).onEnded({ value in
            offset = .zero
        })
        )
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
