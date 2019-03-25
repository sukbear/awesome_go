package main

/*1、包的查找顺序：

（1）先去go编译器的安装目录下的src路径中，去查找该包

（2）如果上述路径中查找不到，就去go项目目录下的src路径中，去查找该包


如： import  "go.com/manager/admin/login"

            先去$GOROOT/src/go.com/manager/admin/login 去找login这个包

            再去$GOPATH/src/go.com/manager/admin/login  去找login包*/

/*2、获取goimports 第三方包：

             因为访问的是goole的服务器，所以下载包的时候，会导致连接超时

        如何解决上述问题呢？？

            （1）go  get  -v  github.com/gpmgo/gopm   //绕过防火墙的包

            （2）使用gopm安装第三方包：

                    gopm get -v  -g -u  ... /goimports

                    -g  把包安装到gopath目录中*/
