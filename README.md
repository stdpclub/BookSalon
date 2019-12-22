# Book salon V0.2.0

## **It's still being designed and built**

## What it is
It's a individual practice project.

The project is built with `golang` `nginx` `docker`

## Installation

First, you need the environment of `golang`, `gin`, `nginx` and `docker`.
> actually after version1.0.0 you just need a docker

Clone the project.

run `runbooksalon.bat` in windows
or run `./runbooksalon.sh` in unix

...

## How to use

First, your server must running.

Then, here are APIs you can use.

Before you use other API, you must login.

Please change `localhost:8080` to yourself address

- 登录
  - GET => `localhost:8080/login`
- 获取所有用户
  - GET => `localhost:8080/user/`
- 查询某个用户的信息
  - GET => `localhost:8080/user/:userid`
- 新建一个user
  - POST => `localhost:8080/user/`
- 删除一个user
  - DELETE => `localhost:8080/user/:userid`
- 获取user用户的所有team
  - GET => `localhost:8080/user/:userid/teams`
- 新建一个隶属于user的team
  - POST => `localhost:8080/user/:userid/team`
- 获取user用户的某个team的信息
  - GET => `localhost:8080/user/:userid/team/:teamid`
- 更新user下的team的信息
  - PUT => `localhost:8080/user/:userid/team/:teamid`
- 删除user下的某个team
  - DELETE => `localhost:8080/user/:userid/team/:teamid`
- 获取user参加的team的leader
  - GET => `localhost:8080/user/:userid/team/:teamid/leader`
- 获取user参加的team的所有组员
  - GET => `localhost:8080/user/:userid/team/:teamid/members`
- 增加user下的某个team的组员
  - POST => `localhost:8080/user/:userid/team/:teamid/member`
- 删除user下的某个team的某个组员
  - DELETE => `localhost:8080/user/:userid/team/:teamid/member/:id`


## What's more

Thanks for read my project. And if there is any question or problem, please feel free to contact me.
