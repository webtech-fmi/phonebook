FROM postgres:13.1

RUN apt-get update \
      && apt-get install -y --no-install-recommends \
           postgresql-13-postgis-3 \
           postgresql-13-postgis-3-scripts \
      && rm -rf /var/lib/apt/lists/*