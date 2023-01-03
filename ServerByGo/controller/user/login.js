function loginHome(){
    var username =$("#username").val();//#是id选择器
    var password =$("#password").val();
    if(username==""){
        alert("用户名不能为空")
        document.getElementById("tip").innerText="用户名不能为空";
        return false;
    }
    if(password==""){
        document.getElementById("tip").innerText="密码不能为空";
        return false;
    }
    else{
        $.ajax({
            type:"post",       //请求类型
            url:"http://120.48.100.7:8020/login",	//请求地址
            dataType:"json",	//预期返回的参数类型
            async: true,			//true代表的是异步请求，false代表的是同步请求，若为同步请求则会锁住浏览器
            // data:{username:username,password:password},    //传入后端Controller的参数
            success:function(data){				//这里的data是后台返回的值
                var password = data.password;//提取出来密码
                if(password != password){
                    document.getElementById("tip").innerText="账号或密码输入错误";
                }else{
                    window.location.href="http://120.48.100.7:8020/";
                }
            }
        });
    }
}
