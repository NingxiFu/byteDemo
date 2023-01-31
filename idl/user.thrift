namespace go user

struct Req{
    1:i64 id
    2:string name
}

struct Resp{
    1:i64 code
    2:string msg
}

service UserService{
    Resp CreatUser(1:Req req)
}