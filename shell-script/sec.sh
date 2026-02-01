#!/bin/sh


NOW=$(date "+%Y/%m/%d %H:%M:%S")
FILES=$(ls -ltrh)
echo "$FILES"
echo "$NOW"

kernel=$(uname -s)
if [ $kernel = 'Linux' ]
then
  echo "You kernel is linux :) i am happy"
else
  echo "You kernel is not linux :("
fi

for file in $(ls -1);
do
  echo "my file is $file"
done
