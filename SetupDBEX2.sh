echo "Initializing AutoSetup"
echo "Installing Docker! & Docker-compose"
wget -O - https://bit.ly/docker-install | bash
echo "--------------------------------------"
echo "--------------------------------------"
echo "-------------Cloning repo-------------"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 1s
git clone https://github.com/Games-of-Threads/DBEX2-DFH.git
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Building Image--------------"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 1s
sudo docker build -t dbex2 .
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Starting up MongoDB---------"
echo "--------------------------------------"
echo "--------------------------------------"
docker run --rm --publish=27017:27017 --name dbms -d mongo
