<form method="POST" action="/pdf/publish?type={{ .data.Type }}" class="form-horizontal">
  <div class="modal-header">
    <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
    <h4 class="modal-title">版本列表</h4>
  </div>
  <div class="modal-body">
    <div class="form-group">
      <label class="col-sm-3 control-label">版本号</label>
      <div class="col-sm-3">
        <select class="form-control select2" name="version" data-ajax--url="/pdf/versions/list" data-ajax--cache="true">
        </select>
      </div>
    </div>
  </div>
  <div class="modal-footer">
    <a class="btn btn-default" data-dismiss="modal">取消</a>
    <button type="submit" class="btn bg-purple">确认</button>
  </div>
</form>