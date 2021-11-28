# GoBIMS
The realization of Books Information Management System(BIMS) with Go+Vue in full stack

## api接口说明
|接口|功能|参数|type|
|---|---|---|---|
|http://locallhost:3000/api/v1/user/login|登录|userpassword,userpassword|POST|
|http://locallhost:3000/api/v1/user/joinup|注册|user_name,pass_word,role|POST|
|http://locallhost:3000/api/v1/user/|用户列表|pagesize,pagenum|GET|
|http://locallhost:3000/api/v1/user/:id|用户信息编辑|id|PUT|
|http://locallhost:3000/api/v1/user/:id|删除用户|id|DELETE|
