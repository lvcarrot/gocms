<form method="POST" action="/pdf/versions/{{.Version}}?action=save" class="form-horizontal">
  <div class="modal-header">
    <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
    <h4 class="modal-title">版本管理</h4>
  </div>
  <div class="modal-body">
    <div class="form-group">
      <label class="col-sm-2 control-label">版本号</label>
      <div class="col-sm-3">
        <input type="text" class="form-control" name="version" data-rule="{'maxlength':64}" required>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">更新类型</label>
      <div class="col-sm-6 icheck">
        <label class="radio-inline">
          <input type="radio" name="update_type" value="1" checked>
          普通更新
        </label>
        <label class="radio-inline">
          <input type="radio" name="update_type" value="2">
          自动弹窗
        </label>
        <label class="radio-inline">
          <input type="radio" name="update_type" value="2">
          静默更新
        </label>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">更新类型</label>
      <div class="col-sm-6 icheck">
        <label class="radio-inline">
          <input type="radio" name="version_type" value="1" checked>
          Release
        </label>
        <label class="radio-inline">
          <input type="radio" name="version_type" value="2">
          Beat
        </label>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">发布说明</label>
      <div class="col-sm-8">
        <textarea class="form-control" rows="3" name="release_note"></textarea>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">发布时间</label>
      <div class="col-sm-3">
        <input type="text" class="form-control" name="release_date" required disabled>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">包大小</label>
      <div class="col-sm-2">
        <input type="text" class="form-control" name="pkg_size" required disabled>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">MD5</label>
      <div class="col-sm-6">
        <input type="text" class="form-control" name="md5" required disabled>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">URL</label>
      <div class="col-sm-8">
        <input type="text" class="form-control" name="pkg_url" required disabled>
      </div>
    </div>
    <div class="form-group">
      <label class="col-sm-2 control-label">上传</label>
      <div class="col-sm-8">
        <input type="file" class="file" data-upload-url="/upload" data-show-preview="false" data-show-remove="false"
          data-language="zh">
      </div>
    </div>
  </div>
  <div class="modal-footer">
    <a class="btn btn-default" data-dismiss="modal">取消</a>
    <button type="submit" class="btn btn-danger">确定</button>
  </div>
</form>

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
</script>