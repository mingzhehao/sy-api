### 启动 
~~~
bee run -downdoc=true -gendoc=true
~~~


### 配置所有端口，如果80端口被占用，调整所有的文件中的80端口为其他设置端口
  调整包括，js，jquery，index.html，conf配置文件中等等


### 配置fonts.googleapis.com 为国内地址
~~~
http://fonts.useso.com/css
~~~


### 配置外网ip访问（防止跨域问题）
~~~
文件： swagger/swagger-1/index.html
配置： url: "http://192.168.56.101:8085/docs",
~~~

### url 访问swagger
~~~
http://192.168.56.101:8085/swagger/swagger-1/
~~~

### url访问接口两种方式 
** restful url **
~~~
http://192.168.56.101:8085/v1/app/getjson/1
~~~
** router url **
~~~
http://192.168.56.101:8085/getjson/1
~~~


### Post body (添加指定数据)
~~~
    {
        "Id": 1,
        "CreateDate": 1460111310,
        "AppCode": "100005",
        "AppName": "海岛奇兵",
        "PublishData": 1460111310 
    }
~~~
### Update body (更新指定数据)
~~~
    {
        "PublishData": 1460111310
    }
~~~

# Sql 数据 为 
    INSERT INTO `app` VALUES ('1', 1460111310, '100000', '神庙逃亡', null);

