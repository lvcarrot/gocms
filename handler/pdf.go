package handler

import (
	"encoding/json"
	"gocms/model"
	"gocms/util"
	"log"
	"net/http"
	"sdbackend/domain"
	"time"

	"github.com/gorilla/mux"
)

var local *time.Location

type Select struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func init() {
	var err error
	local, err = time.LoadLocation("Asia/Chongqing")
	if err != nil {
		panic(err)
	}
}

func PDFInstallRuns(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalPDFInstallRuns(); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if installRuns, err := model.GetPDFInstallRuns(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = installRuns
		}
		if grandTotal, err := model.GetPDFInstallRunsByDate("total"); err == nil {
			data["total"] = grandTotal
		}
		data["page"] = p
	}
	rLayout(w, r, "pdf_install_runs.tpl", data)
}

func PDFRentions(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalRetentions(); err == nil && nums > 0 {
		p := util.NewPaginator(r, nums)
		if retentions, err := model.GetPDFRentions(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = retentions
		}
		data["page"] = p
	}
	if result, err := model.GetAvgPDFRetentions(); err == nil {
		data["avg"] = *result
	}
	rLayout(w, r, "pdf_retentions.tpl", data)
}

func MFShowVersions(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	result, err := model.GetMFShowVersions()
	if err == nil {
		data["list"] = result
	}
	rLayout(w, r, "pdf_mfshow_versions.tpl", data)

}

func Feedbacks(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalFeedbacks(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if feedbacks, err := model.GetFeedbacks(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = feedbacks
		}
		data["page"] = p
	}
	rLayout(w, r, "pdf_feedbacks.tpl", data)

}

func UninstallOpts(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalUninstallOpts(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if uninstallOpts, err := model.GetUninstallOpts(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = uninstallOpts
		}
		data["page"] = p
	}
	if results, err := model.GetUninstallResults(); err == nil {
		data["results"] = results
	}
	rLayout(w, r, "pdf_uninstall_opts.tpl", data)
}

func BundleInstall(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalBundleInstalls(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if list, err := model.GetBundleInstalls(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "bundle_installation.tpl", data)
}

func MiniNewsStats(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalMiniNewsStats(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if list, err := model.GetMiniNewsStats(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "mininews_stats.tpl", data)
}

func Crashs(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalPDFInstallRuns(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if crashs, err := model.GetPDFInstallRuns(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = crashs
		}
		data["page"] = p
	}

	rLayout(w, r, "pdf_crashs.tpl", data)
}

func CrashsDetail(w http.ResponseWriter, r *http.Request) {
	start, err := time.ParseInLocation("2006-01-02", r.URL.Query().Get("date"), local)
	if err != nil {
		start = time.Now().AddDate(0, 0, 1)
	}
	end := start.AddDate(0, 0, 1)

	data := make(map[string]interface{})
	if nums, err := model.GetCrashsTotal(&start, &end); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if rates, err := model.GetCrashVersioRate(&start, &end); err == nil {
			data["crash_rate"] = rates
		}
		if crashs, err := model.GetCrashsByDay(p.PerPageNums, p.Offset(), &start, &end); err == nil {
			data["crash_list"] = crashs
		}
		data["page"] = p
	}

	rLayout(w, r, "pdf_crashs_detail.tpl", data)
}

func KitTipStats(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.GetTotalKitTip(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if list, err := model.GetKitTipStats(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = list
		}
		data["page"] = p
	}
	rLayout(w, r, "kit_tip_stats.tpl", data)
}

func GetPDFVersions(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if nums, err := model.TotalPDFVersions(); err == nil && nums > 0 {
		p := util.NewPaginator(r, int64(nums))
		if versions, err := model.GetPDFVersions(p.PerPageNums, p.Offset()); err == nil {
			data["list"] = versions
			api, web := model.GetPDFReleaseVersion()
			log.Println("api:", api, " web:", web)
			for i := range versions {
				if versions[i].Version.Version == api {
					data["api"] = &versions[i]
					versions[i].ReleaseOnApi = true
				} else {
					versions[i].ReleaseOnApi = false
				}
				if versions[i].Version.Version == web {
					data["web"] = versions[i]
					versions[i].ReleaseOnWeb = true
				} else {
					versions[i].ReleaseOnWeb = false
				}
			}
		}
		data["page"] = p
	}
	rLayout(w, r, "pdf_version.tpl", data)
}

func GetPDFVersion(w http.ResponseWriter, r *http.Request) {
	var (
		version *domain.Version
		err     error
		v       = mux.Vars(r)["version"]
		data    = make(map[string]interface{})
	)
	if v == "new" {
		now := time.Now().In(local)
		version = &domain.Version{
			ReleaseDate: &now,
		}
	} else {
		version, err = model.GetPDFVersion(v)
		if err != nil {
			rLayout(w, r, "error.tpl", nil)
			return
		}
	}
	data["version"] = version
	rLayout(w, r, "version_edit.tpl", data)
}

func GetPDFVersionList(w http.ResponseWriter, r *http.Request) {
	versions, err := model.GetPDFVersions(16, 0)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	var selects []select2
	for i := range versions {
		selects = append(selects, select2{versions[i].Version.Version, versions[i].Version.Version})
	}
	data, _ := json.Marshal(struct {
		Results []select2 `json:"results"`
	}{selects})
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ModalPDFPublish(w http.ResponseWriter, r *http.Request) {
	var (
		typ  = r.URL.Query().Get("type")
		data = map[string]interface{}{
			"Type": typ,
		}
	)
	rLayout(w, r, "version_publish.tpl", data)
}

func PDFPublish(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	var (
		version = r.Form.Get("version")
		typ     = r.Form.Get("type")
	)

	if version == "" {
		jFailed(w, http.StatusBadRequest, "empty version num")
		return
	}
	if typ == "WebSite" {
		if err := model.PublishPDFWebsite(version); err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
	} else if typ == "Api" {
		if err := model.PublishPDFApi(version); err != nil {
			jFailed(w, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		jFailed(w, http.StatusBadRequest, "invalid type")
		return
	}
	jSuccess(w, nil)
}

func SavePDFVersion(w http.ResponseWriter, r *http.Request) {
	var version domain.Version
	if err := util.ParseForm(r.Form, &version); err != nil {
		jFailed(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Printf("%#v", version)
	if version.Version == "" ||
		version.PkgURL == "" ||
		version.PkgSize == 0 ||
		version.MD5 == "" ||
		version.ReleaseDate == nil ||
		version.NeedUpdate == 0 ||
		version.VersionType == 0 {
		jFailed(w, http.StatusBadRequest, "invalid param")
		return
	}
	typ := mux.Vars(r)["version"]
	if typ == "new" {
		if _, err := model.GetPDFVersion(version.Version); err == nil {
			jFailed(w, http.StatusBadRequest, "当前版本号已存在")
			return
		}
	} else {
		v, err := model.GetPDFVersion(version.Version)
		if err != nil {
			jFailed(w, http.StatusBadRequest, "版本不存在")
			return
		}
		version.ID = v.ID
	}
	model.FlushVesionCache()
	jSuccess(w, nil)
}
