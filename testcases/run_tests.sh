#!/bin/bash
for i in *.raw; do
  ../cabfixer $i
  filename=$(basename $i .raw)
  output_file=$filename.xcbr
  echo $output_file
  answer_file=${filename}_answer.xcbr
  echo $answer_file
  dos2unix $answer_file
  diff --unified $output_file $answer_file
  if  [[ $? -eq 0 ]]; then
    echo "test passed"
  else
    echo "test failed"
  fi

done
