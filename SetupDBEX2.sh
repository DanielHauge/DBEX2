
echo "Initializing AutoSetup"
echo "Installing essentials"
sudo apt-get update
sudo apt-get install -y wget
sudo apt-get install -y git-all
echo "Installing Docker! & Docker-compose"
wget -O - https://bit.ly/docker-install | bash
echo "--------------------------------------"
echo "--------------------------------------"
echo "-------------Cloning repo-------------"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 3s
git clone https://github.com/Games-of-Threads/DBEX2-DFH.git
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Building Image--------------"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 3s
sudo docker build -t dbex2:dfh $(pwd)/DBEX2-DFH/.
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Starting up MongoDB---------"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 3s
sudo docker run --rm -v $(pwd)/DBEX2-DFH/ImportData.sh:/ImportData.sh --publish=27017:27017 --name dbms -d mongo
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Initializing Data to DB-----"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 3s
sudo docker exec dbms sh -c 'chmod +x ImportData.sh'
sudo docker exec dbms sh -c './ImportData.sh'
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Starting Application--------"
echo "--------------------------------------"
echo "--------------------------------------"
sleep 3s
sudo docker run --rm -dit --link dbms:mongo --publish=9191:9191 --name GoApp dbex2:dfh
echo "--------------------------------------"
echo "--------------------------------------"
echo "----------Application Started---------"
echo "--------------------------------------"
echo "--------------------------------------"
cat <<"EOF"
You can now attach to the container with: docker attach GoApp
Or use the API endpoint by curling using browser on this machines IP. REMEMBER PORT 9191.
http://<ip>:9191/
a likely IP will be one of the following below
EOF
hostname -I

