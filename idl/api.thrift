namespace go api

struct Req{   // api请求参数,由于user服务需要id和name,因此api服务的请求参数也是id和name
    1:i64 id (api.query="id")
    2:string name (api.query="name")
}

struct Resp{
    1:i64 code
    2:string msg
}

service ApiService{
    Resp CreatUser(1:Req req)(api.get="/user") // 本demo通过http的get请求获取id和name,通过调用user服务在服务端数据库创建一条记录
}