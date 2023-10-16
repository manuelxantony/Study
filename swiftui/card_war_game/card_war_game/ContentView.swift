//
//  ContentView.swift
//  card_war_game
//
//  Created by Manuel Antony on 31/05/23.
//

import SwiftUI

struct ContentView: View {
    
    @State var playerCard:String = "back"
    @State var cpuCard:String = "back"
    
    @State var playerCount: Int = 0
    @State var cpuCount:Int = 0
    
    var body: some View {
        ZStack(alignment: .center){
            Image("background-cloth").resizable()
                .ignoresSafeArea()
            // for align all items
            VStack(alignment: .center, spacing: 40){
                HStack(alignment: .bottom) {
                    Image("logo")
                    Button {
                        refresh()
                    } label: {
                        Text("restart")
                            .foregroundColor(.white)
                    }.padding(.bottom, 200).padding(.leading, 40)
                }.padding(.leading, 120)
                
                // for displaying cards
                HStack(alignment: .center, spacing: 50){
                    Image(playerCard)
                    Image(cpuCard)
                }.padding(.bottom, 40)
                
                // deal button
                Button {
                    deal() // action
                } label: {  // label we can give the UI
                    Image("button")
                }

                
                // count - player vs cpu
                HStack(alignment: .center, spacing: 60){
                    // player
                    VStack(alignment: .center, spacing: 20){
                        Text("Player")
                            .foregroundColor(.white)
                            .font(.system(size: 20))
                            .fontWeight(.semibold)
                        Text(String(playerCount))
                            .foregroundColor(.white)
                            .font(.system(size: 30))
                    }
                    //cpu
                    VStack(alignment: .center, spacing: 20){
                        Text("CPU")
                            .foregroundColor(.white)
                            .font(.system(size: 20))
                            .fontWeight(.semibold)
                        Text(String(cpuCount))
                            .foregroundColor(.white)
                            .font(.system(size: 30))
                    }
                }
            }
        }
    }
    
    func deal(){
        // randomize cards
        var playerCardValue = Int.random(in: 2...14)
        playerCard = "card" + String(playerCardValue)
        
        var cpuCardValue = Int.random(in: 2...14)
        cpuCard = "card" + String(cpuCardValue)
        
        // update the scores
        if playerCardValue > cpuCardValue {
            playerCount += 1
        } else if cpuCardValue > playerCardValue{
            cpuCount += 1
        }
    }
    
    func refresh(){
        playerCount = 0
        cpuCount = 0
        playerCard = "back"
        cpuCard = "back"
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
