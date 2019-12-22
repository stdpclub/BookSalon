# Book salon V0.2.0

## **It's still being designed and built**

## What it is
It's a individual practice project.

The project is built with `golang` `nginx` `docker`

## Installation

First, you need the environment of `golang`, `gin`, `nginx` and `docker`.
> actually after version1.0。0 you just need a docker

Clone the project.

run `runbooksalon.bat` in windows
or run `./runbooksalon.sh` in unix

...

## How to use

First, your server must running.

Then, here are APIs you can use.

Before you use other API, you must login.

- 登录
  - `localhost:8080/login`

- 获取所有用户
  - `localhost:8080/users`
- 查询某个用户的信息
  - `localhost:8080/user/:userid`
- 新建一个user
  - `localhost:8080/user`
- 删除一个user
  - `localhost:8080/user/:userid`

- 获取user用户的所有team
  - `localhost:8080/user/:userid/teams`
- 新建一个隶属于user的team
  - `localhost:8080/user/:userid/team`
- 获取user用户的某个team的信息
  - `localhost:8080/user/:userid/team/:teamid`
- 更新user下的team的信息
  - `localhost:8080/user/:userid/team/:teamid`
- 删除user下的某个team
  - `localhost:8080/user/:userid/team/:teamid`

- 获取user参加的team的leader
  - `localhost:8080/user/:userid/team/:teamid/leader`
- 获取user参加的team的所有组员
  - `localhost:8080/user/:userid/team/:teamid/members`
- 增加user下的某个team的组员
  - `localhost:8080/user/:userid/team/:teamid/member`
- 删除user下的某个team的某个组员
  - `localhost:8080/user/:userid/team/:teamid/member/:id`

## What's more

thanks for read my project. And if there is any question or problem, please feel free to contact me.
