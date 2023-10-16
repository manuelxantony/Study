//
//  test.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 29/06/23.
//

import SwiftUI

struct Home2: View {
    
    // Intros....
    @State var intros: [Intro] = [
    
        Intro(title: "Plan", subTitle: "your routes", description: "View your collection route Follow, change or add to your route yourself", pic: "circle",color: Color("Green")),
        Intro(title: "Quick Waste", subTitle: "Transfer Note", description: "Record oil collections easily and accurately. No more paper!", pic: "square.and.arrow.down.fill",color: Color("DarkGrey")),
        Intro(title: "Invite", subTitle: "restaurants", description: "Know some restaurant who want to optimize oil collection? Invite them with one click", pic: "pencil.tip",color: Color("Yellow")),
    ]
    
    
    var body: some View {
        
        ZStack{
            ForEach(intros.indices.reversed(),id: \.self){index in
                IntroView(intro: intros[index])
                    .ignoresSafeArea()
            }
        }
    }
}









struct Home2_Previews: PreviewProvider {
    static var previews: some View {
        Home2()
    }
}

// Intro Model and Sample Screens...
struct Intro: Identifiable{
    var id = UUID().uuidString
    var title: String
    var subTitle: String
    var description: String
    var pic: String
    var color: Color
    var offset: CGSize = .zero
}


struct IntroView: View {
    let intro: Intro
    
    var body: some View {
        VStack{
            
            Image(systemName: intro.pic)
                .resizable()
                .aspectRatio(contentMode: .fit)
                .padding(40)
            
            VStack(alignment: .leading, spacing: 0) {
                
                Text(intro.title)
                    .font(.system(size: 45))
                
                Text(intro.subTitle)
                    .font(.system(size: 50, weight: .bold))
                
                Text(intro.description)
                    .font(.system(size: 20))
                    .fontWeight(.semibold)
                    .padding(.top)
                    .frame(width: .infinity - 100)
                    .lineSpacing(8)
            }
            .foregroundColor(.white)
            .frame(maxWidth: .infinity,alignment: .leading)
            .padding(.leading,20)
            .padding([.trailing,.top])
        }
        .frame(maxWidth: .infinity,maxHeight: .infinity)
        .background(
        
            intro.color
        )
    }
}
