yum update
yum -y install go docker
mkdir -p /root/go/src/github.com/madoka
cd /
git clone https://github.com/containernetworking/plugins.git
cd plugins
./build_linux.sh
mkdir -p /opt/cni/bin/
mv bin/* /opt/cni/bin/
cd /
mkdir -p /etc/cni/net.d
echo '{"cniVersion": "0.2.0","name": "mynet","type": "bridge","bridge": "cni0","isGateway": true,"ipMasq": true,"ipam": {"type": "host-local","subnet": "10.22.0.0/16","routes": [{ "dst": "0.0.0.0/0" }]}}' > /etc/cni/net.d/mynet.conf

systemctl enable docker
systemctl start docker
yum -y install bridge-utils