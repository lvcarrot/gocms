package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sdbackend/domain"
	"sync"
	"time"

	"gocms/model"

	"github.com/Tomasen/realip"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/time/rate"
)

const (
	defaultMaxMemory = 32 << 20  // 32 MB
	sessName         = "X-GoCMS" // Session 名称
	dateFormate      = "2006-01-02"
)

type select2 struct {
	ID   string `json:"id"`
	Name string `json:"text"`
}

var (
	t           *template.Template
	build       = "0"
	md5Regexp   = regexp.MustCompile("[a-fA-F0-9]{32}$")
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$")
	store       = sessions.NewFilesystemStore(os.TempDir(), securecookie.GenerateRandomKey(32))
)

func aLog(r *http.Request, format string, a ...interface{}) error {
	m := &model.AdminLog{
		Path:   r.URL.String(),
		UA:     r.UserAgent(),
		IP:     realip.FromRequest(r),
		Commit: fmt.Sprintf(format, a...),
	}
	if session, err := store.Get(r, sessName); err != nil {
		return err
	} else if cookie, exist := session.Values["user"]; !exist {
		return http.ErrNoCookie
	} else if user, ok := cookie.(*model.Admin); ok {
		m.AdminID = user.ID
	}
	return m.Create()
}

func jSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Code int         `json:"code"`
		Data interface{} `json:"data,omitempty"`
	}{Code: http.StatusOK, Data: data})
}

func jFailed(w http.ResponseWriter, code int, format string, a ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Code int    `json:"code"`
		Msg  string `json:"msg,omitempty"`
	}{Code: code, Msg: fmt.Sprintf(format, a...)})
}

// rLayout 渲染模板
func rLayout(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	if session, err := store.Get(r, sessName); err != nil {
		Error(w, http.StatusBadRequest, "页面错误 %v", err)
	} else if s := mux.CurrentRoute(r); s == nil {
		Error(w, http.StatusBadRequest, "页面错误")
	} else if tpl, err := s.GetPathTemplate(); err != nil {
		Error(w, http.StatusNotFound, "页面错误 %v", err)
	} else if err = t.ExecuteTemplate(w, name, map[string]interface{}{
		"menu": model.GetNodes(),
		"node": model.GetNodeByPath(tpl),
		"user": session.Values["user"],
		"form": r.Form,
		"data": data,
	}); err != nil {
		w.Write([]byte(err.Error()))
	}
}

// Error 错误页面
func Error(w http.ResponseWriter, code int, format string, a ...interface{}) {
	w.WriteHeader(code)
	t.ExecuteTemplate(w, "error.tpl", map[string]interface{}{
		"code": code, "text": http.StatusText(code),
		"msg": fmt.Sprintf(format, a...),
	})
}

// Check 检查用户登录
func Check(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session, err := store.Get(r, sessName); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if cookie, exist := session.Values["user"]; !exist {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if user, ok := cookie.(*model.Admin); !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if token, exist := tokenMap[user.ID]; exist && token != session.ID {
			session.Options.MaxAge = -1
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusFound)
		} else if !user.Status && r.URL.Path != "/profile" {
			http.Redirect(w, r, "/profile", http.StatusFound)
		} else if c := mux.CurrentRoute(r); c == nil {
			Error(w, http.StatusNotFound, "页面错误")
		} else if tpl, err := c.GetPathTemplate(); err != nil {
			Error(w, http.StatusNotFound, "页面错误 %v", err)
		} else if user.Access(tpl) {
			h.ServeHTTP(w, r)
		} else {
			Error(w, http.StatusForbidden, "无权访问 %s", tpl)
		}
	})
}

// Limit 请求限制
func Limit(b int, f func(http.ResponseWriter, *http.Request)) http.Handler {
	var (
		bucket = make(map[string]*rate.Limiter)
		lock   sync.Mutex
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip    = realip.FromRequest(r)
			l     *rate.Limiter
			exist bool
		)
		lock.Lock()
		if l, exist = bucket[ip]; !exist {
			l = rate.NewLimiter(rate.Limit(1), b)
			bucket[ip] = l
		}
		lock.Unlock()
		if l.Allow() {
			f(w, r)
			return
		}
		Error(w, http.StatusTooManyRequests, "请求太频繁")
	})
}

// WriteLog 日志打印
func WriteLog(w io.Writer, p handlers.LogFormatterParams) {
	var u string
	if session, err := store.Get(p.Request, sessName); err == nil {
		if cookie, exist := session.Values["user"]; exist {
			u = fmt.Sprint(cookie)
		}
	}
	fmt.Fprintf(w, "%s %s %d %s %d %s (%s) %s\n", p.TimeStamp.Format("2006/01/02 15:04:05"),
		p.Request.Method, p.StatusCode, p.URL.RequestURI(), p.Size,
		realip.FromRequest(p.Request), time.Now().Sub(p.TimeStamp), u)
}

// Start 初始化控制层
func Start(path string) error {
	// 注册自定义函数
	funcMap := template.FuncMap{
		"date": func(t *time.Time) string {
			if t == nil {
				return "无"
			}
			return t.In(time.Local).Format(time.RFC3339)
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"url": func(r url.Values, a ...string) template.URL {
			if len(a) <= 0 {
				return template.URL(r.Encode())
			}
			u := make(url.Values)
			for _, n := range a {
				if v := r.Get(n); len(v) > 0 {
					u.Add(n, v)
				}
			}
			return template.URL(u.Encode())
		},
		"version": func() string {
			return fmt.Sprintf("1.0.%s", build)
		},
		"rate": func(r int64) string {
			if r == 0 {
				return "-"
			}
			return fmt.Sprintf("%.2f", float64(r)/100)
		},
		"price": func(r, w int64) string {
			if r == 0 {
				return "0.00"
			}
			if w == 4 {
				return fmt.Sprintf("%.4f", float64(r)/10000)
			}
			return fmt.Sprintf("%.2f", float64(r)/100)
		},
		"retention": func(mfshow, serverRun int64) string {
			var (
				m = fmt.Sprintf("%0.2f", float64(mfshow)/float64(100))
				s = fmt.Sprintf("%0.2f", float64(serverRun)/float64(100))
			)
			if mfshow == 0 {
				m = "-"
			}
			if serverRun == 0 {
				s = "-"
			}

			return fmt.Sprintf("%s/%s", m, s)
		},
		"updateType": func(updateType domain.UpdateTypeEnum) string {
			switch updateType {
			case domain.UpdateTypeNormal:
				return "普通更新"
			case domain.UpdateTypeSilent:
				return "静默升级"
			case domain.UpdateTypePop:
				return "自动弹窗"
			}
			return "未知"
		},
		"versionType": func(versionType domain.VersionTypeEnum) string {
			switch versionType {
			case domain.VersionBeat:
				return "Beat"
			case domain.VersionRelease:
				return "Release"
			}
			return "未知"
		},
		"boolNote": func(f bool) string {
			if f {
				return "是"
			} else {
				return "否"
			}
		},
	}
	// 文件监控
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %v", err)
	}
	go func() {
		pattern := filepath.Join(path, "views", "*.tpl")
		for {
			select {
			case e := <-watcher.Events:
				log.Printf("load %s: %d", filepath.Base(e.Name), e.Op)
				if t, err = template.New("").Funcs(funcMap).ParseGlob(pattern); err != nil {
					log.Printf("parse %s failed: %v", e.Name, err)
				}
			case err := <-watcher.Errors:
				log.Printf("Watcher error: %v", err) // No need to exit here
			}
		}
	}()
	watcher.Events <- fsnotify.Event{}
	return watcher.Add(filepath.Join(path, "views"))
}
