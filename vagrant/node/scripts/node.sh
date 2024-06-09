# deploy node
sudo ufw disables
# sudo hostnamectl set-hostname master-node
sudo hostnamectl set-hostname worker-node-$1

# install networks plugin
sudo kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl get pods --all-namespaces

# kubeadm join 172.16.133.166:6443 --token 92asxq.4omfwmf9sh7gchu9 \
        # --discovery-token-ca-cert-hash sha256:30aa5c04c6dd3fa3695c7915e30ac2f4309d2ba47dfd52b718fb148d3a31b08a 