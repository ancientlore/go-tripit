include $(GOROOT)/src/Make.inc

TARG=tripit
GOFILES=\
	tripit.go\
	http.go\
	webauth.go\
	oauth.go

include $(GOROOT)/src/Make.pkg

