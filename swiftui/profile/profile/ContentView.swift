//
//  ContentView.swift
//  profile
//
//  Created by Manuel Antony on 30/05/23.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        
        // main view
        ZStack{
            Image("background").resizable()
                .edgesIgnoringSafeArea(.all)
            
            // main vertical view
            VStack(alignment: .center){ // why .leading and .trailing is similar to .center
                
                // view for alligning pic title and description
                VStack(alignment: .center, spacing: 15){ // why alignment have no effect
                    Image("profile").resizable()
                        .aspectRatio(contentMode: .fit)
                        .frame(width: 150, height: 150)
                        .clipShape(Circle())
                        .shadow(color: .purple, radius: 0, x: -2, y: 4)
                        .shadow(color: .black, radius: 7, x: -7, y: 10)
                    
                    Text("Purple")
                        .font(.system(size: 40))
                        .fontWeight(.bold)
                        .foregroundColor(.white)
                        .shadow(color: .black, radius: 7, x: -7, y: 10)
                    
                    Text("IOS Engineer")
                        .font(.footnote)
                        .foregroundColor(.white)
                    
                    HStack(alignment: .center, spacing: 30){
                        Image(systemName: "heart").resizable()
                            .aspectRatio(contentMode: .fit)
                            .frame(width: 20, height: 20)
                            .foregroundColor(.white)
                        
                        Image(systemName: "network").resizable()
                            .aspectRatio(contentMode: .fit)
                            .frame(width: 20, height: 20)
                            .foregroundColor(.white)
                        
                        Image(systemName: "message").resizable()
                            .aspectRatio(contentMode: .fit)
                            .frame(width: 20, height: 20)
                            .foregroundColor(.white)
                        
                        Image(systemName: "phone").resizable()
                            .aspectRatio(contentMode: .fit)
                            .frame(width: 20, height: 20)
                            .foregroundColor(.white)
                        
                    }
                }
                Spacer()
                
                // view for button and count
                VStack(alignment: .center, spacing: 20){
                    RoundedRectangle(cornerRadius: 8)
                        .frame(width: 150, height: 50)
                        .foregroundColor(.white)
                        .shadow(color: .purple, radius: 0, x:-2, y:4)
                        .shadow(color: .black, radius: 7, x: -7, y: 10)
                        .overlay(
                            Text("Follow")
                                .font(.system(size: 20))
                                .fontWeight(.bold)
                        )
                    HStack(alignment: .center, spacing: 30){
                        VStack(alignment: .center, spacing: 5){
                            Text("699")
                                .foregroundColor(.white)
                            Text("hearts")
                                .foregroundColor(.gray)
                        }
                        VStack(alignment: .center, spacing: 5){
                            Text("499")
                                .foregroundColor(.white)
                            Text("followers")
                                .foregroundColor(.gray)
                        }
                        VStack(alignment: .center, spacing: 5){
                            Text("109")
                                .foregroundColor(.white)
                            Text("following")
                                .foregroundColor(.gray)
                        }
                    }
                }
                Spacer()
                VStack(alignment: .center,spacing: 20){
                    Text("About Me")
                        .foregroundColor(.white)
                        .font(.system(size: 30))
                        .fontWeight(.bold)
                    Text("Hey all, I'm a budding ios engineer. Currently exploring swiftUI and developing some cool apps")
                        .foregroundColor(.white)
                        .multilineTextAlignment(.center)
                }.padding().padding()
                Spacer()
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
