package handler

import (
	"encoding/json"
	"fmt"
	"gocms/util"
	"net/http"
	"sdbackend/domain"
	"sdbackend/utils"
	"strconv"
	"strings"
	"time"

	"gocms/model"

	"github.com/gorilla/mux"
)

func QDList(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.Admin
		ok   bool
		qds  []string
		err  error
	)
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
		return
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
		return
	} else if user, ok = cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
		return
	}

	if user.Group.ID == 1 || user.Group.Name == "data_admin" {
		qds, err = model.AllQDs()
		if err != nil {
			http.NotFound(w, r)
			return
		}
	} else {
		qds, err = model.AdmindQDs(user.ID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	data := make(map[string]interface{})
	selects := []select2{
		select2{"all", "all"},
	}
	for i := range qds {
		selects = append(selects, select2{ID: qds[i], Name: qds[i]})
	}
	data["results"] = selects
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)

}

func QDSettleDay(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.Admin
		ok   bool
		qds  []string
		err  error
	)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
		return
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
		return
	} else if user, ok = cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
		return
	}

	tpl := "settle_day_all.tpl"
	if user.Group.ID == 1 || user.Group.Name == "data_admin" {
		qds, err = model.AllQDs()
		if err != nil {
			http.NotFound(w, r)
			return
		}
	} else {
		qds, err = model.AdmindQDs(user.ID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tpl = "settle_day_qd.tpl"
	}

	data := make(map[string]interface{})
	if qds != nil {
		if qd := strings.TrimSpace(r.Form.Get("qd")); len(qd) > 0 && qd != "all" {
			check := false
			for i := range qds {
				if (qds)[i] == qd {
					qds = []string{qd}
					check = true
					break
				}
			}
			if !check {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}

		var viewList []model.QDInstallRuns
		if nums, err := model.TotalInstallRunsByQD(qds); err == nil && nums > 0 {
			p := util.NewPaginator(r, nums)
			if qdStats, err := model.InstallRunsByQD(qds, p.PerPageNums, p.Offset()); err == nil {
				for i := range qdStats {
					coefficient := qdStats[i].Coefficient
					if user.Group.ID == 1 || user.Group.Name == "data_admin" {
						coefficient = 100
					}
					qdStats[i].Total = (qdStats[i].InstallEnd * qdStats[i].Coefficient / 100) * qdStats[i].Price
					qdStats[i].InstallStart = qdStats[i].InstallStart * coefficient / 100
					qdStats[i].UninstallStart = qdStats[i].UninstallStart * coefficient / 100
					qdStats[i].UninstallEnd = qdStats[i].UninstallEnd * coefficient / 100
					qdStats[i].MFShow = qdStats[i].MFShow * coefficient / 100
					qdStats[i].ServerRun = qdStats[i].ServerRun * coefficient / 100
					qdStats[i].InstallEnd = qdStats[i].InstallEnd * coefficient / 100
					viewList = append(viewList, qdStats[i])
				}
				data["list"] = viewList
			}
			data["page"] = p
		}
	}
	rLayout(w, r, tpl, data)
}

func QDSettleMonth(w http.ResponseWriter, r *http.Request) {
	var (
		user *model.Admin
		ok   bool
		qds  []string
		err  error
	)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if session, err := store.Get(r, sessName); err != nil {
		http.NotFound(w, r)
		return
	} else if cookie, exist := session.Values["user"]; !exist {
		http.NotFound(w, r)
		return
	} else if user, ok = cookie.(*model.Admin); !ok {
		http.NotFound(w, r)
		return
	}

	tpl := "settle_month_all.tpl"
	if user.Group.ID == 1 || user.Group.Name == "data_admin" {
		qds, err = model.AllQDs()
		if err != nil {
			http.NotFound(w, r)
			return
		}
	} else {
		qds, err = model.AdmindQDs(user.ID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		tpl = "settle_month_qd.tpl"
	}

	data := make(map[string]interface{})
	if qds != nil {
		if qd := strings.TrimSpace(r.Form.Get("qd")); len(qd) > 0 && qd != "all" {
			check := false
			for i := range qds {
				if (qds)[i] == qd {
					qds = []string{qd}
					check = true
					break
				}
			}
			if !check {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}

		if nums, err := model.TotalMonthSettleByQD(qds); err == nil && nums > 0 {
			p := util.NewPaginator(r, nums)
			if qdStats, err := model.MonthSettleByQD(qds, p.PerPageNums, p.Offset()); err == nil {
				data["list"] = qdStats
			}
			data["page"] = p
		}
	}
	rLayout(w, r, tpl, data)
}

func QDRetentions(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}

	var (
		qd        = strings.TrimSpace(r.Form.Get("qd"))
		from      = r.Form.Get("from")
		to        = r.Form.Get("to")
		roundStrs = r.Form["r"]
		rounds    []int
	)
	for i := range roundStrs {
		rd, err := strconv.Atoi(roundStrs[i])
		if err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
		rounds = append(rounds, rd)
	}
	if rounds == nil {
		rounds = []int{1, 3, 7, 30}
	}
	if qd == "all" {
		qd = ""
	}

	data := make(map[string]interface{})
	if nums, err := model.TotalQDRetention(rounds, from, to, qd); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if retentions, err := model.GetQDRetentions(p.PerPageNums, p.Offset(), rounds, from, to, qd); err == nil {
			data["list"] = retentions
		}
		data["page"] = p
	}
	if result, err := model.GetAvgPDFRetentions(); err == nil {
		data["avg"] = *result
	}
	data["rounds"] = rounds
	rLayout(w, r, "qd_retentions.tpl", data)
}

func QDPrices(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.QDPricesTotal("", ""); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if prices, err := model.QDPrices(p.PerPageNums, p.Offset(), "", ""); err == nil {
			data["list"] = prices
		}
		data["page"] = p
	}
	rLayout(w, r, "qd_prices.tpl", data)
}

func QDPriceAdd(w http.ResponseWriter, r *http.Request) {
	now := time.Now().In(utils.Location())
	data := &domain.GroupCoefficient{
		Start: now.AddDate(0, 0, 1).Format("2006-01-02"),
	}
	if err := t.ExecuteTemplate(w, "qd_price_add.tpl", data); err != nil {
		w.Write([]byte(err.Error()))
	}

}

func QDPriceEdit(w http.ResponseWriter, r *http.Request) {
	qd := mux.Vars(r)["qd"]
	ins, err := model.GetQDPriceByQD(qd)
	if err != nil || ins == nil {
		rLayout(w, r, "error.tpl", nil)
		return
	}

	fmt.Println(*ins)
	if err := t.ExecuteTemplate(w, "qd_price_edit.tpl", ins); err != nil {
		w.Write([]byte(err.Error()))
	}
}

func SaveQDPrice(w http.ResponseWriter, r *http.Request) {
	var ins domain.GroupCoefficient
	if err := r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := util.ParseForm(r.Form, &ins); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	if ins.Coefficient == 0 || ins.QD == "" {
		jFailed(w, http.StatusBadRequest, "invalid param")
		return
	}
	if err := model.SaveQDPrice(&ins); err != nil {
		jFailed(w, http.StatusInternalServerError, "保存错误，请联系管理员!")
		return
	}
	jSuccess(w, nil)
}
