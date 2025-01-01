#!/usr/bin/env bash

words=("hello" "world" "programming" "is" "fun" "and" "unpredictable")

#
## Generate Sentence out of random words
#
generate_sentence() {
  shuffled_words=($(printf "%s\n" "${words[@]}" | shuf))
  word_count=$((RANDOM % ${#shuffled_words[@]} + 2))
  sentence=$(printf "%s " "${shuffled_words[@]:0:word_count}")
  echo "${sentence% }"
}

#
## Main Function
#
main() {
  for i in {1..10}; do
    random_sentence=$(generate_sentence)
    curl -X POST -H "Content-Type: application/json" http://localhost:9999/api/v1/producer/message -d "{\"message\": \"$random_sentence\"}"
  done
}

#
## Run main
#
main
