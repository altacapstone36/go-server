#!/bin/sh

CONFIG="
server:
\n SECRET: \"6568135b96cd0a3d6adcae0a475f0da473e8e2bee0d800b2ba5010186b37d950\"
\n PORT: \"8080\"
\nmysql:
\n DATABASE: \"hospital\"
\n USERNAME: \"root\"
\n PASSWORD: \"animelovers\"
\n HOST: \"localhost\"
\nmongo:
\n STRING: \"mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb\"
\n DATABASE: \"hospital\"
"

echo -e $CONFIG > config/config.yaml
make run
