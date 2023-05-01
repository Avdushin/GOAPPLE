#!/bin/bash
echo -e "\e[0;92MainMenuanjaro script starting..."
sudo pacman -Syyyuuu
clear
echo -e "\e[0;92mInstalling programs..."
sudo pacman -Syu i3 i3blocks i3status i3lock polybar kitty krita fish ack vim bottom neofetch flameshot variety feh rofi discord python-pip steam telegram-desktop gcolor3 lxappearance picom flatpak xdg-desktop-portal-gtk awesome-terminal-fonts noto-fonts-emoji noto-fonts --noconfirm
echo -e "\e[0;92mCreating directories..."
sudo mkdir -p ~/.config/i3/
sudo mkdir -p ~/.local/share/Trash/files
echo -e "\e[0;92mInstalling programs..."
sudo tar -C $HOME -xzf src/packages/NoiseTorch_x64.tgz
sudo tar -xf src/packages/sublime_text.tar.xz -C /opt/ && gtk-update-icon-cache && flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
sudo cp -rf src/distros/Manjaro/i3/* ~/.config/i3/
sudo cp -r src/polybar/ ~/.config/  
sudo cp -r src/kitty/ ~/.config/
#sudo flatpak install flathub com.rafaelmardojai.Blanket -y 
sudo cp -r src/dots/.bashrc ~/ 
sudo cp -rf src/dots/rofi ~/.config/
sudo chsh -s /usr/bin/fish
chsh -s /usr/bin/fish
echo -e "\e[0;92mInstalling themes and fonts..."
sudo cp -r src/assets/walls/ ~/
git clone https://github.com/PIN3APPLEZZ/tella-icons
sudo cp -r tella-icons/Tela-blue tella-icons/Tela-blue-dark /usr/share/icons/
git clone https://github.com/PIN3APPLEZZ/themes-for-pin3apple
sudo cp -r themes-for-pin3apple/Solarized-Dark-Blue themes-for-pin3apple/Solarized-Dark-Cyan themes-for-pin3apple/Nordic-darker-standard-buttons /usr/share/themes/
sudo cp -r src/dots/.fonts ~/
sudo cp -r src/dots/git_token ~/
echo -e "\e[0;92mDONE!"
echo -e "\e[0;91m!\e[0;1;33mYour system will be \e[0;91mREBOOT!"
sleep 5
reboot 