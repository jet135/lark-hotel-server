# 基于飞书搭建的酒店管理服务
没有前端代码，完全基于飞书多维表格搭建，以飞书表格数据为主，程序对接飞书开放平台接口将表格数据同步至Mysql
web框架为Goframe
这是一个业务属性比较重的项目，实际上通用性不强，但是这个服务实现了信息数字化，让数据能够被使用，成本仅为一台服务器费用（飞书个人企业版是免费使用的），作为一个go web服务，希望对大家有参考意义

## Features
- Chrome浏览器插件（manifest/chrome_plugin）实现身份证录入时，调用服务器接口上报信息，服务器调用飞书接口将信息新增至表格
- 飞书表格内记录客房信息（房号、房价等..），多维表格实时展示数据报表
- 定时生成每天的飞书表格，以及初始化数据
- 定时同步飞书数据到Mysql，记录入住多日的用户，第二天生成表格时进行数据初始化
- 接入飞书机器人，在聊天中输入指令进行业务操作（快捷退房、查询用户信息等..）
- 年度统计报表功能，数据定时同步至报表
- 查询某日账单信息（飞书机器人每日定时推送消息、通过Iphone快捷指令查询）

## Lark Word Template Link
创建副本到自己的云档目录中

运营数据-多维表格模板：
https://da4jh6e6v50.feishu.cn/base/WT4Nbxh1ua1ybfsCJKzcjukanzb
年度账单-多维表格模板：
https://da4jh6e6v50.feishu.cn/base/IKZQblceDaAuD8sa0YQcon9Sn3f


## Quick Start
1、Clone the repository:
```bash
git clone https://github.com/jet135/lark-hotel-server.git
```
2、Profile configuration
```bash
# 在 manifest/config/下创建config.yaml并配置，参考manifest/config/config.example.yaml
```
(manifest/config/config.yaml)
3、Initialize database table(manifest/sql/lark-hotel-server.sql)
4、Run locally
```bash
# prefix：Need to install gf
go mod tidy
gf run main.go
```

## Deploy To Server

