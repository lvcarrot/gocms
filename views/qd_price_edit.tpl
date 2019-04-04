<form method="POST" action="/pdf/qd/price?action=save" class="form-horizontal">
    <div class="modal-header">
        <a class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></a>
        <h4 class="modal-title">渠道价格修改</h4>
    </div>
    <div class="modal-body">
        <div class="form-group">
            <div class="col-sm-3">
                <input type="hidden" class="form-control" name="qd" value="{{.QD}}" data-rule="{'maxlength':64}" required>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">价格</label>
            <div class="col-sm-2">
                <input type="text" class="form-control" name="price" value="{{.Price}}" required>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">比例</label>
            <div class="col-sm-2">
                <input type="text" class="form-control" name="coefficient" value="{{.Coefficient}}" required>
            </div>
        </div>
    </div>
    <div class="modal-footer">
        <a class="btn btn-default" data-dismiss="modal">取消</a>
        <button type="submit" class="btn btn-danger">确定</button>
    </div>
</form>