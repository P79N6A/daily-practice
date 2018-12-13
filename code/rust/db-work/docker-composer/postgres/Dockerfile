# vim:set ft=dockerfile:
FROM daocloud.io/debian:8

RUN echo "deb http://mirrors.163.com/debian/ jessie main non-free contrib" > /etc/apt/sources.list  && \
echo "deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb http://mirrors.163.com/debian/ jessie-backports main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian/ jessie main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian/ jessie-updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian/ jessie-backports main non-free contrib" >> /etc/apt/sources.list && \
echo "deb http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib" >> /etc/apt/sources.list

# explicitly set user/group IDs
RUN groupadd -r postgres --gid=999 && useradd -r -g postgres --uid=999 postgres

ADD ./common/gosu-amd64 /usr/local/bin/gosu
ADD ./common/gosu-amd64.asc /usr/local/bin/gosu.asc

# grab gosu for easy step-down from root
ENV GOSU_VERSION 1.7
RUN set -x \
	&& apt-get update && apt-get install -y --no-install-recommends ca-certificates wget && rm -rf /var/lib/apt/lists/*
	
RUN export GNUPGHOME="$(mktemp -d)" \
	&& gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
	&& gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
	&& rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
  && chmod +x /usr/local/bin/gosu \
	&& gosu nobody true \
	&& apt-get purge -y --auto-remove ca-certificates wget

# make the "en_US.UTF-8" locale so postgres will be utf-8 enabled by default
RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
	&& localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

RUN mkdir /docker-entrypoint-initdb.d

RUN set -ex; \
# pub   4096R/ACCC4CF8 2011-10-13 [expires: 2019-07-02]
#       Key fingerprint = B97B 0AFC AA1A 47F0 44F2  44A0 7FCC 7D46 ACCC 4CF8
# uid                  PostgreSQL Debian Repository
	key='B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8'; \
	export GNUPGHOME="$(mktemp -d)"; \
	gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key"; \
	gpg --export "$key" > /etc/apt/trusted.gpg.d/postgres.gpg; \
	rm -r "$GNUPGHOME"; \
	apt-key list

ENV PG_MAJOR 9.6
ENV PG_VERSION 9.6.6-0+deb9u1~bpo8+1

RUN echo 'deb http://apt.postgresql.org/pub/repos/apt/ jessie-pgdg main' $PG_MAJOR > /etc/apt/sources.list.d/pgdg.list

RUN apt-get update \
	&& apt-get install -y postgresql-common \
	&& sed -ri 's/#(create_main_cluster) .*$/\1 = false/' /etc/postgresql-common/createcluster.conf
	
RUN apt-get install -y \
		postgresql-$PG_MAJOR=$PG_VERSION \
		postgresql-contrib-$PG_MAJOR=$PG_VERSION \
	&& rm -rf /var/lib/apt/lists/*

# make the sample config easier to munge (and "correct by default")
RUN mv -v /usr/share/postgresql/$PG_MAJOR/postgresql.conf.sample /usr/share/postgresql/ \
	&& ln -sv ../postgresql.conf.sample /usr/share/postgresql/$PG_MAJOR/ \
	&& sed -ri "s!^#?(listen_addresses)\s*=\s*\S+.*!\1 = '*'!" /usr/share/postgresql/postgresql.conf.sample

RUN mkdir -p /var/run/postgresql && chown -R postgres /var/run/postgresql

ENV PATH /usr/lib/postgresql/$PG_MAJOR/bin:$PATH
ENV PGDATA /var/lib/postgresql/data
VOLUME /var/lib/postgresql/data

COPY docker-entrypoint.sh /
RUN chmod 755 /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 5432
CMD ["postgres"]
