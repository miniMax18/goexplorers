# How to launch our service urlshort golang: 
#

## Start of go by docker container in bash terminal (Linux or Mac):

### by Docker:
  
  ```sh
   docker build -t your_go_container123 .
  ```
  
  ```sh
   docker run --rm -it --privileged --group-add keep-groups -p 8081:8080 your_go_container123 bash
  ```
  
  ```sh
   root# go run ./main
  ```

### or by Podman with docker compatibility plug:
  ```sh
   podman build -t your_go_container123 .
  ```
  
  ```sh
   podman run --rm -it --privileged --group-add keep-groups -p 8081:8080 your_go_container123 bash
  ```
  
  ```sh
   root# go run ./main
  ```
