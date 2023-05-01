#!/bin/bash
echo -e "\e[0;92MainMenuanjaro script starting...\e[0m"
# Updating
sudo pacman -Syyyuuu
clear
echo -e "\e[0;92mInstalling programs...\e[0m"
# installing packages
sudo pacman -Syu i3 i3blocks i3status i3lock polybar kitty krita fish ack vim bottom neofetch flameshot variety feh rofi discord python-pip steam telegram-desktop gcolor3 lxappearance picom flatpak xdg-desktop-portal-gtk awesome-terminal-fonts noto-fonts-emoji noto-fonts --noconfirm
# Creating directories
echo -e "\e[0;92mCreating directories...\e[0m"
sudo mkdir -p ~/.config/i3/
sudo mkdir -p ~/.local/share/Trash/files
# Unpacking apps (NoiseTorch & Sublime Text)
echo -e "\e[0;92mInstalling programs...\e[0m"
sudo tar -C $HOME -xzf src/packages/NoiseTorch_x64.tgz
sudo tar -xf src/packages/sublime_text.tar.xz -C /opt/ && gtk-update-icon-cache && flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
# I3 config
sudo cp -rf src/distros/Manjaro/i3/* ~/.config/i3/
# Polybar config
sudo cp -r src/polybar/ ~/.config/
# Kitty terminal config
sudo cp -r src/kitty/ ~/.config/
#sudo flatpak install flathub com.rafaelmardojai.Blanket -y 
sudo cp -r src/dots/.bashrc ~/ 
# Rofi config
sudo cp -rf src/dots/rofi ~/.config/
# Setting fish-shell by default
sudo chsh -s /usr/bin/fish
chsh -s /usr/bin/fish
# Themes and styles
echo -e "\e[0;92mInstalling themes and fonts...\e[0m"
# set wallpaper
sudo cp -r src/assets/walls/ ~/
# Set icons
git clone https://github.com/PIN3APPLEZZ/tella-icons
sudo cp -r tella-icons/Tela-blue tella-icons/Tela-blue-dark /usr/share/icons/
# Set Themes
git clone https://github.com/PIN3APPLEZZ/themes-for-pin3apple
sudo cp -r themes-for-pin3apple/Solarized-Dark-Blue themes-for-pin3apple/Solarized-Dark-Cyan themes-for-pin3apple/Nordic-darker-standard-buttons /usr/share/themes/
# Set fonts
sudo cp -r src/dots/.fonts ~/
# Get git_token file example
sudo cp -r src/dots/git_token ~/
echo -e "\e[0;92mDONE!\e[0m"
echo -e "\e[0;91m!\e[0;1;33mYour system will be \e[0;91mREBOOT!\e[0m"
sleep 5
reboot 