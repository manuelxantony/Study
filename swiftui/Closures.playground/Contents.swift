import Foundation

func add(a: Int, b: Int, _ action: (Int) -> ()) {
    let sum = a + b
    action(sum)
}

add(a: 1, b: 2, { sum in
    let squareSum = sum * sum
    print(squareSum)
})


func howareyou(_ responseHandler: @escaping (String) -> ()) {
    print("How are you")
    
    // response after 2 seconds
    DispatchQueue.main.asyncAfter(deadline: .now() + 2, execute: {
            responseHandler("Im gooood....!")
    })
}


howareyou({ response in
    print(response)
})
