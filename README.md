# go-qywechat

![GitHub issues](https://img.shields.io/github/issues/xqk/go-qywechat)
![GitHub forks](https://img.shields.io/github/forks/xqk/go-qywechat)
![GitHub stars](https://img.shields.io/github/stars/xqk/go-qywechat)
![GitHub license](https://img.shields.io/github/license/xqk/go-qywechat)
![Twitter](https://img.shields.io/twitter/url?url=https%3A%2F%2Fgithub.com%2Fxqk%2Fgo-qywechat)

```
import (
    "github.com/xqk/go-qywechat" // package qywechat
)
```
<details>
<summary>客户联系 API</summary>

* [ ] 成员对外信息
* [ ] 企业服务人员管理
  - [ ] 获取配置了客户联系功能的成员列表
  - [x] 客户联系「联系我」管理
* [x] 客户管理
  - [x] 获取客户列表
  - [x] 获取客户详情
  - [x] 批量获取客户详情
  - [ ] 修改客户备注信息
  - [ ] 客户联系规则组管理
* [ ] 在职继承
  - [ ] 分配在职成员的客户
  - [ ] 查询客户接替状态
  - [ ] 分配在职成员的客户群
* [ ] 离职继承
  - [ ] 获取待分配的离职成员列表
  - [ ] 分配离职成员的客户
  - [ ] 查询客户接替状态
  - [ ] 分配离职成员的客户群
* [ ] 客户标签管理
  - [ ] 管理企业标签
  - [ ] 编辑客户企业标签
* [ ] 客户分配
  - [ ] 获取离职成员列表
  - [ ] 分配在职或离职成员的客户
  - [ ] 查询客户接替结果
  - [ ] 分配离职成员的客户群
* [x] 变更回调通知
  - [x] 添加企业客户事件
  - [x] 编辑企业客户事件
  - [x] 外部联系人免验证添加成员事件
  - [x] 删除企业客户事件
  - [x] 删除跟进成员事件
  - [x] 客户接替失败事件
  - [x] 客户群变更事件

</details>