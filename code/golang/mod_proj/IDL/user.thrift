namespace go services

struct UserRequest{
    1: string name = "",
}

struct Profile{
    1: i32 ID = 0,
    2: string Description = "",
    3: string Avatar = "",
}

struct UserResponse{
    1: i32 ID = 0,
    2: string Name = "",
    3: string Password = "",
    4: i32 Status = 0,
    5: Profile profile,
}

service UserService {
    UserResponse GetUser(1: UserRequest req)
}