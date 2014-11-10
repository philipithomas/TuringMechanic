#!/bin/sh
set -e

bundle install
bundle exec jekyll build


# If we're in travis and on master branch - Deploy all the things!
# (we treat deployment as a test)

if [ "$TRAVIS" == "true" ] \
    && [ "$TRAVIS_BRANCH" == "master" ] \
    && [ "$TRAVIS_PULL_REQUEST" == "false" ] \
    && [ "$TRAVIS_REPO_SLUG" == "philipithomas/TuringMechanic" ] \
    && [ "$TRAVIS_SECURE_ENV_VARS" == "true" ]
then
      # Push the build using the s3_website.yml settings
      bundle exec s3_website push

      # Clear the Cloudflare cache
      curl https://www.cloudflare.com/api_json.html \
          -d 'a=fpurge_ts' \
          -d "tkn=$cloudflare_token" \
          -d "email=$cloudflare_email" \
          -d "z=$cloudflare_zone" \
          -d 'v=1'
fi
