include $(GOROOT)/src/Make.inc

TARG=tripit
GOFILES=\
	tripit.go\
	http.go\
	webauth.go\
	oauth.go\
	activityobjectptrvector.go\
	airobjectptrvector.go\
	carobjectptrvector.go\
	closenessmatchvector.go\
	cruiseobjectptrvector.go\
	directionsobjectptrvector.go\
	errorvector.go\
	groupvector.go\
	inviteevector.go\
	lodgingobjectptrvector.go\
	mapobjectptrvector.go\
	noteobjectptrvector.go\
	pointsprogramvector.go\
	profileemailaddressvector.go\
	profilevector.go\
	railobjectptrvector.go\
	restaurantobjectptrvector.go\
	transportobjectptrvector.go\
	tripcrsremarkvector.go\
	tripptrvector.go\
	warningvector.go\
	weatherobjectvector.go\
	imageptrvector.go\
	travelerptrvector.go\
	airsegmentptrvector.go\
	railsegmentptrvector.go\
	transportsegmentptrvector.go\
	cruisesegmentptrvector.go\
	pointsprogramactivityvector.go\
	pointsprogramexpirationvector.go

include $(GOROOT)/src/Make.pkg

