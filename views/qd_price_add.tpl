<form method="POST" action="/pdf/qd/price?action=save" class="form-horizontal">
    <div class="modal-header">
        <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
        <h4 class="modal-title">渠道价格添加</h4>
    </div>
    <div class="modal-body">
        <div class="form-group">
            <label class="col-sm-2 control-label">渠道号</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" name="qd" data-rule="{'maxlength':64}" required>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">价格</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" name="price" required>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">比例</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" name="coefficient" required>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">生效时间</label>
            <div class="col-sm-3">
                <input type="text" class="form-control" name="start" value="{{.Start}}" data-provide="datepicker"
                    data-date-language="zh-CN" data-date-format="yyyy-mm-dd"
                    data-date-autoclose="true" data-date-orientation="bottom" readonly>
            </div>
        </div>
    </div>
    <div class="modal-footer">
        <a class="btn btn-default" data-dismiss="modal">取消</a>
        <button type="submit" class="btn btn-danger">确定</button>
    </div>
</form>