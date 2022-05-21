#!/bin/bash
echo -e "\e[0;92mFedora script starting..."
sudo dnf update -y
clear
echo -e "\e[0;92mInstalling programs..."
sudo dnf install \
  https://download1.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm -y && sudo dnf install \
  https://download1.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
 i3 i3-gaps i3blocks i3status i3lock polybar kitty krita fish git ack vim bottom neofetch flameshot variety feh rofi discord steam telegram-desktop picom gcolor3 lxappearance flatpak xdg-desktop-portal-gtk fontawesome-fonts -y --allowerasing --skip-broken
echo -e "\e[0;92mCreating directories..."
sudo mkdir -p ~/.appz
sudo mkdir -p ~/.config/i3/
sudo mkdir -p ~/.local/share/Trash/files
echo -e "\e[0;92mInstalling programs..."
sudo tar -C $HOME -xzf src/packages/NoiseTorch_x64.tgz
sudo tar -xf src/packages/sublime_text.tar.xz -C ~/.appz && gtk-update-icon-cache && flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
sudo cp src/distros/Fedora/config ~/.config/i3/
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