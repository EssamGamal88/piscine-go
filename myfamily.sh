#!/bin/bash
curl -s "https://01.nextera.education/assets/superhero/all.json" | \
jq --argjson id "$HERO_ID" -r '.[] | select(.id == $id) | .connections.relatives | gsub("\\n"; "\\n")'