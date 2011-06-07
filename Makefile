include $(GOROOT)/src/Make.inc

TARG=tripit
GOFILES=\
	tripit.go\
	http.go\
	webauth.go\
	oauth.go\
	vector/activityobjectptrvector.go\
	vector/airobjectptrvector.go\
	vector/carobjectptrvector.go\
	vector/closenessmatchvector.go\
	vector/cruiseobjectptrvector.go\
	vector/directionsobjectptrvector.go\
	vector/errorvector.go\
	vector/groupvector.go\
	vector/inviteevector.go\
	vector/lodgingobjectptrvector.go\
	vector/mapobjectptrvector.go\
	vector/noteobjectptrvector.go\
	vector/pointsprogramvector.go\
	vector/profileemailaddressvector.go\
	vector/profilevector.go\
	vector/railobjectptrvector.go\
	vector/restaurantobjectptrvector.go\
	vector/transportobjectptrvector.go\
	vector/tripcrsremarkvector.go\
	vector/tripptrvector.go\
	vector/warningvector.go\
	vector/weatherobjectvector.go\
	vector/imageptrvector.go\
	vector/travelerptrvector.go\
	vector/airsegmentptrvector.go\
	vector/railsegmentptrvector.go\
	vector/transportsegmentptrvector.go\
	vector/cruisesegmentptrvector.go\
	vector/pointsprogramactivityvector.go\
	vector/pointsprogramexpirationvector.go

include $(GOROOT)/src/Make.pkg

