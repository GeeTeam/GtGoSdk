Gt Golang SDK
===============
极验验证的Golang SDK目前提供基于Beego框架的DEMO。

本项目是面向服务器端的，具体使用可以参考极验的 `文档 <http://www.geetest.com/install/sections/idx-server-sdk.html>`_ ,客户端相关开发请参考极验的 `前端文档。 <http://www.geetest.com/install/>`_.

开发环境
----------------

 - Golang (推荐1.6.2以上版本）
 - Beego (推荐1.6.1以上版本)

快速开始
---------------

下面使用示例代码的均以Beego框架为例。

1. 获取代码

通过go get获取代码:

.. code-block:: bash

    $ go get github.com/GeeTeam/GtGoSdk

2. 初始化验证


在调用GeetestLib前请自行在app.conf设定公钥和私钥,用户id为可选项，默认为空字符串：

.. code-block :: ini

  PrivateKey = "你的私钥"
  CaptchaID = "你的公钥"

根据自己的私钥初始化验证

.. code-block :: go

    func (ctl *RegisterController)Get() {
        userID := ""
        gt := GtGoSdk.GeetestLib(PrivateKey, CaptchaID)
        status := gt.PreProcess(userID)
        ctl.SetSession(GtGoSdk.GT_STATUS_SESSION_KEY, status)
        ctl.SetSession("user_id", userID)
        responseStr := gt.GetResponseStr()
        ctl.Ctx.WriteString(responseStr)
    }

4. 二次验证

.. code-block :: go

    func (ctl *ValidateController)Post() {
        var result bool
        var respstr string
        gt := GtGoSdk.GeetestLib(PrivateKey, CaptchaID)
        challenge := ctl.GetString(GtGoSdk.FN_CHALLENGE)
        validate := ctl.GetString(GtGoSdk.FN_VALIDATE)
        seccode := ctl.GetString(GtGoSdk.FN_SECCODE)
        status := ctl.GetSession(GtGoSdk.GT_STATUS_SESSION_KEY).(int)
        userID := ctl.GetSession("user_id").(string)
        if status == 0 {
            result = gt.FailbackValidate(challenge, validate, seccode)
        } else {
            result = gt.SuccessValidate(challenge, validate, seccode, userID)
        }
        if result {
            respstr = "success"
        } else {
            respstr = "fail"
        }
        ctl.Ctx.WriteString(respstr)
    }


运行demo
---------------------

1. Beego demo运行：进入demo文件夹，运行bee run：

.. code-block:: bash

    $ cd $GOPATH/src/github.com/GeeTeam/GtGoSdk/demo
    $ bee run

在浏览器中访问http://localhost:8080即可看到Demo界面

发布日志
-----------------
+ 3.2.0

 - 参照 `gt-python-sdk <https://github.com/GeeTeam/gt-python-sdk/>`_ 3.2.0版实现极验接口
