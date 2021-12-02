#!/bin/bash -e

echo "Branch is '$1'"

BRANCH_NAME=$(echo  $1 | awk -F "/" '{print $NF}')
if [ "$BRANCH_NAME" == "main" ];
then
  DAYS=$(ls -d1 day*)
else
  DAYS=$BRANCH_NAME
fi

for day in $DAYS;
do
  echo Processing $day
  (cd $day; go run main.go)
done
