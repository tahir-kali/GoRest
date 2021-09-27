# GoRest

A simple restful Api server in Golang + firebase using jwt authentications 
to get serviceaccount.json please contact the developer in telegram 
QuickStart:
1. Replace the serviceaccount.json with real data 
2. run main.go 
test user 

UserID: tahir@emailroar.com
Password: 123Alif@Academy

EndPoints

1. Post localhost:2021/createuser 
	body {
	"email":"",
	"first":"",
	"last":"",
	"password":"",
	}
2. Get localhost:2021/		Generates a JWT token and sends it as header X-Digest: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//All E Wallat operations must have header X-Digest setup after this point
3. Post localhost:2021/topup/{email}
	body{
	"user":"", //email of user
	"type":"topup",
	"balance":int,
	"date":""
	}
4.Get localhost:2021/balance/{email}
5.Get localhost:2021/transactions/{email}
6.Get localhost:2021/validate/{email}

Project Structure:

│   go.mod
│   go.sum
│   main.go   // main file
│
├───config
│       serviceaccount.json		//Firebase Service Account
│
├───endpoints
│       authenticate.go			//Http Handler  
│       transaction.go			//Http Handler
│       users.go			//Http Handler	
│       wallets.go
│
├───entities
│       authenticate.go			//struct 
│       transaction.go			//struct
│       user.go				//struct
│       wallet.go			//struct
│
├───router
│       router.go			//all http handlers
│
└───services
        authenticate-service.go		//functions to handle authentications
        db-service.go			//functions to establish connection to firestore
        transactions-service.go		//recording transactions
        user-service.go			//create update users
        wallet-service.go		//get balance or topup wallet
