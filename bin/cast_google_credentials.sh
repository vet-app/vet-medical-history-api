#!/bin/sh

# Generate the google credentials json file at runtime
echo "------> Generating google-credentials.json from Heroku config var"
echo $GOOGLE_CREDENTIALS > google-credentials.json
exec "$@"