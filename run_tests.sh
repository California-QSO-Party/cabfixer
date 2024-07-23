#!/bin/bash

./cabfixer a.log
diff a.xcbr a_answer.xcbr
if  [[ $? -eq 0 ]]; then
  echo "test passed"
else
  echo "test failed"
fi