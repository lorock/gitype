# typing

简单的博客系统，具有以下特性：

1. 单用户；
1. 无分类，只能通过标签来归类；



### 支持的数据库

1. mysql
1. sqlite3
1. postgresql



### 安装

1. 下载代码:`go get github.com/caixw/typing`；
1. 将typing复制到指定目录；
1. 执行`typing -init=config`输出基本的配置内容；
1. 修改`config`下的相关配置；
1. 执行`typing -init=db`初始化数据库配置；
1. 将源文件目录下的static复制到config/app.json指定的目录中；
1. 登录后台，作一些自定义设置，默认登录密码为`123`；



### 配置文件

#### logs.xml
logs.xml为日志的配置文件，可以定义日志的输入形式和输出日的地，
具体配置可参考[logs](https://github.com/issue9/logs)的相关文档。
文件位于程序当前目录的config子目录下。

#### app.json
app.json位于程序当前目录的config子目录下，包含了以下可配置字段，修改后需要重启程序才能启作用。

名称  | 描述
:---- |:----
debug | 是否处于调试模式
dbDSN | 数据库dsn
dbPrefix | 数据表名前缀
dbDriver | 数据库类型，可以是mysql, sqlite3, postgresql
frontAPIPrefix | 前端api地址前缀
adminAPIPrefix | 后台api地址前经，需要同时修改static/admin/app.js中有关api地址的内容。
themeURLPrefix | 各主题公开文件的根URL
themeDir | 主题文件所在的目录
tempDir | 临时目录



### 开发

typing以自用为主，暂时*不支持新功能的PR*。
bug可在[此处](https://github.com/caixw/typing/issues)提交或是直接PR。

详细的开发文档声明在[DEV](DEV.md)中。



### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
