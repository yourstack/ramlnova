# !Notice
This Project still in Testing version 0.1.

- **Don’t use it in production.** We can’t support you if things go awry. And upgrading to subsequent versions might involve getting your hands dirty.
- **Report feedbacks responsibly.** Poorly written bug reports take time to deal with, distracting us from adding new features and making ramlnova stable.


# Info
ramlnova is the next-generation RESTful-APIs Generator software that makes you Reduce the 70% server coding workload. It's scalable, intelligent, and fast.

This is ramlnova generator for Multi-Language(Laravel5.2/golang/etc.) project.

这是一个通过解析RAML文件，输出Laravel5.2/Golang等多种语言HTTP RESTful服务端项目代码的生成器.

# Installation
``` 
go get -u github.com/anasinc/ramlnova
```

# FEATURES
Main Features:

- [x] Generate RESTful APIs Server code, 自动生成RESTfulAPIs项目代码

- [x] 快速批量生成文件, faster to generate static *.php/*.lang files

- [x] 多语言模板支持, Multi-Language template file support, [Guide](#Contributing).
``` bash
## template dictionary struct
## .tp writing format guide, please views [Guide]
./template
  laravel5.2/
    routeHead.tp
    xxx.tp
  golang/
    xxx.tp
```

- [ ] Sql-Orm Generator in Cloud, 云端设计生成数据库

Powered by:

- [go-raml(forked buddhamagnet)](https://github.com/buddhamagnet/raml)

- [ramlapi(*.go generator)](https://github.com/EconomistDigitalSolutions/ramlapi)

Compatibility:

- [x] RAML 0.8 support

- [ ] Preliminary 1.0 support

- [ ] Full 1.0 support

# Todolist
FrameWork Main Features:

- [ ] Multi-User template account support.

- [ ] Launch an RESTful APIs Server to Access.

- [ ] Zip up the output project files to .zip/.tar.gz

Laravel5.2:
- [x] route.php Head & Resource Methods Generating from raml.Resourse.Methods()

- [ ] route.php Group/MiddleWare Generating such as Auth/Login/etc.

- [ ] Controllers Generating such as Dictionaries "app/Http/Controllers/xxx/xxx.php"

- [ ] Models Generation from raml.type

# Changelog

# Contributing
## Guide
- template(.tp)
Format Guide : text/template <https://gowalker.org/text/template>

- Reporting Bugs by ISSUE

- Waiting

## Tech Trees
General Language:
- [raml](http://raml.org)
- [golang](http://golang.org)
- [php](http://php.net)
- regular expression

Golang Package:
- [text/template](https://gowalker.org/text/template)
- regexp
- github.com/buddhamagnet/raml

RESTful FrameWork:
- [Laravel 5.2](laravel.com)

Tools:
- vim, vim-go(LiteIDE/Atom/etc. is OK)
- linux, better Debian jessie 8.2 or newer
- nginx, best web server for destributing api server-end 

Extra Skill:
- shell, do some nice trick
- Shadowsocks/VPN, fuck the GFW
- mysql 5.7+

# Contributor

- CIO [ryanemax](https://ryanamax.github.io/), The Founder.

# Copyright

http://www.anasit.com

Copyright © 2015 ANAS IT Inc. All Rights Reserved
