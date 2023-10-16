//
//  IsmometricRectangle.swift
//  ToDo
//
//  Created by Manuel Antony on 08/06/23.
//

import SwiftUI

struct IsmometricRectangle: Shape {
    
    let adjustment: CGFloat
    
    func path(in rect: CGRect) -> Path {
            var path = Path()
            // origin
            path.move(to: CGPoint(x: rect.minX, y: rect.maxY))
            // paths
            path.addLine(to: CGPoint(x: rect.minX + adjustment, y: rect.minY))
            path.addLine(to: CGPoint(x: rect.maxX - adjustment, y: rect.minY))
            path.addLine(to: CGPoint(x: rect.maxX, y: rect.maxY))
        
            return path
        }
}

struct IsmometricRectangle_Previews: PreviewProvider {
    static var previews: some View {
        IsmometricRectangle(adjustment: 20)
            .fill(.yellow)
            .frame(width: 350, height: 100)
    }
}
