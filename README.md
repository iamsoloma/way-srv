**Way server is a chain management system.**

Create metadata/config.yaml  
>env: "prod"  
mode: "node"  
storage_path: "./blockchains"  
http_server:  
  address: "0.0.0.0:1436"  
  timeout: 4s  
  idle_timeout: 30s  
  
Create a empty folder 
>/blockchains  
  
Run: 
>docker run -d -p 1436:1436 -v $(pwd)/blockchains:/way/blockchains -v $(pwd)/metadata:/way/metadata way-srv:
prod
