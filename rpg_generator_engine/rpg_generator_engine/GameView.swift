//
//  ContentView.swift
//  rpg_generator_engine
//
//  Created by Caim on 2026/4/15.
//

import SwiftUI

struct GameView: View {
    var body: some View {
        VStack {
            Image(systemName: "square")
                .imageScale(.large)
                .foregroundStyle(.tint)
            Text("Hello, world!")
        }
        .padding()
    }
}

#Preview {
    GameView()
}
