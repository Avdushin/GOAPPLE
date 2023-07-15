# If PREFIX isn't provided, we check for $(DESTDIR)/usr/local and use that if it exists.
# Otherwice we fall back to using /usr.

LOCAL != test -d $(DESTDIR)/usr/local && echo -n "/local" || echo -n ""
LOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo "/local" || echo "")
PREFIX ?= /usr$(LOCAL)

default:
	# Run "sudo make install" to install the application.
	# Run "sudo make uninstall" to uninstall the application.

install:
	install -Dm00644 usr/local/share/applications/goapple-gui.desktop $(DESTDIR)$(PREFIX)/share/applications/goapple-gui.desktop
	install -Dm00755 usr/local/bin/goapple-gui $(DESTDIR)$(PREFIX)/bin/goapple-gui
	install -Dm00644 usr/local/share/pixmaps/goapple-gui.jpg $(DESTDIR)$(PREFIX)/share/pixmaps/goapple-gui.jpg
	sudo cp -rf src/ /usr/local/bin/
	sudo cp -rf src/ /usr/local/share/applications/
	# sudo cp -rf language.txt ~/.local/share/
	sh install.sh
uninstall:
	-rm $(DESTDIR)$(PREFIX)/share/applications/goapple-gui.desktop
	-rm $(DESTDIR)$(PREFIX)/bin/goapple-gui
	-rm $(DESTDIR)$(PREFIX)/share/pixmaps/goapple-gui.jpg
	-rm -rf /usr/local/bin/src/
	-rm -rf /usr/local/share/applications/src/
	-rm -rf language.txt ~/.local/share/