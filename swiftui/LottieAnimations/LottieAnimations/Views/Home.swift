//
//  Home.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 29/06/23.
//

import SwiftUI

struct Home: View {
    
    // offset value will be tracked so state
    @State var onBoardingContent:[OnBoadingContentModel] = [
        OnBoadingContentModel(title: "Welcome", lottieFile: "waterdrop", color: Color.red),
        OnBoadingContentModel(title: "Stay Hydrated", lottieFile: "liquid-blobby-loader", color: Color.yellow),
        OnBoadingContentModel(title: "Start Your Journey", lottieFile: "liquid-water-drop", color: Color.green)
    ]
    
    @GestureState var isDragging: Bool = false
    @State var currrentIndex: Int = 0
    @State var fakeIndex: Int = 0
    
    var body: some View {
        ZStack{
            
            ForEach(onBoardingContent.indices.reversed(), id: \.self){ index in
                OnBoardingView(info: onBoardingContent[index])
                    .clipShape(LiquidLeftSwip(offset: onBoardingContent[index].offSet))
                    .ignoresSafeArea()
            }
        }
        .overlay(
            Image(systemName: "chevron.left")
                .resizable()
                .aspectRatio(contentMode: .fit)
                .frame(width: 10)
                .foregroundColor(.black)
                .offset(x: -7, y:583)
                .gesture(
                    DragGesture()
                        .updating($isDragging, body: { value, out, _ in
                            out = true
                        })
                        .onChanged({ value in
                            withAnimation(
                                .interactiveSpring(response: 0.5,dampingFraction: 0.5, blendDuration: 0.4)){                            onBoardingContent[currrentIndex].offSet = value.translation
                                    
                                }
                            
                        }).onEnded({ value in
                            withAnimation(.spring()){
                                onBoardingContent[currrentIndex].offSet = .zero
                            }
                        })
                )
                .opacity(isDragging ? 0 : 1)
            ,
            alignment: .topTrailing
            
            
        )
        
    }
}

struct Home_Previews: PreviewProvider {
    static var previews: some View {
        Home()
    }
}

