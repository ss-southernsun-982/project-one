# deploy node
# sudo swapoff –a
# swapoff -a
sudo ufw disables
sudo hostnamectl set-hostname master-node
# sudo hostnamectl set-hostname w1

# init master node
sudo kubeadm init --pod-network-cidr=100.0.0.0/16
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

export KUBECONFIG=/etc/kubernetes/admin.conf

# install networks plugin
sudo kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl get pods --all-namespaces