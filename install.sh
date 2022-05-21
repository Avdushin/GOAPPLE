echo "INSTALLING COMPONENTS..."
sudo pacman -S st --noconfirm
clear
sudo dnf install st -y
clear
sudo eopkg it st -y
clear
sudo apt install st -y
clear
chmod +x *.desktop
sudo mkdir -p /etc/goapple-gui/
sudo cp -rf language.txt /etc/goapple-gui
sudo chmod 777 /etc/goapple-gui/language.txt
sudo chmod 777 /usr/local/bin/*
chmod +X src/distros/Fedora/*.sh
chmod +X src/distros/Manjaro/*.sh
chmod +X src/distros/Solus/*.sh
echo "DONE..."