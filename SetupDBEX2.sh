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
sudo docker build -t dbex2 $(pwd)/DBEX2-DFH/.
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Starting up MongoDB---------"
echo "--------------------------------------"
echo "--------------------------------------"
sudo docker run --rm -v $(pwd)/DBEX2-DFH/ImportData.sh:ImportData.sh --publish=27017:27017 --name dbms -d mongo
sudo docker exec dbms sh -c 'sudo ./ImportData.sh'
