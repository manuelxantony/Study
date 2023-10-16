enum UserError: Error{
    case invalidId
    case invalidUser
}

func userDetails(id: Int, name: String) throws{
    if id == 0{
        throw UserError.invalidId
    }
    if name == ""{
        throw UserError.invalidUser
    }
}

do {
    try userDetails(id: 0, name: "")
}catch UserError.invalidId{
    print("Invalid used ID")
}catch UserError.invalidUser{
    print("Invalid user")
}
