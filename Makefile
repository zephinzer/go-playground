build:
	cd basic-server && $(MAKE) init
	cd logger && $(MAKE) init
	cd server-routing && $(MAKE) init

create:
	cp -r .template ${PROJECT}