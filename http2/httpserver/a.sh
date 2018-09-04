#!/bin/bash

PEER=""
ORG=""
CNAME="
while getopts "h?peer:org:f:c" opt; do 
	case "$opt" in
		h|\?)
		echo "print help"
		exit 0
		;;
		peer) PEER=$OPTARG
		;;
		org)  ORG=$OPTARG
		;;
		c)    CNAME=$OPTARG 
		;;
		f)    CNAME+$OPTARG
		;;
    	esac
done 

echo "peer: $PEER"

