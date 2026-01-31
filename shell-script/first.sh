#!/bin/bash

echo "Hello welcome to shell script!"
# dirname="test-dir-name"
# if [ -d $dirname ]
# then
#   echo "$dirname is exists"
# else
#   echo "$dirname not exists!"
# fi

# read input
# read -p "Enter a number: " num
# if (( $num == 12 )); then
#   echo "$num is valid"
# else 
#   echo "$num is in valid"
# fi
# if [ $num -gt 10 ]; then
#   echo "Number is greater than 10"
# fi

# read -r -p "Enter file name: " file
# if [ -f "$file" ]; then
#   echo "$file is exists"
# else 
#   echo "$file is not exists!!"
# fi

# read -r -p "Enter a number: " num
# if [ $num -gt 0 ]; then
#   echo "Negative"
# elif [ $num -lt 0 ]; then
#   echo "Positive"
# elif [ $num -eq 0 ]; then
#   echo "Zero"
# fi


# read -r -p "Enter dir name: " dirname
# if [ -d "$dirname" ]; then
#   echo "$dirname is exists"
# else echo "$dirname is not exists"
# fi 


read -r -p "Enter first number: " first_num
read -r -p "Enter secound number: " sec_num
if [ "$first_num" -gt "$sec_num" ]; then
  echo "$first_num is gt $sec_num"

elif [ "$first_num" -lt "$sec_num" ]; then
  echo "$first_num is lt $sec_num"
  
elif [ "$first_num" -eq $sec_num ]; then
  echo "$first_num is eq $sec_num"
fi