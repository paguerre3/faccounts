# faccounts
f3 accounts challenge

### run challenge
build test image and execute tests based on docker compose dependencies defined
<pre>
docker-compose up --build --abort-on-container-exit
</pre>

clean up resources stored in volumes
<pre>
docker-compose down --volumes
</pre>

### troubleshoot commands
list network
<pre>
docker network ls  
</pre>
inspect default network created
<pre>
docker network inspect faccounts_default
</pre>

