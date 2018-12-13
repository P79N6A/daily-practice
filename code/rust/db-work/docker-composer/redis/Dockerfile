FROM daocloud.io/debian:8

RUN echo "deb http://mirrors.163.com/debian/ jessie main non-free contrib" > /etc/apt/sources.list  && \
echo "deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb http://mirrors.163.com/debian/ jessie-backports main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian/ jessie main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian/ jessie-updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian/ jessie-backports main non-free contrib" >> /etc/apt/sources.list && \
echo "deb http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.163.com/debian-security/ jessie/updates main non-free contrib" >> /etc/apt/sources.list

# add our user and group first to make sure their IDs get assigned consistently, regardless of whatever dependencies get added
RUN groupadd -r redis && useradd -r -g redis redis

ADD ./common/gosu-amd64 /usr/local/bin/gosu
ADD ./common/gosu-amd64.asc /usr/local/bin/gosu.asc

# grab gosu for easy step-down from root
ENV GOSU_VERSION 1.7
RUN set -x \
	&& apt-get update && apt-get install -y --no-install-recommends ca-certificates wget && rm -rf /var/lib/apt/lists/* \
	&& export GNUPGHOME="$(mktemp -d)" \
	&& gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B42F6819007F00F88E364FD4036A9C25BF357DD4 \
	&& gpg --batch --verify /usr/local/bin/gosu.asc /usr/local/bin/gosu \
	&& rm -r "$GNUPGHOME" /usr/local/bin/gosu.asc \
	&& chmod +x /usr/local/bin/gosu \
	&& gosu nobody true \
	&& apt-get purge -y --auto-remove ca-certificates wget

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates \
		wget \
	&& rm -rf /var/lib/apt/lists/*

ENV REDIS_VERSION 3.0.7
ENV REDIS_DOWNLOAD_URL http://download.redis.io/releases/redis-3.0.7.tar.gz
ENV REDIS_DOWNLOAD_SHA1 e56b4b7e033ae8dbf311f9191cf6fdf3ae974d1c

# for redis-sentinel see: http://redis.io/topics/sentinel
RUN buildDeps='gcc libc6-dev make' \
	&& set -x \
	&& apt-get update && apt-get install -y $buildDeps --no-install-recommends \
	&& rm -rf /var/lib/apt/lists/* \
	&& wget -O redis.tar.gz "$REDIS_DOWNLOAD_URL" \
	&& echo "$REDIS_DOWNLOAD_SHA1 *redis.tar.gz" | sha1sum -c - \
	&& mkdir -p /usr/src/redis \
	&& tar -xzf redis.tar.gz -C /usr/src/redis --strip-components=1 \
	&& rm redis.tar.gz \
	&& make -C /usr/src/redis \
	&& make -C /usr/src/redis install \
	&& rm -r /usr/src/redis \
	&& apt-get purge -y --auto-remove $buildDeps

RUN mkdir /data && chown redis:redis /data
VOLUME /data
WORKDIR /data

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod 755 /usr/local/bin/docker-entrypoint.sh

COPY redis.conf /usr/local/etc/redis/

RUN ln -s usr/local/bin/docker-entrypoint.sh /entrypoint.sh # backwards compat
ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 6379
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
