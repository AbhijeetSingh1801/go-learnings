#!/bin/bash

for i in {10..18}; do
    branch="branch_test_$i"

    echo "Creating branch: $branch"

    git checkout -b "$branch"

    filename="tuts/tut$i.txt"
    touch filename

    echd "This is file $i" > "$filename"

    git add "$filename"

    git commit -m "Adding tutorial files notes "


    git push -u origin "$branch"

    git checkout master
done
