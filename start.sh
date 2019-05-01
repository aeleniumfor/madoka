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