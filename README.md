# GoBIMS
The realization of Books Information Management System(BIMS) with Go+Vue in full stack

## api接口说明

|接口|功能|参数|type|
|:-:|:-:|:-:|:-:|
|http://localhost:3000/api/v1/admin/userlist|获取用户列表(或者根据用户名搜索)|pagesize,pagenum,(username)|POST|
|http://localhost:3000/api/v1/admin/editbook/:id|修改(:id)的图书信息|(:id),name,categories,author,price,date|PUT|
|http://localhost:3000/api/v1/admin/addbook|增加图书信息|name,categories,author,price,date|POST|
|http://localhost:3000/api/v1/admin/deletebook/:id|删除(:id)的图书|id|DELETE|
|http://localhost:3000/api/v1/admin/login|管理员登录|user_name,pass_word|POST|
|http://localhost:3000/api/v1/user/:id|删除(:id)的用户|id|DELETE|
|http://localhost:3000/api/v1/user/:id|修改(:id)的用户信息|(:id),user_name,role|PUT|
|http://localhost:3000/api/v1/base/booklist|获取图书信息列表|pagesize,pagenum|GET|
|http://localhost:3000/api/v1/base/searchlist|根据书名、类别搜索图书|name,categories,pagesize,pagenum|GET|
|http://localhost:3000/api/v1/base/joinup|用户注册|user_name,pass_word,role|POST|
|http://localhost:3000/api/v1/base/login|用户登录|user_name,pass_word|POST|