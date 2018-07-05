<!DOCTYPE html>
<html>
<head>
  {{template "header" .node.Name}}
  <style>
    .select2 {min-width:120px;}
  </style>
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
  {{template "navbar" .}}
  <div class="content-wrapper">
    {{template "title" .}}
    <section class="content">
      <div class="box">
        <div class="box-header with-border">
          <h3 class="box-title">安装活跃</h3>
          <div class="box-tools">
            <form class="form-inline">
              <div class="form-group">
                <label>渠道</label>
                <select class="form-control select2 input-group-sm" name="qd" data-ajax--url="/admin/qd/list" data-ajax--cache="true">
                  <option selected>{{.form.Get "qd"}}</option>
                </select>
              </div>
              <button type="submit" class="btn bg-purple btn-sm" title="筛选">筛选 <i class="fa fa-filter"></i></button>
              <a class="btn bg-olive btn-sm btn-export" href="?export=xls" title="导出">导出 <i class="fa fa-file-excel-o"></i></a>
            </form>
          </div>
        </div>
      {{if .data.list}}
        <div class="box-body table-responsive">
          <table class="table table-bordered">
            <tbody>
              <tr>
                <th>日期</th>
                <th>渠道</th>
                <th>激活</th>
                <th>单价(元)</th>
                <th>结算金额(元)</th>
              </tr>
            {{range .data.list}}
              <tr>
                <td>{{.Date}}</td>
                <td>{{.QD}}</td>
                <td>{{.InstallEnd}}</td>
                <td>{{price .Price}}</td>
                <td>{{price .Total}}</td>
              </tr>
            {{end}}
            </tbody>
          </table>
        </div>
        <div class="box-footer clearfix">
          <a href="javascript:history.go(-1);" class="btn btn-sm bg-navy">返回</a>
          {{template "paginator" .data}}
        </div>
      {{else}}
        <div class="box-body">
          <p class="lead text-center">无数据</p>
        </div>
      {{end}}
      </div>
    </section>
  </div>
  {{template "modal"}}
  {{template "footer"}}
</div>
</body>
</html>