<%
' ���ܣ���ݵ�¼�ӿڽ���ҳ
' �汾��3.3
' ���ڣ�2012-07-17
' ˵����
' ���´���ֻ��Ϊ�˷����̻����Զ��ṩ���������룬�̻����Ը����Լ���վ����Ҫ�����ռ����ĵ���д,����һ��Ҫʹ�øô��롣
' �ô������ѧϰ���о�֧�����ӿ�ʹ�ã�ֻ���ṩһ���ο���
	
' /////////////////ע��/////////////////
' ������ڽӿڼ��ɹ������������⣬���԰��������;�������
' 1���̻��������ģ�https://b.alipay.com/support/helperApply.htm?action=consultationApply�����ύ���뼯��Э�������ǻ���רҵ�ļ�������ʦ������ϵ��Э�����
' 2���̻��������ģ�http://help.alipay.com/support/232511-16307/0-16307.htm?sh=Y&info_type=9��
' 3��֧������̳��http://club.alipay.com/read-htm-tid-8681712.html��
' /////////////////////////////////////

%>
<html>
<head>
	<META http-equiv=Content-Type content="text/html; charset=gb2312">
<title>֧������ݵ�¼�ӿ�</title>
</head>
<body>

<!--#include file="class/alipay_submit.asp"-->

<%
'/////////////////////�������/////////////////////

        'Ŀ������ַ
        target_service = "user.auth.quick.login"
        '����
        '���ҳ����תͬ��֪ͨҳ��·��
        return_url = "http://�̻����ص�ַ/alipay.auth.authorize-ASP-GBK/return_url.asp"
        '��http://��ʽ������·�����������?id=123�����Զ������
        '������ʱ���
        anti_phishing_key = ""
        '��Ҫʹ����������ļ�submit�е�query_timestamp����
        '�ͻ��˵�IP��ַ
        exter_invoke_ip = ""
        '�Ǿ�����������IP��ַ���磺221.0.0.1

'/////////////////////�������/////////////////////

'���������������
sParaTemp = Array("service=alipay.auth.authorize","partner="&partner,"_input_charset="&input_charset  ,"target_service="&target_service   ,"return_url="&return_url   ,"anti_phishing_key="&anti_phishing_key   ,"exter_invoke_ip="&exter_invoke_ip  )

'��������
Set objSubmit = New AlipaySubmit
sHtml = objSubmit.BuildRequestForm(sParaTemp, "get", "ȷ��")
response.Write sHtml


%>
</body>
</html>
