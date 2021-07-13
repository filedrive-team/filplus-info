filplus.info
==================
[![](https://img.shields.io/github/go-mod/go-version/filedrive-team/filplus-info)]()
[![](https://goreportcard.com/badge/github.com/filedrive-team/filplus-info)](https://goreportcard.com/report/github.com/filedrive-team/filplus-info)
[![](https://img.shields.io/github/license/filedrive-team/filplus-info)](https://github.com/filedrive-team/filplus-info/blob/main/LICENSE)

## Project Description
Filplus.info is the first visual dashboard for Filecoin Plus, which can be regarded as an incentive program for useful storage on Filecoin. Through monitoring on-chain messages on Filecoin and off-chain information on GitHub, filplus.info records relevant real-time data of each DataCap allocation and verified deal transactions on Filecoin from public APIs and open-source platforms.
On filplus.info, there are visual data analysis charts, information of notaries and clients, and deal-flow of DataCap. It also supports multiple parameters search and a hyperlink jump to the Filecoin browser filscan.io.

*We have called [filscan.io](https://filscan.io/) api in our codes!*

## Project Structure
```
├── Makefile
├── Readme.md  
├── api
│   ├── api.go
│   ├── public
│   │   ├── clientdeals.go
│   │   ├── datacapallocated.go
│   │   └── notary.go
│   └── requestparam.go
├── common
│   └── globalcache.go
├── conf
│   └── app.toml
├── errormsg
│   └── errormsg.go
├── frontend
│   ├── Makefile
│   ├── README.md
│   ├── babel.config.js
│   ├── package.json
│   ├── public
│   │   ├── favicon.png
│   │   └── index.html
│   ├── src
│   │   ├── App.vue
│   │   ├── api
│   │   │   ├── fetch.js
│   │   │   └── notary.js
│   │   ├── assets
│   │   │   ├── fil.png
│   │   │   ├── filedrive-logo.png
│   │   │   ├── logo-color-dark.png
│   │   │   └── logo-color-dark.svg
│   │   ├── components
│   │   │   ├── footer
│   │   │   │   └── index.vue
│   │   │   ├── header
│   │   │   │   └── index.vue
│   │   │   ├── index.js
│   │   │   ├── search
│   │   │   │   └── index.vue
│   │   │   └── table
│   │   │       ├── AdvancedTable.vue
│   │   │       ├── BaseTable.vue
│   │   │       └── Expand.vue
│   │   ├── i18n.js
│   │   ├── locales
│   │   │   └── en.json
│   │   ├── main.js
│   │   ├── mixin
│   │   │   └── index.js
│   │   ├── router
│   │   │   └── index.js
│   │   ├── store
│   │   │   └── index.js
│   │   ├── style
│   │   │   ├── class.scss
│   │   │   ├── index.scss
│   │   │   └── variable.scss
│   │   └── views
│   │       ├── Home.vue
│   │       ├── allocation
│   │       │   └── index.vue
│   │       ├── client
│   │       │   └── index.vue
│   │       └── notary
│   │           ├── index.vue
│   │           └── notary.vue
│   └── vue.config.js
├── go.mod
├── go.sum
├── http_api_test
│   ├── api.http
│   └── http-client.env.json
├── jobs
│   ├── crawler.go
│   └── syncer.go
├── log
│   └── logger.go
├── main.go
├── middleware
│   └── cors
│       └── cors.go
├── models
│   ├── client.go
│   ├── client_allowance.go
│   ├── client_test.go
│   ├── models.go
│   ├── notary.go
│   ├── notary_allowance.go
│   └── notary_test.go
├── routers
│   └── routers.go
├── rpc
│   ├── rpc_call.go
│   └── rpc_notary.go
├── settings
│   ├── constants.go
│   ├── loadconfig.go
│   └── settingtypes
│       └── tomltypes.go
├── types
│   ├── common.go
│   ├── normaltime.go
│   └── unixtime.go
└── utils
└── common.go
```
## Dependency
Environment
```
go version 1.14
node.js version 14.17.3
npm version 6.14.13
```

## Config
#### Back-end config
Please set databases with your username and password in conf/app.toml.
```
[app]
# development,test,product
runmode = "product"

[product]
[product.server]
httpPort = 8088
readTimeOut = 60
writeTimeout = 60

[product.database]
# mysql,postgres
type = "postgres"
user = "your_user"
password = "your_password"
host = "your_host"
name = "your_dbname"
```

#### Front-end config
Please set back-end api with your url in frontend/.env.xxx.

local server: .env.development
```
VUE_APP_BASE_URL=http://localhost:8088
```
product server: .env.production
```
VUE_APP_BASE_URL=YOUR_BACKEND_API_URL
```

## Deployment
### Build all
```bash
make
./filplus-info
````
### Only build back-end
```bash
make backend
./filplus-info
```
### Only build front-end
```bash
make frontend
```

### Dist
```bigquery
make frontend
copy all files of dist into the root of website
```


## Contribute
PRs are welcome!

## License
MIT

