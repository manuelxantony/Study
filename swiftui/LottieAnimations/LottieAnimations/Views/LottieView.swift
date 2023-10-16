//
//  LottieView.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 28/06/23.
//

import SwiftUI
import Lottie


// MARK: UIViewRepresentable is used for creating and managing UIView in SwiftUI
struct LottieView: UIViewRepresentable {
       
    // name of the lottie json file
    let lottieFile: String
    // creating lottie animation view
    let animationView = LottieAnimationView()
    
    // loopmode
    var loopMode: LottieLoopMode = .loop
    
    // MARK: Creates the view object and configures its initial state.
    func makeUIView(context: Context) -> some UIView {
        // creating UIView instance
        let view = UIView(frame: .zero)
        
        animationView.animation = LottieAnimation.named(lottieFile)
        animationView.contentMode = .scaleAspectFit
        animationView.loopMode = loopMode
        animationView.play()
        
        // adding view to receiver, superview
        view.addSubview(animationView)
        
        // authoresization for translating views mask into auto layouts
        animationView.translatesAutoresizingMaskIntoConstraints = false
        
        // height and width constraints of the view frame
        NSLayoutConstraint.activate([
            animationView.heightAnchor.constraint(equalTo: view.heightAnchor),
            animationView.widthAnchor.constraint(equalTo: view.widthAnchor),
        ])
        
        return view
    }
    
    // for updating view with new data
    func updateUIView(_ uiView: UIViewType, context: Context) {}
}

struct LottieView_Previews: PreviewProvider {
    static let lottieFile: String = "liquid-water-drop"
    static var previews: some View {
        LottieView(lottieFile: lottieFile)
            .background(.black.opacity(0.9))
    }
}
