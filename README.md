## Hyperledger Fabric Sample Application

This application demonstrates the creation and transfer of assets between actors leveraging Hyperledger Fabric. We will set up the minimum number of nodes required to develop the chaincode (smartcontracts).There will be a single peer and a single orgranization in this application to show how the network operates. 

**Please ensure you have the following components on your computer before proceeding:**
*cURL
*Node.js
*npm package manager
*Go language
*Docker (https://docs.docker.com)
*Docker Compose

**Lets get started!**

Please open your local terminal and clone the following repository

$ git clone https://github.com/hyperledger/education.git

Next navigate into your the asset-app.

Please make sure Docker is running at this time. Should you have used it prior, please run the code below to clear your containers.


Now lets start the Hyperledger Fabric network. Please enter the following command:

$ ./startFabric.sh

If you continue to watch your terminal, you will begin to see the network start up and great an initial "genesis block" which we will be the initial data we work with in this app. 

Please now run the following command:

$ npm install

Now lets spin up the application. Enter the following commands in the following order. You will also be directed in your terminal to enter this source code.

$ node registerAdmin.js

$ node registerUser.js

$ node server.js

The app is live and should be fully functioning. Please go to localhost:8000 in your web browser as your terminal will also direct you. There an application will appear. 

In this application you will be able to query all assets currently registered in the genesis block as well as create additional asset you wish to track. Furthermore, have the ability to change owners(holders) of a particular asset. 

To ensure your application is live and looks as it should, run the index.html file locally.

The itention of this application is for educational purposes and utilizes source code from the Hyplerledger Fabric Project (https://github.com/hyperledger/fabric) as well as the edEx Linux Foundation Hyplerledger course(https://github.com/hyperledger/education.git).
