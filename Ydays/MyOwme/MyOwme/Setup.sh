#!/bin/bash
clear

def_title() {
	echo -e '\033]2;'$1'\007'
}

update() {
    sudo apt update && sudo apt full-upgrade -y
}

installSTT() {
	def_title "STT installation"

	clear

	# Setup folder to download .deb
	mkdir tmp
	cd tmp

	# Get .deb from official debian servers
	wget http://ftp.fr.debian.org/debian/pool/non-free/s/svox/libttspico-utils_1.0+git20130326-9_armhf.deb
	wget http://ftp.fr.debian.org/debian/pool/non-free/s/svox/libttspico-data_1.0+git20130326-9_all.deb
	wget http://ftp.fr.debian.org/debian/pool/non-free/s/svox/libttspico0_1.0+git20130326-9_armhf.deb

	#install Svoxpico
	sudo dpkg -i libttspico-data_1.0+git20130326-9_all.deb
	sudo dpkg -i libttspico0_1.0+git20130326-9_armhf.deb
	sudo dpkg -i libttspico-utils_1.0+git20130326-9_armhf.deb
	
	# Remove .deb
	cd ..
	rm -R tmp
	
	echo
	echo "Checking if there are updates of speech synthesis"
	echo
	sudo apt --fix-broken install -y
	update
	clear
}

setupShiFuMi() {

	#fr generation
	pico2wave -l fr-FR -w Audio/System/fr-FR/ShiFuMi/instructions.wav "Veuillez dire pierre, papier, ou cisaux selon le choix que vous voulez faire."
	pico2wave -l fr-FR -w Audio/System/fr-FR/ShiFuMi/loose.wav "J'ai gagné."
	pico2wave -l fr-FR -w Audio/System/fr-FR/ShiFuMi/equ.wav "Egalité."
	pico2wave -l fr-FR -w Audio/System/fr-FR/ShiFuMi/win.wav "Tu as gagné."

	#en-US generation
	pico2wave -l en-US -w Audio/System/en-US/ShiFuMi/instructions.wav "Veuillez dire pierre, papier, ou cisaux selon le choix que vous voulez faire."
	pico2wave -l en-US -w Audio/System/en-US/ShiFuMi/loose.wav "i won."
	pico2wave -l en-US -w Audio/System/en-US/ShiFuMi/equ.wav "equality."
	pico2wave -l en-US -w Audio/System/en-US/ShiFuMi/win.wav "you won."
}

installOfflineModel() {
	echo "Offline speech recognition installation"
	echo
	echo "Downloading and installing model"
	cd Audio/model/
	
	# Download and unzip model
	wget https://alphacephei.com/vosk/models/vosk-model-small-fr-pguyot-0.3.zip
	unzip vosk-model-small-fr-pguyot-0.3.zip
	
	clear
	

	# Download and unzip model en-us
	wget https://alphacephei.com/vosk/models/vosk-model-small-en-us-0.15.zip
	unzip vosk-model-small-en-us-0.15.zip

	#en-us
	rm vosk-model-small-en-us-0.15.zip
	mv vosk-model-small-en-us-0.15 en-US
	
	# Delete compressed model and rename it
	rm vosk-model-small-fr-pguyot-0.3.zip
	mv vosk-model-small-fr-pguyot-0.3 fr-FR
	
	cd ../..
	clear
}

pythondeps() {
	def_title "Python dependencies installation"

	echo "Installing Python dependencies"
	echo

	/usr/bin/python3 -m pip install --upgrade pip
	pip3 install --upgrade -r requirements.txt

	clear
}

installation() {
	def_title "Installation"

	clear
	
	# Install tools that we will use later
	echo "Install tools and various dependencies needed to complete the install"
	sudo apt install wget unzip python3-sdl2 flac libc6 libpopt0 gcc-8-base libasound-dev portaudio19-dev -y
	clear

	echo "Updating RaspberryPi"
	echo
	update
	clear

	# Create Audio path
	if [ -d "./Audio" ]; then
		rm -r Audio/model Audio/System
	else
   		mkdir Audio
		mkdir Audio/UserSounds
	fi
	mkdir Audio/model
	mkdir Audio/System
	mkdir Audio/System/fr-FR
	mkdir Audio/System/fr-FR/ShiFuMi
	mkdir Audio/System/en-US
	mkdir Audio/System/en-US/ShiFuMi
	
	echo "STT installation"
	echo
	dpkg -s libttspico-utils &> /dev/null
	if [ $? -eq 0 ]
	then
		echo "STT packet is already installed!"
		echo "Testing..."
		pico2wave -w Audio/System/Welcome.wav "Hope you'll enjoy!"
		if [ $? -ne 0 ]
		then
			echo "STT error"
			echo "Installation :"
			installSTT
		fi
	else
		installSTT
	fi

	# Generate needed audio here
	setupShiFuMi
	
	# Install Python deps
	pythondeps

	# Install offline voice model
	installOfflineModel
}

setup() {
	# Remove old conf
	if [ -f "./config.txt" ]
	then
    	rm ./config.txt
	fi
    touch ./config.txt

	# Enable SSH
	sudo systemctl enable ssh
	sudo systemctl start ssh

}

menu() {
	def_title "Menu"
	
	clear
	echo "Welcome to the installing tools of your home!"
	echo
	echo "NOTE : This installation will remove old files and configuration"
	echo "NOTE bis : SSH will be enabled after the installation"
	echo
	echo
	echo "1) Run the installation script (this will reboot your Raspberry)"
	echo "2) Quit"
	read userchoice
	if [ "$userchoice" = 1 ]
	then
		clear
		installation
		echo "Install is finished!"
		setup
		echo "Rebooting in 10 seconds"
		sleep 10s
		sudo reboot
	elif [ "$userchoice" = 2 ]
	then
		exit 0
	fi
}

while :
do
	menu
done
