//
//  LiquidSwip.swift
//  LottieAnimations
//
//  Created by Manuel Antony on 29/06/23.
//

import SwiftUI

struct LiquidLeftSwip: Shape {
    
    var offset: CGSize
    
    var animatableData: CGSize{
        get { return offset }
        set { offset = newValue }
    }
    
    // top and bottom of the drop shape
    let yTop: CGFloat = 600
    let yBottom: CGFloat = 700
    // x axis of the shape, this decides the initial bulging
    let xOffset: CGFloat = 50
    
    // y axis of the curve in the drop shape
    var yCurve: CGFloat {
        (yTop + yBottom) / 2
    }
    
    func path(in rect: CGRect) -> Path {
        var path = Path()
        
        let rectWidth = rect.width
        let rectHeight = rect.height
        let offsetWidth = (offset.width >= 0) ? 0 : offset.width
        let offsetHeight = offset.height
        // double the offset.width value for showing smooth expansion
        let offsetWidthDouble = offsetWidth * 2
        
        
        // rectangle
        path.move(to:CGPoint(x:0, y: 0))
        path.addLine(to: CGPoint(x: rectWidth, y: 0))
        path.addLine(to: CGPoint(x: rectWidth, y: rectHeight))
        path.addLine(to: CGPoint(x: 0, y: rectHeight))
        
        
        // setting the path for drop shape
        path.move(to: CGPoint(x: rectWidth, y: yTop + offsetWidthDouble + offsetHeight))
        path.addCurve(to: CGPoint(x: rectWidth, y: yBottom - offsetWidthDouble + offsetHeight),
                      control1: CGPoint(x: rectWidth - xOffset + offsetWidth, y: yCurve + offsetHeight),
                      control2: CGPoint(x: rectWidth  - xOffset + offsetWidth, y: yCurve + offsetHeight))
        path.move(to: CGPoint(x: rectWidth, y: yTop + offsetWidthDouble + offsetHeight))
        
        return path
    }
}

struct LiquidLeftSwip_Previews: PreviewProvider {
    //@State static var offset: CGSize = CGSize(width: 0, height: 0)
    static var previews: some View {
        //LiquidLeftSwip(offset: offset)
        Home()
    }
}
