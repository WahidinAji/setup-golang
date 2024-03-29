# if you want to use air, you can use this script to install air and run it
# note: this script is only for development environment and assumes that you have postgres installed on your machine

export SENTRY_DSN="https://<your-sentry-dsn>"
export DB_NAME="<your-db-name>"
export DB_USER="<your-db-user>"
export DB_PASSWORD="<your-db-password>"
export DB_HOST="<your-db-host>"
export DB_PORT="<your-db-port>"
export ENVIRONMENT="development"
export RELEASE="0.0.1"
export PORT="<your-app-port>"
export PASSWORD_SECRET="<your-password-secret>"
export CORS_ALLOW_ORIGINS="http://localhost:${PORT}"
export CORS_ALLOW_METHODS="GET, POST, PUT, DELETE, OPTIONS"
export CORS_ALLOW_HEADERS="Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"


if [ ! -d "./bin" ]; then
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
fi

./bin/air