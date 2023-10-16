//
//  OnBoardingModel.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 30/06/23.
//

import Foundation
import SwiftUI

struct OnBoadingContentModel: Identifiable{
    let id: String = UUID().uuidString
    let title: String
    let lottieFile: String
    let color: Color
    var offSet: CGSize = .zero
}
