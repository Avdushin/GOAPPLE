### GOAPPLE GUI

![](src/assets/demo/app.dark.png)![](src/assets/demo/app.light.png)

GUI version of [GOAPPLE-CLI](https://github.com/Avdushin/GOAPPLE), written at Go Fyne.

### Dependencies

`make`

```
# to install make:

# Manjaro
sudo pacman -S make
# Solus
sudo eopkg it make
# Fedora
sudo dnf install make
```



### How to install

```
git clone https://github.com/Avdushin/goapple-gui
cd goapple-gui
sudo make install
# Click at the goapple-gui icon or 
goapple-gui
```

or u can use

1) 
  ```
   git clone https://github.com/Avdushin/goapple-gui
   cd goapple-gui
   ./install.sh
   ./goapple-gui
   ```
   
2) 
  ```
   sudo make install
   ./goapple-gui
   ```

### Usage

 * app's menu

![app.empty](src/assets/demo/app.dark.png)

* select  distro to install

  ![app.distro.select](src/assets/demo/app.distro.select.png)

#### menu

* File -> about

  Open Browser link about app

  ![apps.file.about](src/assets/demo/apps.file.about.png)



* Settings -> Themes ->

  You can select app's theme dark/light

![app.srttings.themes](src/assets/demo/app.srttings.themes.png)  ![app.light](src/assets/demo/app.light.png)

* App running ->

  When app is running - progress bar running! At this moment, application installing selected configuration

  ![app.running](src/assets/demo/app.running.png)

* Error (if you don't selected distro) ->

  ![app.select.errors](src/assets/demo/app.select.errors.png)



<p align="center">2022 © <a href="https://github.com/Avdushin" target="_blank">Avdushin</a></p>
