[TOC]

# 你丫才美工API

## 广告

### addAd

- 接口含义：添加广告
- 请求方式：POST
- 参数：
```
{
    type: 'ninepic' | 'petpic' | 'wildpic', //分别对应九图，萌宠搞事，野表情的广告
    file: file,
    startTime: UNIXtimestamp
    endTime: UNIXtimestamp
}
```
- 响应：
```
{
    status: Number(200 | 500)，   // 200表示成功，其他返回状态用别的状态码标识，并注明message
    data: {
        pid: String // 图片的id
    },
    message: String // 出错了则在这里注明出错信息，成功返回‘成功信息’
}
```

### getAdd

- 接口含义：获取广告信息
- 请求方式：GET
- 参数：
```
{
    type: String('ninepic' | 'petpic' | 'wildpic'), //分别对应九图，萌宠搞事，野表情的广告
}
```
- 响应：
```
{
    status: Number,
    data: {
        type: String('ninepic' | 'petpic' | 'wildpic'), //分别对应九图，萌宠搞事，野表情的广告
        pid: String
    },
    message: String    
}
```

### updateAd

- 接口含义：更新广告
- 请求方式：POST
- 参数：
```
{
    type: 'ninepic' | 'petpic' | 'wildpic', //分别对应九图，萌宠搞事，野表情的广告
    file: file,
    startTime: UNIXtimestamp
    endTime: UNIXtimestamp
}
```
- 响应：
```
{
    status: Number(200 | 500)，   // 200表示成功，其他返回状态用别的状态码标识，并注明message
    data: {
        pid: String // 图片的id
    },
    message: String // 出错了则在这里注明出错信息，成功返回‘成功信息’
}
```

## 用户

### addUser

- 接口含义：添加/修改用户
- 请求方式：POST
- 参数：
```
{
    uid: Number,            // 注意：有这个参数的时候，为修改用户，没有的时候，为添加用户
    username: String,
    password: String,
    authority: Number,
    userLink: String,       // 点击头像跳转的地址，可以为空
    userHeadIcon: String    // 头像的图片地址
}
```
- 响应：
```
{
    status: Number,
    data: null,
    message: String
}
```

### deleteUser

- 接口含义：删除用户
- 请求方式：POST
- 参数：
```
{
    uid: Number
}
```
- 响应：
```
{
    status: Number,
    data: null,
    message: String
}
```
### queryUser

- 接口含义：查询用户
- 请求方式：GET
- 参数：
```
{
    limit: Number,
    page: Number
}
```
- 响应：
```
{
    status: Number,
    data: [
        {
            uid: Number,
            username: String,
            authority: Number,
            userLink: String,       // 点击头像跳转的地址，可以为空
            userHeadIcon: String    // 头像的图片地址
        }
    ],
    message: String
}
```

### login

- 接口含义：用户登录
- 请求方式：POST
- 参数：
```
{
    username: String,
    password: String
}
```
- 响应
```
{
    status: Number,
    data: {
        uid: Number,
        username: String,
        authority: Number,
        userLink: String,       // 点击头像跳转的地址，可以为空
        userHeadIcon: String    // 头像的图片地址
    },
    message: String
}
```

### logout

- 接口含义：用户登出
- 请求方式：GET
- 参数：无
- 响应：
```
{
    status: Number,
    data: null,
    message: String
}
```

## 图片管理


## 野表情和萌宠

