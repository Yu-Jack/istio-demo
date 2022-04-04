rebuild-product:
	docker rmi jk82421/product:v1
	docker build services/product/ -t jk82421/product:v1

rebuild-review-v1:
	docker rmi jk82421/review:v1
	docker build services/review/ -t jk82421/review:v1

rebuild-review-v2:
	docker rmi jk82421/review:v2
	docker build services/review/ -t jk82421/review:v2

build:
	docker build services/product/ -t jk82421/product:v1
	docker build services/review-v1/ -t jk82421/review:v1
	docker build services/review-v2/ -t jk82421/review:v2

push:
	docker push jk82421/product:v1
	docker push jk82421/review:v1
	docker push jk82421/review:v2

clear:
	docker rmi jk82421/product:v1
	docker rmi jk82421/review:v1
	docker rmi jk82421/review:v2

up:
	istioctl kube-inject -f k8s.yaml | kubectl apply -f -
	kubectl apply -f istio.yaml

down: 
	kubectl delete -f k8s.yaml
	kubectl delete -f istio.yaml