include $(GOROOT)/src/Make.$(GOARCH)

TARG=acme/oppai
GOFILES=\
	oppai.go\

include $(GOROOT)/src/Make.pkg
