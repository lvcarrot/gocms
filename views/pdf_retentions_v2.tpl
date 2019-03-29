<!DOCTYPE html>
<html>

<head>
  {{template "header" .node.Name}}
  <!-- Latest compiled and minified CSS -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-select@1.13.8/dist/css/bootstrap-select.min.css">
  <style>
    .select2 {
      min-width: 120px;
    }
  </style>
</head>

<body class="hold-transition skin-blue sidebar-mini">

  <div class="wrapper">
    {{template "navbar" .}}
    <div class="content-wrapper">
      {{template "title" .}}
      <section class="content">
        <!-- Small boxes (Stat box) -->
        <div class="row">
          <div class="col-lg-3 col-xs-6">
            <!-- small box -->
            <div class="small-box bg-aqua">
              <div class="inner">
                <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow}}/{{rate .data.avg.RServerRun}}</sup></h3>
                <p>Avg次日留存</p>
              </div>
              <div class="icon">
                <i class="fa fa-bar-chart" aria-hidden="true"></i>
              </div>
            </div>
          </div>
          <!-- ./col -->
          <div class="col-lg-3 col-xs-8">
            <!-- small box -->
            <div class="small-box bg-green">
              <div class="inner">
                <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow3}}/{{rate .data.avg.RServerRun3}}</sup></h3>
                <p>Avg三日留存</p>
              </div>
              <div class="icon">
                <i class="fa fa-area-chart" aria-hidden="true"></i>
              </div>
            </div>
          </div>
          <!-- ./col -->
          <div class="col-lg-3 col-xs-6">
            <!-- small box -->
            <div class="small-box bg-yellow">
              <div class="inner">
                <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow7}}/{{rate .data.avg.RServerRun7}}</sup></h3>
                <p>Avg七日留存</p>
              </div>
              <div class="icon">
                <i class="fa fa-bar-chart" aria-hidden="true"></i>
              </div>
            </div>
          </div>
          <!-- ./col -->
          <div class="col-lg-3 col-xs-6">
            <!-- small box -->
            <div class="small-box bg-red">
              <div class="inner">
                <h3><sup style="font-size: 20px">{{rate .data.avg.RMFShow30}}/{{rate .data.avg.RServerRun30}}</sup></h3>
                <p>Avg三十日留存</p>
              </div>
              <div class="icon">
                <i class="fa fa-line-chart" aria-hidden="true"></i>
              </div>
            </div>
          </div>
        </div>
        <div class="box">
          <div class="box-header with-border">
            <div class="box-title">PDF留存率</div>
            <div class="box-tools visible-lg">
              <form class="form-inline">
                <label>开始/结束日期</label>
                <div class="form-group">
                  <div class="input-group input-group-sm input-daterange" data-provide="datepicker"
                    data-date-language="zh-CN" data-date-format="yyyy-mm-dd" data-date-end-date="0d"
                    data-date-autoclose="true" data-date-orientation="bottom">
                    <input type="text" class="form-control" placeholder="请选择开始" name="from" value="{{.form.Get "from"}}"
                      readonly>
                    <span class="input-group-addon"><i class="fa fa-calendar"></i></span>
                    <input type="text" class="form-control" placeholder="请选择结束" name="to" value="{{.form.Get "to"}}"
                      readonly>
                  </div>
                </div>
                <label>留存类型</label>
                <div class="form-group">
                  <select class="selectpicker" name="r" multiple title="请选择留存类型" data-actions-box="true">
                    <option value="0" {{if roundSelected 0 .data.rounds }}selected{{end}}>当日</option>
                    <option value="1" {{if roundSelected 1 .data.rounds}}selected{{end}}>次日</option>
                    <option value="3" {{if roundSelected 3 .data.rounds}}selected{{end}}>3日</option>
                    <option value="4" {{if roundSelected 4 .data.rounds}}selected{{end}}>4日</option>
                    <option value="5" {{if roundSelected 6 .data.rounds}}selected{{end}}>5日</option>
                    <option value="6" {{if roundSelected 6 .data.rounds}}selected{{end}}>6日</option>
                    <option value="7" {{if roundSelected 7 .data.rounds}}selected{{end}}>7日</option>
                    <option value="8" {{if roundSelected 8 .data.rounds}}selected{{end}}>8日</option>
                    <option value="9" {{if roundSelected 9 .data.rounds}}selected{{end}}>9日</option>
                    <option value="10" {{if roundSelected 10 .data.rounds}}selected{{end}}>10日</option>
                    <option value="11" {{if roundSelected 11 .data.rounds}}selected{{end}}>11日</option>
                    <option value="12" {{if roundSelected 12 .data.rounds}}selected{{end}}>12日</option>
                    <option value="13" {{if roundSelected 13 .data.rounds}}selected{{end}}>13日</option>
                    <option value="14" {{if roundSelected 14 .data.rounds}}selected{{end}}>14日</option>
                    <option value="15" {{if roundSelected 15 .data.rounds}}selected{{end}}>15日</option>
                    <option value="16" {{if roundSelected 16 .data.rounds}}selected{{end}}>16日</option>
                    <option value="17" {{if roundSelected 17 .data.rounds}}selected{{end}}>17日</option>
                    <option value="18" {{if roundSelected 18 .data.rounds}}selected{{end}}>18日</option>
                    <option value="19" {{if roundSelected 19 .data.rounds}}selected{{end}}>19日</option>
                    <option value="20" {{if roundSelected 20 .data.rounds}}selected{{end}}>20日</option>
                    <option value="21" {{if roundSelected 21 .data.rounds}}selected{{end}}>21日</option>
                    <option value="22" {{if roundSelected 22 .data.rounds}}selected{{end}}>22日</option>
                    <option value="23" {{if roundSelected 23 .data.rounds}}selected{{end}}>23日</option>
                    <option value="24" {{if roundSelected 24 .data.rounds}}selected{{end}}>24日</option>
                    <option value="25" {{if roundSelected 25 .data.rounds}}selected{{end}}>25日</option>
                    <option value="26" {{if roundSelected 26 .data.rounds}}selected{{end}}>26日</option>
                    <option value="27" {{if roundSelected 27 .data.rounds}}selected{{end}}>27日</option>
                    <option value="28" {{if roundSelected 28 .data.rounds}}selected{{end}}>28日</option>
                    <option value="29" {{if roundSelected 29 .data.rounds}}selected{{end}}>29日</option>
                    <option value="30" {{if roundSelected 30 .data.rounds}}selected{{end}}>30日</option>
                  </select>
                </div>
                <button type="submit" class="btn bg-purple btn-sm" title="筛选">筛选 <i class="fa fa-filter"></i></button>
              </form>
            </div>
          </div>
          {{if .data.list}}
          <div class="box-body table-responsive">
            <table class="table table-bordered">
              <tbody>
                <tr>
                  <th>日期</th>
                  <th>留存类型</th>
                  <th>新增</th>
                  <th>留存人数</th>
                  <th>留存率</th>
                </tr>
                {{range .data.list}}
                <tr>
                  <td>{{.Date}}</td>
                  <td>{{retentionType .Round}}</td>
                  <td>{{.Install}}</td>
                  <td>{{.MFShow}}/{{.ServerRun}}</td>
                  <td>{{rate .MFShowRate}}/{{rate .ServerRunRate}}</td>
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

<!-- Latest compiled and minified JavaScript -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap-select@1.13.8/dist/js/bootstrap-select.min.js"></script>

<!-- (Optional) Latest compiled and minified JavaScript translation files -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap-select@1.13.8/dist/js/i18n/defaults-zh_CN.min.js"></script>

</html>
