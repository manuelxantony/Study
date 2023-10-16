//
//  Wave.swift
//  Wave
//
//  Created by Manuel Antony on 27/06/23.
//

import SwiftUI


struct WaveAnimation: View {

    // controlling the height of the wave
    @Binding var progress: Double
    
    @State var phase: Double = 0.0
    
    var strength: Double = 40
    var frequency: Double = 7
    var isPhaseReverse: Bool = false
   
    var body: some View{
            WaveShape(progress: $progress, isPhaseReverse: isPhaseReverse, strength: strength, frequency: frequency, phase: phase)
            .onAppear{
                withAnimation(Animation.linear(duration: 2).repeatForever(autoreverses: false)){
                    self.phase = .pi * 2
                }
            }
    }
}

struct WaveAnimation_Previews: PreviewProvider {
    @State static var progress: Double = 0.0
    @State static var phase: Double = 0.0
    static var previews: some View {
        WaveAnimation(progress: $progress, phase: phase)
            .foregroundColor(.blue)
    }
}
