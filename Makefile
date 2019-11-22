generate-db:
	sqlboiler psql

port-tiller:
	kubectl port-forward svc/tiller-deploy -n kube-system 44134

gen-proto:
	cd iproto && protoc --go_out=. *.proto