package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-cas/cas"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
)

const mySessionCookieName = "_my_session_id"
const myListeningPort = "8080"

type myHandle struct {
	c *cas.Client
	g *gin.Engine
}

type mySession struct {
	id string
}

// Command line parameters
var casServerURL string
var ifName string

var myIpAddress string

// A very basic session mechanism
var mySessionCount uint32
var mySessionList = map[string]*mySession{}
var mySessionListMux sync.Mutex

func init() {
	flag.StringVar(&casServerURL, "cas", "", "Your CAS server URL")
	flag.StringVar(&ifName, "iface", "", "Your Ethernet interface")
}

func main() {
	flag.Parse()
	if casServerURL == "" {
		panic("Missing CAS server URL.")
	}
	if ifName == "" {
		panic("Missing Interface name.")
	}

	initMyIpAddress(ifName)
	if myIpAddress == "" {
		panic("Cannot detect my IP address")
	}

	u, _ := url.Parse(casServerURL)
	casClient := cas.NewClient(&cas.Options{
		URL: u,
	})

	ginRender := multitemplate.New()
	ginRender.AddFromFiles("index", "index.html")

	ginEngine := gin.Default()
	ginEngine.HTMLRender = ginRender

	ginEngine.Use(static.Serve("/", static.LocalFile(".", false)))

	ginEngine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", nil)
	})

	h := &myHandle{
		g: ginEngine,
		c: casClient,
	}
	srv := &http.Server{
		Addr:    ":" + myListeningPort,
		Handler: casClient.Handle(h),
	}

	_ = srv.ListenAndServe()
}

func initMyIpAddress(ifName string) {
	if iface, err := net.InterfaceByName(ifName); err == nil {
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					myIpAddress = ipnet.IP.String()
				}
			}
		}
	}
}

// ServeHTTP() processes CAS authentication and calls GIN server.
func (h *myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := findSession(r)
	if r.URL.Path == "/logout" {
		unregisterSession(sess)
		cas.RedirectToLogout(w, r)
		fmt.Printf("User logged out.\n")
		return
	}
	if sess != nil {
		// Ok. This guy is authenticated --> Call GIN
		h.g.ServeHTTP(w, r)
		return
	}

	if !cas.IsAuthenticated(r) {
		cas.RedirectToLogin(w, r)
		return
	}

	fmt.Printf("User authenticated: '%s'\n", cas.Username(r))
	for attr, value := range cas.Attributes(r) {
		fmt.Printf("- attribute : %-15s = %s\n", "'"+attr+"'", value[0])
	}

	sess = newSession()
	registerSession(sess)

	// Set a session cookie, so that findSession() will find this session, when
	// this client calls us.
	cookie := &http.Cookie{
		Name:     mySessionCookieName,
		Value:    sess.id,
		MaxAge:   0,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

	// At this point, the INDEX page is still not loaded.
	// Redirect the client to our /, so that it will have it, served by GIN server.
	http.Redirect(w, r, "http://"+myIpAddress+":"+myListeningPort+"/", http.StatusFound)
}

func newSession() *mySession {
	// Todo: session ID should be randomized.
	id := atomic.AddUint32(&mySessionCount, 1)
	sess := &mySession{id: fmt.Sprintf("%08x", id)}
	return sess
}

func registerSession(sess *mySession) {
	mySessionListMux.Lock()
	mySessionList[sess.id] = sess
	mySessionListMux.Unlock()
}

func unregisterSession(sess *mySession) {
	if sess == nil {
		return
	}
	mySessionListMux.Lock()
	delete(mySessionList, sess.id)
	mySessionListMux.Unlock()
}

func findSession(r *http.Request) *mySession {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == mySessionCookieName {
			id := cookie.Value

			mySessionListMux.Lock()
			if sess, ok := mySessionList[id]; ok {
				mySessionListMux.Unlock()
				return sess
			}
			mySessionListMux.Unlock()

			return nil
		}
	}

	return nil
}
