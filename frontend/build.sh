docker build --tag bggridley/frontend .
sudo docker save -o frontend.tar bggridley/frontend
sudo ctr -n k8s.io image import frontend.tar