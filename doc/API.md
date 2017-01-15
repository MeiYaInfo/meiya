[TOC]

# 你丫才美工API

## 广告

### editAd

- 接口含义：添加/修改广告
- 请求方式：POST
- 参数：
```
{
    adId: Number,   // 如果没有这个属性就是添加，有的话就是修改
    type: 'ninepic' | 'petpic' | 'wildpic', //分别对应九图，萌宠搞事，野表情
    pid: String,    // 图片的id
    startTime: Number(UNIXtimestamp),
    endTime: Number(UNIXtimestamp)
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
    type: String('ninepic' | 'petpic' | 'wildpic'), //分别对应九图，萌宠搞事，野表情
}
```
- 响应：
```
{
    status: Number,
    data: {
        type: String('ninepic' | 'petpic' | 'wildpic'), //分别对应九图，萌宠搞事，野表情
        pid: String
    },
    message: String    
}
```
### deleteAd

- 接口含义：删除广告
- 请求方式：POST
- 参数：
```
{
    adId: number
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

## 用户

### editUser

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
### getUser

- 接口含义：查询用户列表
- 请求方式：GET
- 参数：
```
{
    uid: Number,    // 有这个字段就只返回该用户信息，没有则接收liimt和page字段返回用户列表
    limit: Number,
    page: Number,
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

### getLoginInfo

- 接口含义：查询当前登录用户信息
- 请求方式：GET
- 参数：无
- 响应：
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

### uploadPic

- 接口含义：添加图片
- 请求方式：POST
- 参数：
```
{
    file: file
}
```
- 响应：
```
{
    status: Number,
    data: {
        pid: String
    },
    message: String
}
```

### editNinePic

- 接口含义：添加九图文章
- 请求方式：POST
- 参数：
```
{
    nid: Number,    // 有则是修改，没有就是添加
    pids: String,   // 由多个图片的pid拼接而成的字符串，用逗号分开
    uid: String, // 只能是user用户表中的用户uid
    description: String,
    original: Boolean
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

### getNinePic

- 接口含义：获取九图
- 请求方式：GET
- 参数：
```
{
    limit: Number,  // 有这个属性就是获取nid之前的limit个列表，没有这个属性就是获取nid这条九图的信息 
    nid: Number
}
```
- 响应：
```
{
    status: Number,
    data: [
        {
            pids: String,   // 由多个图片的pid拼接而成的字符串，用逗号分开
            uid: String, // 只能是user用户表中的用户uid
            description: String,
            original: Boolean
        }
    ]
    message: String
}
```

### deleteNinePic

- 接口含义：删除九图
- 请求方式：post
- 参数：
```
{
    nid: Number
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

### getPic

查询野表情或萌宠表情

- 接口含义：查询野表情/萌宠信息
- 请求方式：get
- 参数：
```
{
    type: String('petpic' | 'wildpic'),
    limit: Number,  // 需要显示的天数
    time: Number   // 需要此时间之前的几天的数据
}
```
- 响应：
```
{
    status: Number,
    data: [
        {
            pid: String,
            addTime: String(UNIXstamp)
        }
    ],
    message: String
}
```
























