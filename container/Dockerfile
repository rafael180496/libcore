#Ambiente go oracle 
FROM golang
LABEL maintainer="rafael180496@gmail.com"
#Utilidades
RUN apt-get update 
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends apt-utils

#dependencia de la imagen
RUN apt-get install vim -y && \
    apt-get install curl -y && \
    apt-get install wget -y && \
    apt-get install gzip -y && \
    apt-get install unzip -y && \
    apt-get install zip -y && \
    apt-get install p7zip-full -y && \
    apt-get install inotify-tools -y && \
    apt-get install software-properties-common -y && \
    apt-get install libaio1 -y && \
    apt-get install alien -y 
#intalacion de rpm
RUN apt-get install rpm -y
#intalar git
RUN apt-get install git -y
#oracle cli 12.1
RUN mkdir /usr/auxprogram && \
    cd /usr/auxprogram && \
    git clone https://github.com/rafael180496/dockerglash.git

RUN cd  /usr/auxprogram/dockerglash && \
    7z e jdkclient.7z.001 
RUN cd  /usr/auxprogram/dockerglash && \
    rm -R jdkclient.7z.0* 
RUN cd  /usr/auxprogram/dockerglash && \
    alien -i *.rpm

RUN cd  /usr/auxprogram/dockerglash && \
    rm *.rpm && \
    rm -R Comprimer && \
    rm README.md && \
    mkdir -p "/usr/lib/oracle/12.1/client64/network/admin"

RUN rm -R  /usr/auxprogram
# creacion de variable oci
RUN mkdir -p "/usr/local/share/pkgconfig"
COPY configssh/oci8.pc /usr/local/share/pkgconfig/oci8.pc 
#variables de entorno

ENV PKG_CONFIG_PATH /usr/local/share/pkgconfig
ENV ORACLE_HOME /usr/lib/oracle/12.1/client64
ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:$ORACLE_HOME/lib:$ORACLE_HOME
ENV OCI_HOME /usr/lib/oracle/12.1/client64/lib
ENV OCI_LIB_DIR $ORACLE_HOME/lib
ENV OCI_INC_DIR /usr/include/oracle/12.1/client64
ENV OCI_INCLUDE_DIR /usr/include/oracle/12.1/client64
ENV OCI_VERSION 12

ENV C_INCLUDE_PATH /usr/include:/usr/include/oracle/12.1/client64:/usr/local/include
ENV TNS_ADMIN $ORACLE_HOME/network/admin
ENV TNSADMIN $ORACLE_HOME/network/admin
ENV GOBIN $GOPATH/bin
ENV GCCGO gccgo
ENV CC gcc
ENV CXX g++
ENV CGO_ENABLED 1
ENV PKG_CONFIG pkg-config
ENV GO111MODULE on
ENV GIT_TERMINAL_PROMPT 1
ENV GOPROXY direct
ENV GOSUMDB off
ENV  PATH $PATH:$ORACLE_HOME/bin
#depuracion de la imagen
RUN  apt-get remove --purge unzip -y && \
     apt-get remove --purge p7zip-full -y && \    
     apt-get remove --purge alien -y && \
     apt-get remove --purge rpm -y 

RUN  apt-get autoremove -y

#configuracion ssh

RUN apt-get install openssh-server -y
RUN apt-get install sudo -y

USER root
ENV  SSH_MASTER_USER master
ENV  SSH_MASTER_PASS abc123

COPY configssh/ssh_config /etc/ssh/ssh_config
COPY configssh/sshd_config /etc/ssh/sshd_config
COPY configssh/user.sh /usr/local/bin/user.sh
RUN chmod +x /usr/local/bin/user.sh
RUN /usr/local/bin/user.sh
 
COPY configssh/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD tail -f /dev/null