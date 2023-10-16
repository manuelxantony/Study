//
//  temp.swift
//  neopoptest
//
//  Created by Manuel Antony on 12/06/23.
//

import SwiftUI
import NeoPop

class Temp {
    
    let model = PopView.Model(
        popEdgeDirection: .bottomRight,
        backgroundColor: Color(.black)
    )

    let popView = PopView(
        frame: CGRect(x: 0, y: 0, width: 100, height: 100),
        model: model
    )
}
