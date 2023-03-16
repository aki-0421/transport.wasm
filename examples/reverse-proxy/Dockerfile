FROM golang:1.20-bullseye

RUN apt update && \
    apt install -y cmake ninja-build software-properties-common

# download tinygo
WORKDIR /tmp
RUN git clone --recursive https://github.com/tinygo-org/tinygo.git

# checkout release
WORKDIR /tmp/tinygo
RUN git checkout release && \
    git submodule update --init

# install llvm-15
RUN curl -O https://apt.llvm.org/llvm-snapshot.gpg.key &&  apt-key add llvm-snapshot.gpg.key
RUN add-apt-repository -y "deb http://apt.llvm.org/bullseye/ llvm-toolchain-bullseye-15 main"
RUN apt update && \
    apt install -y clang-15 llvm-15-dev lld-15 libclang-15-dev

# build tinygo
RUN make binaryen && \
    make wasi-libc && \
    go install

# install nodejs
RUN curl -fsSL https://deb.nodesource.com/setup_current.x | bash -
RUN apt install -y nodejs

# install  wrangler
RUN npm install -g wrangler

WORKDIR /usr/src/app
