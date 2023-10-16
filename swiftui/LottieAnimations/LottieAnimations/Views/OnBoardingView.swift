//
//  OnBoardingView.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 29/06/23.
//

import SwiftUI

struct OnBoardingView: View {
    
    var info: OnBoadingContentModel
    
    var body: some View {
        VStack{
            //ZStack{
           

                VStack{
                    Spacer()
                    Text(info.title)
                        .fontWeight(.heavy)
                        .font(.largeTitle)
                        .foregroundColor(.white)
                    
                    LottieView(lottieFile: info.lottieFile)
                }
                .padding(.bottom, 500)
                
           // }
                .background(info.color)
        }
    }
}

struct OnBoardingView_Previews: PreviewProvider {
    static var info: OnBoadingContentModel = OnBoadingContentModel(title: "Welcome", lottieFile: "waterdrop", color: Color.red)
    static var previews: some View {
        OnBoardingView(info: info)
            .background(info.color)
    }
}
