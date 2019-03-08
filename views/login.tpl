<!DOCTYPE html>
<html>

<head>
  {{template "header" "后台登录"}}
  <style type="text/css">
    footer {
      margin-left: 0px !important;
      position: absolute;
      width: 100%;
      bottom: 0;
    }
  </style>
</head>

<body class="hold-transition login-page">
  <div class="login-box">
    <div class="login-logo">
      <a href="#"><b>Go</b>CMS</a>
    </div>
    <div class="login-box-body box">
      <p class="login-box-msg">登录系统后台</p>
      <form action="/login?refer={{urlquery .ref}}" method="post">
        <div class="form-group has-feedback">
          <input name="username" type="email" class="form-control" autocomplete="off" placeholder="请输入管理员邮箱"
            data-message="{'required':'登录名称不能为空'}" required>
          <span class="glyphicon glyphicon-user form-control-feedback"></span>
        </div>
        <div class="form-group has-feedback">
          <input name="password" type="password" class="form-control" autocomplete="off" placeholder="请输入密码"
            data-message="{'required':'密码不能为空'}" required>
          <span class="glyphicon glyphicon-lock form-control-feedback"></span>
        </div>
        <div class="row">
          <div class="col-xs-6 form-group has-feedback">
            <input name="code" type="text" class="form-control" autocomplete="off" placeholder="验证码" data-message="{'required':'验证码不能为空'}"
              required>
          </div>
          <div class="col-xs-4 pull-right">
            <input type="hidden" name="id" value="{{.captcha}}">
            <button type="submit" class="btn btn-primary btn-block btn-flat">登录</button>
          </div>
        </div>
        <img class="img-responsive" src="/captcha/{{.captcha}}.png?">
      </form>
    </div>
  </div>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/3.4.1/js/bootstrap.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/admin-lte/2.4.9/js/adminlte.min.js"></script>

  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery.form/4.2.2/jquery.form.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.19.0/jquery.validate.min.js"></script>
  <script src="//cdnjs.cloudflare.com/ajax/libs/blueimp-md5/2.10.0/js/md5.min.js"></script>
  <script src="/static/js/global.js?v=20181213" type="text/javascript"></script>
  <script type="text/javascript">
    $(document).on('click', 'img', function (e) {
      var src = $(e.target).attr('src');
      $(e.target).attr('src', src.substr(0, src.indexOf('?') + 1) +
        'reload=' + (new Date()).getTime());
    })
  </script>
</body>

</html>