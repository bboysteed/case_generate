#!/bin/bash
DIR=`dirname $0`

echo "dir: " $DIR
echo "0:" $0
echo "1:" $1
echo "2:" $2


case $2 in
p1) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/2.out $DIR/../../tests/whitebox/2.in && exit 0;;
p2) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/1.out $DIR/../../tests/whitebox/1.in && exit 0;;
n1) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/9.out $DIR/../../tests/whitebox/9.in && exit 0;;
n2) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/8.out $DIR/../../tests/whitebox/8.in && exit 0;;
n3) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/3.out $DIR/../../tests/whitebox/3.in && exit 0;;
n4) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/7.out $DIR/../../tests/whitebox/7.in && exit 0;;
n5) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/6.out $DIR/../../tests/whitebox/6.in && exit 0;;
n6) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/5.out $DIR/../../tests/whitebox/5.in && exit 0;;
n7) $DIR/../../../bin/genprog_tests.py --program grade $1 $DIR/../../tests/whitebox/4.out $DIR/../../tests/whitebox/4.in && exit 0;;
esac
exit 1
