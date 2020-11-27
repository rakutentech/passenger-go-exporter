FROM ruby:2.7

#
RUN apt-get update \
  && apt-get install -y nodejs

WORKDIR /app
# Create rails new application.
RUN gem install rails -v "6.0.3.4" || echo "ignore errors."
RUN rails new example || echo "ignore errors."

#
#  WORK FOR EXAMPLE DIRECTORY.
#
WORKDIR /app/example
RUN rails generate controller Welcome index

# Passenger configuration.
RUN echo 'gem "passenger", "6.0.5", require: "phusion_passenger/rack_handler"' >> Gemfile \
  && bundle install

ENTRYPOINT [ "/usr/local/bin/bundle","exec","passenger","start" ]
