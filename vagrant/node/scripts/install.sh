# update os
sudo apt update
sudo apt -y upgrade

swapoff -a

# install docker
sudo apt install -y docker.io
sudo systemctl enable docker
sudo systemctl start docker

# cretae folder /etc/apt/keyrings
mkdir /etc/apt/keyrings

# install kubernets
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.30/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.30/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt update
sudo apt-get install -y kubeadm kubelet kubectl
sudo apt-mark hold kubeadm kubelet kubectl
systemctl start kubelet