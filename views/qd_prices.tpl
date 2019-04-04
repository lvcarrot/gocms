<!DOCTYPE html>
<html>

<head>
  {{template "header" .node.Name}}
  <!-- Latest compiled and minified CSS -->
</head>

<body class="hold-transition skin-blue sidebar-mini">

  <div class="wrapper">
    {{template "navbar" .}}
    <div class="content-wrapper">
      {{template "title" .}}
      <section class="content">
        <!-- Small boxes (Stat box) -->
        <div class="box">
          <div class="box-header with-border">
            <h3 class="box-title"># 价格列表</h3>
            <div class="box-tools">
              <a class="btn btn-sm bg-purple pull-right" data-href="/pdf/qd/prices/new?action=add" data-target="#modal-edit"
                data-toggle="modal">添加 <i class="fa fa-plus"></i>
              </a>
            </div>
          </div>
          {{if .data.list}}
          <div class="box-body table-responsive">
            <table class="table table-bordered">
              <tbody>
                <tr>
                  <th>ID</th>
                  <th>渠道</th>
                  <th>比例</th>
                  <th>单价</th>
                  <th>开始时间</th>
                  <th>操作</th>
                </tr>
                {{range .data.list}}
                <tr>
                  <td>{{.ID}}</td>
                  <td>{{.QD}}</td>
                  <td>{{.Coefficient}}</td>
                  <td>{{.Price}}</td>
                  <td>{{.Start}}</td>
                  <td>
                    <a class="btn btn-default btn-xs" title="渠道价格修改" data-href="/pdf/qd/prices/{{.QD}}?action=edit"
                      data-target="#modal-edit" data-toggle="modal"><i class="fa fa-pencil text-green"></i></a>
                  </td>
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