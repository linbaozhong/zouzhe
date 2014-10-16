<%@LANGUAGE="VBSCRIPT" CODEPAGE="65001"%>
<%
' 功能：快捷登录接口接入页
' 版本：3.3
' 日期：2012-07-17
' 说明：
' 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写,并非一定要使用该代码。
' 该代码仅供学习和研究支付宝接口使用，只是提供一个参考。
	
' /////////////////注意/////////////////
' 如果您在接口集成过程中遇到问题，可以按照下面的途径来解决
' 1、商户服务中心（https://b.alipay.com/support/helperApply.htm?action=consultationApply），提交申请集成协助，我们会有专业的技术工程师主动联系您协助解决
' 2、商户帮助中心（http://help.alipay.com/support/232511-16307/0-16307.htm?sh=Y&info_type=9）
' 3、支付宝论坛（http://club.alipay.com/read-htm-tid-8681712.html）
' /////////////////////////////////////

%>
<html>
<head>
	<META http-equiv=Content-Type content="text/html; charset=utf-8">
<title>支付宝快捷登录接口</title>
</head>
<body>

<!--#include file="class/alipay_submit.asp"-->

<%
'/////////////////////请求参数/////////////////////

        '目标服务地址
        target_service = "user.auth.quick.login"
        '必填
        '必填，页面跳转同步通知页面路径
        return_url = "http://商户网关地址/alipay.auth.authorize-ASP-UTF-8/return_url.asp"
        '需http://格式的完整路径，不允许加?id=123这类自定义参数
        '防钓鱼时间戳
        anti_phishing_key = ""
        '若要使用请调用类文件submit中的query_timestamp函数
        '客户端的IP地址
        exter_invoke_ip = ""
        '非局域网的外网IP地址，如：221.0.0.1

'/////////////////////请求参数/////////////////////

'构造请求参数数组
sParaTemp = Array("service=alipay.auth.authorize","partner="&partner,"_input_charset="&input_charset  ,"target_service="&target_service   ,"return_url="&return_url   ,"anti_phishing_key="&anti_phishing_key   ,"exter_invoke_ip="&exter_invoke_ip  )

'建立请求
Set objSubmit = New AlipaySubmit
sHtml = objSubmit.BuildRequestForm(sParaTemp, "get", "确认")
response.Write sHtml


%>
</body>
</html>
