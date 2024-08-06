docker build --tag bggridley/backend .
sudo docker save -o backend.tar bggridley/backend
sudo ctr -n k8s.io image import backend.tar
