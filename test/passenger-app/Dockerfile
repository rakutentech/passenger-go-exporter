FROM ruby:2.7

ARG ARG_RAILS_VERSION=6.1.3
ARG ARG_PASSENGER_VERSION=6.0.5
ENV RAILS_VERSION $ARG_RAILS_VERSION
ENV PASSENGER_VERSION $ARG_PASSENGER_VERSION
#
RUN apt-get update \
  && apt-get install -y nodejs

WORKDIR /app
# Create rails new application.
RUN echo "gem: --no-document" > ~/.gemrc \
 && gem install rails -v ${RAILS_VERSION} || echo "ignore errors."
RUN rails new example || echo "ignore errors."

#
#  WORK FOR EXAMPLE DIRECTORY.
#
WORKDIR /app/example
RUN rails generate controller Welcome index

# Passenger configuration.
RUN echo "gem \"passenger\", \"${PASSENGER_VERSION}\", require: \"phusion_passenger/rack_handler\"" >> Gemfile \
 && bundle install

ENTRYPOINT [ "/usr/local/bin/bundle","exec","passenger","start" ]
