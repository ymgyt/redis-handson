#!/bin/bash

VERSION=stable
SRC=redis-${VERSION}
ARCHIVE=${SRC}.tar.gz

CWD=$(cd $(dirname ${BASH_SOURCE[0]}) && pwd)

function install() {
  curl -O http://download.redis.io/releases/${ARCHIVE}
  tar xzvf ${ARCHIVE}
  cd ${SRC}
  make
  make test
  make install
}

function clean() {
  rm ${ARCHIVE} 
}


cd ${CWD}/../
install && clean
