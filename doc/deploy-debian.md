
# Debian Jessie Deploy guide (assumes amd64)

Get dependancies:

    # apt update
    # apt install \
    git \
    build-essential \
    wget \
    nginx \
    libssl-dev \
    libboost-chrono-dev \
    libboost-date-time-dev \
    libboost-filesystem-dev \
    libboost-program-options-dev \
    libboost-system-dev \
    libboost-thread-dev \
    fail2ban

Build i2pd:

    # mkdir -p /usr/local/src/i2pd
    # cd /usr/local/i2pd
    # wget https://github.com/PurpleI2P/i2pd/archive/2.8.0.tar.gz -O i2pd-2.8.0.tar.gz
    # tar -xzvf i2pd-2.8.0.tar.gz
    # cd i2pd-2.8.0
    # make 

Set Up i2pd:

    # cp /usr/local/src/i2pd/i2pd-2.8.0/i2pd /usr/bin/
    # useradd i2pd
    # mkdir /home/i2pd/
    # chown i2pd:i2pd /home/i2pd/
    # mkdir -p /etc/i2pd
    # cp /usr/local/src/inproxy/contrib/i2pd/i2pd.conf /etc/i2pd/
    # chown i2pd:i2pd -R /etc/i2pd/
    # cp /usr/local/src/inproxy/contrib/systemd/i2pd.service /etc/systemd/system/
    # systemctl enable i2pd
    # systemctl start i2pd

Get go 1.6:

    # cd /usr/local
    # wget https://storage.googleapis.com/golang/go1.6.3.linux-amd64.tar.gz -O go-1.6.3.tar.gz
    # tar -xzvf go-1.6.3.tar.gz
    # cat >> /etc/bash.bashrc <<END
    export GOROOT=/usr/local/go
    export PATH=$GOROOT/bin
    END

Get proxy server:

    # cd /usr/local/src
    # git clone https://github.com/majestrate/inproxy
    # cd inproxy

Build proxy server:

    # make

Set up proxy server:

    // set this to a domain name you control
    # export INPROXY_DOMAIN=site.tld
    
    # mkdir -p /opt/inproxy/bin
    # cp /usr/local/src/inproxy/inproxy /opt/inproxy/bin/
    # /usr/local/src/inproxy/configure i2p.$INPROXY_DOMAIN > /opt/inproxy/inproxy.ini
    # cp /usr/local/src/inproxy/contrib/systemd/inproxy.service /etc/systemd/system/
    # systemctl enable inproxy
    # systemctl start inproxy


Set Up nginx:
    
    # cp /usr/local/src/inproxy/contrib/nginx/inproxy /etc/nginx/sites-available/
    # ln -s /etc/nginx/sites-available/inproxy /etc/nginx/sites-enabled/
    # systemctl reload nginx

Update Dns Record: (Given you are using `site.tld`)

    *.i2p.site.tld. IN CNAME site.tld.


If everything works then try: http://em763732l4b7b7zhaolctpt6wewwr7zw3nsxfchr6qmceizzmgpa.b32.i2p.site.tld/

Replace `site.tld` with your domain
