# !Notice
This Project still in Testing version.

- **Don’t use it in production.** We can’t support you if things go awry. And upgrading to subsequent versions might involve getting your hands dirty.
- **Report feedbacks responsibly.** Poorly written bug reports take time to deal with, distracting us from adding new features and making ramlnova stable.


# Info
ramlnova is the next-generation RESTful-APIs Generator software that makes you Reduce the 70% server coding workload. It's scalable, intelligent, and fast.

This is ramlnova generator for Multi-Language(Laravel5.2/golang/etc.) project.

这是一个通过解析RAML文件，输出Laravel5.2/Golang等多种语言HTTP RESTful服务端项目代码的生成器.

# Installation
``` 
go get -u github.com/anasinc/ramlgen
```

# FEATURES
Main Features:

- faster to generate static *.php/*.lang files

Powered by:

- [go-raml(forked buddhamagnet)](https://github.com/buddhamagnet/raml)

- [ramlapi(*.go generator)](https://github.com/EconomistDigitalSolutions/ramlapi)

Compatibility:

- [x] RAML 0.8 support

- [ ] Preliminary 1.0 support

- [ ]Full 1.0 support

# Todolist
FrameWork:
- [x] Multi-Language template file support.
``` bash
## template dictionary struct
./template
  laravel5.2/
    routeHead.tp
    xxx.tp
  golang/
    xxx.tp
```
- [ ] Multi-User template account support.
- [ ] Launch an RESTful APIs Server to Access.
- [ ] Zip up the output project files to .zip/.tar.gz

Laravel5.2:
- [x] raml.Resourse.Methods(),生成基础app/Http/route.php，包括uri部分以及同时根据注释生成的方法名

- [ ] route.php实现auth规则及MiddleWare规则

- [ ] raml.Resourse,生成基础app/Http/Controllers/xxx/xxxController.php，包括文件引用部分及空函数

- [ ] 根据type生成Models基础模型

# Changelog

- 修复生成genFile时，不存在目录问题(MkdirAll/Remove连用生成所需目录)


# Contributing
- template(.tp)
Format Guide : text/template <https://gowalker.org/text/template>

- Reporting Bugs by ISSUE

- Waiting

# Contributor

- CIO [ryanemax](https://ryanamax.github.io/), The Founder.

# Copyright

http://www.anasit.com

Copyright © 2015 ANAS IT Inc. All Rights Reserved
