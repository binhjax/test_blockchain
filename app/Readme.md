# **************  Setup Environment ***************
1. Setup blockchain:
  #go run main.go  blockchain peers
  #go run main.go  accounts fillgas
2. Deploy contract & Init
  #go run main.go  contract deploy
  #go run main.go  contract init

3. Start
  + API Server
   #go run main.go  api
  + Process inqueue message  
  #go run main.go  inqueue
  + Process outqueue message  
  #go run main.go  outqueue
# **************  Main Feature ***************
- API server:
  + Manage blockchain network:
    - Add peer
    - Check transaction
    - Get Log
  + Manage contract
    - Load contract from config file
    - Save config to file
  + Manage account
    - Load from config file
    - Generate new account  
    - Store back to file
    - Auto add Gas
  + View report
- Workers:
  +  
- Listeners:
  +  Event Listeners
  +  Block listener
# **************  Structure ***************

# **************  Folder/Packages ***************
+ Account:
  Manage all wallet accounts
+ Admin:
  All admin action
+ Cache:
  Manage redis cache
+ Cmd: Manage commands
    + api: Api server
    + inqueueWorker: Process inqueu message
    + outqueuWorker: Process outqueue message
+ Config:
  Config type and loaddata from config file
+ Connection:
  Manage connection to geth
+ Contracts:
  Manage contract
+ Controllers:
  Business Logic Process for API queries. They will call functions from admin, account
+ docs:
  Docs for system. Using swagger
+ helper:
  Support functions
+ listener
  Listen block event and smart contract event from blockchain
+ middleware
  Middleware process for API server
+ Router: Route for API server
+ Tasks:
  Business Logic Process for Message Queue
    + Admin task will call function from contracts and admin to process task

+ Workers
  Process pub & sub with rabbit mq

# **************  API List  ***************
+ Generate doc:  swagger generate spec -o ./swagger.yaml --scan-models
+ View full: swagger serve -F=swagger swagger.yaml
+ Summary
    + 0. Authentication
    + 1. Admin  
      -> 1.1 Network
        + Connect peers
      ->  1.2. Account
        + From store
        + From file
        + Save to file
        + List token
      ->  1.3. Wallet
        + Load
        + List
        + New
        + Save
        + Fillgas
        + Balance
        + Send Eth between account

      ->  1.3. Contract
        + Deploy Contract
        + Initlization Contract

# **************  Queue Message Type  ***************


#Geth command

eth.getBlock("latest")
