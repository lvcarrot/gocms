<!DOCTYPE html>
<html>

<head>
  {{template "header" .node.Name}}
  <style>
    .select2 {
      min-width: 120px;
    }
  </style>
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/bootstrap-fileinput/4.5.2/css/fileinput.min.css" />
</head>

<body class="hold-transition skin-blue sidebar-mini">
  <div class="wrapper">
    {{template "navbar" .}}
    <div class="content-wrapper">
      {{template "title" .}}
      <section class="content">
        <div class="row">
          <div class="col-md-6">
            <div class="box box-solid">
              <div class="box-header with-border">
                <h3 class="box-title">官网版本</h3>
                <div class="box-tools">
                  <a class="btn bg-primary btn-sm" data-href="/modal/pdf/publish?type=WebSite" data-target="#modal-edit"
                    data-toggle="modal" title="修改">修改 <i class="fa fa-pencil-square-o"></i></a>
                </div>
              </div>
              {{if .data.web }}
              {{$v := .data.web.Version}}
              <div class="box-body box-profile">
                <ul class="list-group list-group-unbordered">
                  <li class="list-group-item">
                    <b>版本号</b> <a class="pull-right">{{ $v.Version }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>更新类型</b> <a class="pull-right">{{updateType $v.UpdateType }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>版本类型</b> <a class="pull-right">{{versionType $v.VersionType }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>发布说明</b> <a class="pull-right">{{ $v.ReleaseNote }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>发布时间</b> <a class="pull-right">{{date $v.ReleaseDate }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>包大小</b> <a class="pull-right">{{ $v.PkgSize }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>MD5</b> <a class="pull-right">{{ $v.MD5 }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>发布人</b> <a class="pull-right">无</a>
                  </li>
                  <li class="list-group-item">
                    <b>下载地址</b> <a class="pull-right">{{ $v.PkgURL }}</a>
                  </li>
                </ul>
              </div>
              {{else}}
              <div class="box-body">
                <p class="lead text-center">无数据</p>
              </div>
              {{end}}
            </div>
          </div>
          <div class="col-md-6">
            <div class="box box-solid">
              <div class="box-header with-border">
                <h3 class="box-title">更新接口版本</h3>
                <div class="box-tools">
                  <form class="form-inline">
                    <a class="btn bg-primary btn-sm" data-href="/modal/pdf/publish?type=Api" data-target="#modal-edit"
                      data-toggle="modal" title="修改">修改 <i class="fa fa-pencil-square-o"></i></a>
                  </form>
                </div>
              </div>
              {{if .data.api }}
              {{$v := .data.api.Version}}
              <div class="box-body box-profile">
                <ul class="list-group list-group-unbordered">
                  <li class="list-group-item">
                    <b>版本号</b>
                    <a class="pull-right">{{ $v.Version }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>更新类型</b> <a class="pull-right">{{updateType $v.UpdateType }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>版本类型</b> <a class="pull-right">{{versionType $v.VersionType }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>发布说明</b> <a class="pull-right">{{ $v.ReleaseNote }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>发布时间</b> <a class="pull-right">{{date $v.ReleaseDate }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>包大小</b> <a class="pull-right">{{ $v.PkgSize }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>MD5</b> <a class="pull-right">{{ $v.MD5 }}</a>
                  </li>
                  <li class="list-group-item">
                    <b>发布人</b> <a class="pull-right">无</a>
                  </li>
                  <li class="list-group-item">
                    <b>下载地址</b> <a class="pull-right">{{ $v.PkgURL }}</a>
                  </li>
                </ul>
              </div>
              {{else}}
              <div class="box-body">
                <p class="lead text-center">无数据</p>
              </div>
              {{end}}
            </div>
          </div>
        </div>
        <div class="box box-info">
          <div class="box-header with-border">
            <h3 class="box-title">版本列表</h3>
            <div class="box-tools">
              <a class="btn btn-sm bg-purple pull-right" data-href="/pdf/versions/new" data-target="#modal-detail"
                data-toggle="modal">添加 <i class="fa fa-plus"></i>
              </a>
            </div>
          </div>
          {{if .data.list}}
          <div class="box-body table-responsive">
            <table class="table table-bordered">
              <tbody>
                <tr>
                  <th>版本号</th>
                  <th>版本类型</th>
                  <th>更新类型</th>
                  <th>官网</th>
                  <th>API</th>
                  <th>大小</th>
                  <th>URL</th>
                  <th>发布时间</th>
                  <th>操作</th>
                </tr>
                {{ range .data.list }}
                <tr>
                  <td>{{ .Version.Version }}</td>
                  <td>{{versionType .VersionType }}</td>
                  <td>{{updateType .UpdateType }}</td>
                  <td>{{boolNote .ReleaseOnWeb }}</td>
                  <td>{{boolNote .ReleaseOnApi }}</td>
                  <td>{{ .PkgSize }}</td>
                  <td>{{ .PkgURL }}</td>
                  <td>{{date .ReleaseDate }}</td>
                  <td>
                    <a class="btn btn-default btn-xs" title="版本修改" data-href="/pdf/versions/{{.Version.Version}}" data-target="#modal-detail"
                      data-toggle="modal"><i class="fa fa-pencil text-green"></i></a>
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
    <script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-fileinput/4.5.2/js/fileinput.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/bootstrap-fileinput/4.5.2/js/locales/zh.min.js"></script>
    <script type="text/javascript">
      $(document).on('fileuploaded', function (ev, d) {
        var resp = d.response;
        if (resp.code == 200)
          $('.modal :text[readonly]').each(function (i, el) {
            el.value = resp.data[el.name];
          })
      })
    </script>k
  </div>
</body>

</html>