FROM ruby:3.3

ARG ARG_RAILS_VERSION=7.0.8
ARG ARG_PASSENGER_VERSION=6.0.20
ENV RAILS_VERSION $ARG_RAILS_VERSION
ENV PASSENGER_VERSION $ARG_PASSENGER_VERSION

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
