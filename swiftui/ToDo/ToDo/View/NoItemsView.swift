//
//  NoItemsView.swift
//  ToDo
//
//  Created by Manuel Antony on 12/06/23.
//

import SwiftUI

struct NoItemsView: View {
    
    @State var animate: Bool = false

    let secondaryColor = Color("SecondaryColor")
    
    var body: some View {
        ScrollView {
            VStack(spacing: 10) {
                Text("There are no items!")
                    .font(.title)
                    .fontWeight(.semibold)
                
                Text("You can add items to accoplish here can be a better productive person.")
                
                NavigationLink(
                    destination: AddItemView(),
                    label: {
                        Text("Add something 🥳")
                            .foregroundColor(.white)
                            .font(.headline)
                            .frame(height: animate ? 55 : 60)
                            .frame(maxWidth: .infinity)
                            .background(animate ?
                                        secondaryColor : .accentColor)
                            .cornerRadius(10)
                            .padding(.top, 30)
                    }
                )
                .padding(.horizontal, animate ? 20 : 60)
                .shadow(
                    color: animate ? secondaryColor.opacity(0.7) : .accentColor.opacity(0.7),
                    radius: animate ? 30 : 10,
                    x: 0,
                    y: animate ? 30 : 50
                )
                .scaleEffect(animate ? 1.1 : 1.0)
                .offset(y: animate ? -10 : 5)
            }
            .multilineTextAlignment(.center)
            .padding(40)
            .onAppear(perform: addAnimation)
        }
    }
    
    func addAnimation() {
        guard !animate else { return }
        DispatchQueue.main.asyncAfter(deadline: .now() + 1.5) {
            withAnimation(
                Animation
                    .easeInOut(duration: 2.0)
                    .repeatForever()
            ) {
                animate.toggle()
            }
        }
    }
}

struct NoItemsView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationView{
            NoItemsView()
                .navigationTitle("")
        }
    }
}
