{{ $v := .data.version }}
<form method="POST" action="/pdf/versions/{{$v.Version}}?action=save" class="form-horizontal">
  <div class="modal-header">
    <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
    <h4 class="modal-title">版本管理</h4>
  </div>
  <div class="modal-body">
    <div class="form-group">
      <label class="col-sm-2 control-label">版本号</label>
      <div class="col-sm-3">
        {{if eq $v.Version "new"}}
        <input type="text" class="form-control" name="version" data-rule="{'maxlength':64}" required>
        {{else}}
        <p class="form-control">{{$v.Version}}</p>
        <input type="hidden" class="form-control" name="version" value="{{$v.Version}}">
        {{end}}
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">更新类型</label>
      <div class="col-sm-6 icheck">
        <label class="radio-inline">
          <input type="radio" name="update_type" value="1" {{if le $v.UpdateType 1}} checked {{end}}>
          普通更新
        </label>
        <label class="radio-inline">
          <input type="radio" name="update_type" value="2" {{if eq $v.UpdateType 2}} checked {{end}}>
          自动弹窗
        </label>
        <label class="radio-inline">
          <input type="radio" name="update_type" value="3" {{if eq $v.UpdateType 3}} checked {{end}}>
          静默更新
        </label>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">版本类型</label>
      <div class="col-sm-6 icheck">
        <label class="radio-inline">
          <input type="radio" name="version_type" value="1" {{if le $v.VersionType 1}} checked {{end}}>
          Release
        </label>
        <label class="radio-inline">
          <input type="radio" name="version_type" value="0" {{if eq $v.VersionType 2}} checked {{end}}>
          Beat
        </label>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">发布说明</label>
      <div class="col-sm-8">
        <textarea class="form-control" rows="3" name="release_note">{{ $v.ReleaseNote }}</textarea>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">发布时间</label>
      <div class="col-sm-4">
        <input type="text" class="form-control" name="release_date" value="{{date $v.ReleaseDate}}" required disabled>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">包大小</label>
      <div class="col-sm-2">
        <input type="text" class="form-control" name="pkg_size" value="{{$v.PkgSize}}" required readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">MD5</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="md5" value="{{$v.MD5}}" required readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">URL</label>
      <div class="col-sm-8">
        <input type="text" class="form-control" name="pkg_url" value="{{$v.PkgURL}}" required readonly>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">上传</label>
      <div class="col-sm-8">
        <input type="file" data-upload-url="/upload" data-show-preview="false" data-show-remove="false" data-language="zh">
      </div>
    </div>
  </div>
  <div class="modal-footer">
    <a class="btn btn-default" data-dismiss="modal">取消</a>
    <button type="submit" class="btn btn-danger">确定</button>
  </div>
</form>