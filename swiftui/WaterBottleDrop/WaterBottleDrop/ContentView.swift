//
//  ContentView.swift
//  WaterBottleDrop
//
//  Created by Manuel Antony on 27/06/23.
//

import SwiftUI

struct ContentView: View {
    
    @State var progress: Double = 0.0
    @State var isFill: Bool = false
    
    // 370 is the height occupied by the drop shape
    // in the total height of the image view
    let fillValue: Double = 370
    
    var body: some View {
        VStack{
            
            ZStack{
                Image(systemName: "drop.fill")
                    .resizable()
                    .aspectRatio(contentMode: .fit)
                    .foregroundColor(.white)
                    .scaleEffect(x: 1.1, y: 1)
                    .frame(width: 250, height: 500)
                 
                // for waves
                ZStack{
                    WaveAnimation(progress: $progress, strength: 18, frequency: 2)
                        .foregroundColor(.blue.opacity(0.5))

                    WaveAnimation(progress: $progress, strength: 15, frequency: 7,  isPhaseReverse: true)
                        .foregroundColor(.blue)
                }
                .mask{
                    Image(systemName: "drop.fill")
                        .resizable()
                        .aspectRatio(contentMode: .fit)
                        .padding(5)
                }
                .overlay {
                    ZStack {
                        Circle()
                            .fill(.white.opacity(0.4))
                            .frame(width: 15)
                            .offset(x: 40, y: 10)
                        
                        Circle()
                            .fill(.white.opacity(0.1))
                            .frame(width: 20)
                        
                        Circle()
                            .fill(.white.opacity(0.4))
                            .frame(width: 10)
                            .offset(x: -10, y: -90)

                        Circle()
                            .fill(.white.opacity(0.2))
                            .frame(width: 15)
                            .offset(x: -30, y: -50)
                        
                        Circle()
                            .fill(.white.opacity(0.1))
                            .frame(width: 15)
                            .offset(x: -70, y: 50)
                        
                        Circle()
                            .fill(.white.opacity(0.1))
                            .frame(width: 15)
                            .offset(x: 50, y: 70)
                        
                        Circle()
                            .fill(.white.opacity(0.4))
                            .frame(width: 15)
                            .offset(x: -40, y: 130)
                    }
                }
            }
            .frame(width: 250)
            
            Button {
                // increase the water level
                progress += 10
                if progress == fillValue {
                    isFill = true
                }

            } label: {
                    Image(systemName: "plus")
                        .font(.system(size: 40, weight: .black))
                        .foregroundColor(.white)
                        .padding(10)
                        .background(.black.opacity(0.4), in: Circle())
                        .shadow(color: .black.opacity(0.9), radius: 4, x: 4, y: 4)
                        .shadow(color: .white.opacity(0.1), radius: 4, x: -5, y: -5)
                
                
            }
            .padding(.bottom, 100)
            // disabling button if water wave is filled in the drop/bottle
            .disabled(isFill)
            
            
        }
        .padding()
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(.black.opacity(0.8))
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
