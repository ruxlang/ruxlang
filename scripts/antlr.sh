#!/bin/sh
FILENAME=antlr-4.11.1-complete.jar
alias antlr='java -jar $HOME/bin/$FILENAME'
mkdir -p $HOME/bin
if [ ! -f $HOME/bin/$FILENAME ]; then
    wget http://www.antlr.org/download/$FILENAME -O $HOME/bin/$FILENAME
fi

echo $PWD
antlr $@