Who's In?
=========

- Client-side Cordova app in /codiqa-app
- Server-side Google App Engine + Go in /appengine

To compile client side app (OS X):

	cd codiqa-app
	cordova/build

APK is in bin/whosin-debug.apk

To push to App Engine:

	cd appengine
	appcfg.py update .
	